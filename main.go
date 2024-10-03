package main

import (
	"net/http"

	"sobczyk.dev/url-shortener/routes"
)

func main() {
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)

	http.ListenAndServe(":3000", mux)
}
