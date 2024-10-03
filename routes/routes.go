package routes

import (
	"net/http"

	"sobczyk.dev/url-shortener/handlers"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /{code}", handlers.HandleGetLink)
	mux.HandleFunc("POST /{code}", handlers.HandlePostLink)
	mux.HandleFunc("DELETE /{code}", handlers.HandleDeleteLink)
	mux.HandleFunc("GET /", handlers.HandleGetAllLinks)
}
