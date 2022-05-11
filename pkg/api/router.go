package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpswagger "github.com/swaggo/http-swagger"
)

type Router interface {
	Routes() chi.Router
	BasePrefix() string
}

func MakeRoutes(basePrefix string, routers []chi.Router) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)

	for _, router := range routers {
		r.Mount(basePrefix, router)
	}

	r.Get("/swagger/*", httpswagger.Handler(
		httpswagger.URL("http://localhost:5000/swagger/doc.json"),
	))

	return r
}
