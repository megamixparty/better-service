package customer

import "github.com/go-chi/chi"

func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", createHandler)
	r.Get("/", listHandler)
	r.Get("/{id}", getHandler)
	r.Delete("/{id}", deleteHandler)
	return r
}
