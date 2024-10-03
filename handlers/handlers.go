package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"sobczyk.dev/url-shortener/links"
)

func HandleGetLink(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	if link, exists := links.Links[code]; exists {
		http.Redirect(w, r, link, http.StatusFound)
	} else {
		http.NotFound(w, r)
	}
}

type UrlData struct {
	Url string
}

func HandlePostLink(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	var data UrlData

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	links.Links[code] = data.Url

	for key, value := range links.Links {
		fmt.Println(key, value)
	}
}

func HandleDeleteLink(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	delete(links.Links, code)
}

func HandleGetAllLinks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(links.Links)
}
