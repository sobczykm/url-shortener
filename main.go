package main

import (
	"log"
	"net/http"

	"sobczyk.dev/url-shortener/routes"
)

func main() {
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)

	log.Default().Println("Listening on :3000")

	http.ListenAndServe(":3000", mux)
}
