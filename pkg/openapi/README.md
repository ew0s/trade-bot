# OpenAPIHandler

`OpenAPIHandler` предоставляет swagger-ui. Документация сервиса генерится из комментариев при помощи инструмента [swaggo/swag](https://github.com/swaggo/swag).
За примерами написания комментариев в нужном формате и генерации документации необходимо обращаться в [readme swaggo/swag](https://github.com/swaggo/swag#swag).

### Usage

Использовать `OpenAPIHandler` нужно там, где инициализируются роуты приложения.

```go
package main

import (
	"fmt"
    "net/http"

    "github.com/go-chi/chi"
    "github.com/swaggo/swag"

    _ "github.com/ew0s/trade-bot/cmd/api/swagger" // импорт пакета, сгенерированного swag
    "github.com/ew0s/trade-bot/pkg/openapi"
)

func Router() http.Handler {
	r := chi.NewMux()
    docsPath := "/docs"

    openapiHandler, err := setupOpenapiHandler(docsPath)
    if err != nil {
    	// handle error - can't init swagger"
    }

    r.Route(docsPath, func(r chi.Router) {
    	r.Get(openapi.DocsJSONPath, openapiHandler.DocJSON)
    	r.Get(openapi.DocsIndexPath, openapiHandler.Index)
    	r.Get("/*", openapiHandler.RedirectToIndex)
    })

    return r
}

func setupOpenapiHandler(docsPath string) (*openapi.Handler, error) {
	doc, err := swag.ReadDoc()
	if err != nil {
		return nil, fmt.Errorf("reading swagger (make sure doc import is presented): %w", err)
	}

	openapiHandler, err := openapi.NewHandler(docsPath, doc)
	if err != nil {
		return nil, fmt.Errorf("initializing openapi handler: %w", err)
	}

	return openapiHandler, nil
}
```
