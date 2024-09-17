package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func loadAsset(w http.ResponseWriter, r *http.Request, file string) {
	data, err := os.ReadFile("." + file)
	if err != nil {
		log.Printf("Error reading file %s: %v", file, err)
		http.Error(w, "Couldn't read file", http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(file)
	var mimeType string

	switch ext {
	case ".css":
		mimeType = "text/css"
	case ".js":
		mimeType = "application/javascript"
	default:
		mimeType = http.DetectContentType(data)
	}

	w.Header().Set("Content-Type", mimeType+"; charset=utf-8")
	w.Write(data)
}

func executeTemplate(w http.ResponseWriter, templateFiles []string, data interface{}) {
	files := append([]string{
		"./templates/base.html",
		"./templates/nav.html",
	}, templateFiles...)

	t, err := template.ParseFiles(files...)
	if err != nil {
		log.Printf("Error parsing templates: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := t.ExecuteTemplate(w, "base", data); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func getArtistList() []Artist {
	var artists []Artist

	resp, err := callApi(Api["artists"])
	if err != nil {
		log.Printf("Error calling API: %v", err)
		return nil
	}

	if err := json.Unmarshal([]byte(resp), &artists); err != nil {
		log.Printf("Error unmarshaling artist list: %v", err)
		return nil
	}

	for i, a := range artists {
		artists[i].AuthorMemb = "Author"
		if len(a.Members) > 1 {
			artists[i].AuthorMemb = "Member"
		}
		artists[i].AllNames = strings.Join(append([]string{a.Name}, a.Members...), "|")
		artists[i].Concerts = getConcerts(a.Relations)
	}

	return artists
}

func getArtist(id int) Artist {
	var artist Artist

	resp, err := callApi(Api["artists"] + "/" + strconv.Itoa(id))
	if err != nil {
		log.Printf("Error calling API for artist %d: %v", id, err)
		return artist
	}

	if err := json.Unmarshal([]byte(resp), &artist); err != nil {
		log.Printf("Error unmarshaling artist data: %v", err)
		return artist
	}

	artist.Concerts = getConcerts(artist.Relations)
	return artist
}

func getConcerts(url string) []Concert {
	var data map[string]interface{}

	resp, err := callApi(url)
	if err != nil {
		log.Printf("Error calling API for concerts: %v", err)
		return nil
	}

	if err := json.Unmarshal([]byte(resp), &data); err != nil {
		log.Printf("Error unmarshaling concert data: %v", err)
		return nil
	}

	var concerts []Concert
	for location, dates := range data["datesLocations"].(map[string]interface{}) {
		locParts := strings.Split(location, "-")
		city, country := locParts[0], locParts[1]

		var concertDates []string
		for _, date := range dates.([]interface{}) {
			concertDates = append(concertDates, date.(string))
		}

		concerts = append(concerts, Concert{
			City:    city,
			Country: country,
			Dates:   concertDates,
		})
	}

	return concerts
}

func home(w http.ResponseWriter, r *http.Request) {
	artists := getArtistList()
	if artists == nil {
		http.Error(w, "Could not retrieve artists", http.StatusInternalServerError)
		return
	}

	sort.Slice(artists, func(i, j int) bool {
		return artists[i].Name < artists[j].Name
	})

	executeTemplate(w, []string{"./templates/home.html"}, artists)
}

func all(w http.ResponseWriter, r *http.Request) {
	artists := getArtistList()
	if artists == nil {
		http.Error(w, "Could not retrieve artists", http.StatusInternalServerError)
		return
	}

	sort.Slice(artists, func(i, j int) bool {
		return artists[i].Name < artists[j].Name
	})

	executeTemplate(w, []string{"./templates/all.html"}, artists)
}

func artist(w http.ResponseWriter, r *http.Request, id int) {
	artist := getArtist(id)
	if artist.Name == "" {
		http.NotFound(w, r)
		return
	}

	executeTemplate(w, []string{"./templates/artist.html"}, artist)
}
