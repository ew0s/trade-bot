package openapi

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/stretchr/testify/suite"
	"github.com/swaggo/swag"

	_ "github.com/ew0s/trade-bot/cmd/api/swagger"
)

type openapiSuite struct {
	suite.Suite

	basePath, correctDoc string
	docsRequest          *http.Request

	handler *Handler
}

func (s *openapiSuite) SetupTest() {
	s.basePath = "/trade-bot/api/v1/docs"
	s.handler = s.setupHandler()
	s.docsRequest = httptest.NewRequest(http.MethodGet, DocsJSONPath, nil)
}

func (s *openapiSuite) setupHandler() *Handler {
	var err error

	s.correctDoc, err = swag.ReadDoc()
	s.NoError(err)

	swaggerHandler, err := NewHandler(s.basePath, s.correctDoc)
	s.NoError(err)

	return swaggerHandler
}

func TestOpenapiHandler(t *testing.T) {
	suite.Run(t, new(openapiSuite))
}

func (s *openapiSuite) TestNewHandler_ValidateErr() {
	testCases := []struct {
		name, swagger, errMsg string
	}{
		{
			name:    "corrupted swag doc",
			swagger: "}{",
		},
		{
			name:    "not convertible to openapi3",
			swagger: WrongSwaggerHostFormat,
		},
	}

	for _, testCase := range testCases {
		tc := testCase

		s.Run(tc.name, func() {
			_, err := NewHandler(s.basePath, tc.swagger)

			s.Error(err)
		})
	}
}

func (s *openapiSuite) TestDocJSON() {
	responseRecorder := httptest.NewRecorder()

	s.handler.DocJSON(responseRecorder, s.docsRequest)

	s.assertDocIsOpenAPI3(responseRecorder.Body)
}

func (s *openapiSuite) assertDocIsOpenAPI3(r io.Reader) {
	oAPI3Doc := &openapi3.T{}
	// check that spec is in openapi3 format
	err := json.NewDecoder(r).Decode(&oAPI3Doc)

	s.NoError(err)
}
