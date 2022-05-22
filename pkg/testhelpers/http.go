package testhelpers

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"

	"github.com/go-chi/chi/v5"
)

func MakeTestHTTPRequest(
	method string,
	target string,
	body string,
	queryParams map[string]string,
	pathParams map[string]string,
) *http.Request {
	req := httptest.NewRequest(method, target, bytes.NewBufferString(body))

	rctx := chi.NewRouteContext()
	for k, v := range pathParams {
		rctx.URLParams.Add(k, v)
	}

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	q := req.URL.Query()
	for k, v := range queryParams {
		q.Add(k, v)
	}

	req.URL.RawQuery = q.Encode()

	return req
}
