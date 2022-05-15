package openapi

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi2conv"
	"github.com/getkin/kin-openapi/openapi3"
)

const (
	DocsJSONPath  = "/doc.json"
	DocsIndexPath = "/index.html"
)

type Handler struct {
	docsPath string
	oAPI3doc []byte
	index    *template.Template
}

func NewHandler(docsPath, swaggerDoc string) (*Handler, error) {
	oAPI3doc, err := initAPIDoc(swaggerDoc)
	if err != nil {
		return nil, err
	}

	index, err := setupIndex()
	if err != nil {
		return nil, err
	}

	return &Handler{docsPath: docsPath, oAPI3doc: oAPI3doc, index: index}, nil
}

func initAPIDoc(swaggerDoc string) ([]byte, error) {
	oAPI2 := &openapi2.T{}

	err := json.Unmarshal([]byte(swaggerDoc), oAPI2)
	if err != nil {
		return nil, fmt.Errorf("unmarshaling doc string to swagger2.0: %w", err)
	}

	oAPI3, err := openapi2conv.ToV3(oAPI2)
	if err != nil {
		return nil, fmt.Errorf("converting swagger2.0 to openapi3: %w", err)
	}

	oAPI3.AddServer(&openapi3.Server{URL: oAPI2.BasePath})

	marshaled, err := json.Marshal(oAPI3)
	if err != nil {
		return nil, fmt.Errorf("marshaling openapi3 doc: %w", err)
	}

	return marshaled, nil
}

func setupIndex() (*template.Template, error) {
	t := template.New("swagger_index.html")

	return t.Parse(indexTempl)
}

func (h *Handler) DocJSON(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(h.oAPI3doc); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) RedirectToIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, h.docsPath+DocsIndexPath, http.StatusMovedPermanently)
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	if err := h.index.Execute(w, h.docsPath+DocsJSONPath); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
