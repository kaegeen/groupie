package main

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
)

func router(w http.ResponseWriter, r *http.Request) {

	artistExp := regexp.MustCompile(`^/artist/([0-9]+)$`)
	assetExp := regexp.MustCompile(`^/templates/(js|css)/([^/]+)\.(js|css)$`)

	switch {
	case r.URL.Path == "/":
		home(w, r)
	case r.URL.Path == "/all":
		all(w, r)
	case artistExp.MatchString(r.URL.Path):

		matches := artistExp.FindStringSubmatch(r.URL.Path)
		if len(matches) > 1 {
			id, err := strconv.Atoi(matches[1])
			if err != nil {
				log.Printf("Invalid artist ID: %s", matches[1])
				http.Error(w, "Invalid artist ID", http.StatusBadRequest)
				return
			}
			artist(w, r, id)
		} else {
			http.NotFound(w, r)
		}
	case assetExp.MatchString(r.URL.Path):
		loadAsset(w, r, r.URL.Path)
	default:
		http.NotFound(w, r)
	}
}
