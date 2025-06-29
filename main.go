package main

import (
	"log"
	"net/http"

	"github.com/Alexpud/MyGoApi/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	ch := handlers.NewCarros(log.Default())
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/api", func(r chi.Router) {
		r.Route("/carros", func(r chi.Router) {
			r.Get("/", ch.HandleList)
			r.Get("/search", ch.HandleSearch)
		})
	})
	http.ListenAndServe(":3000", r)
}
