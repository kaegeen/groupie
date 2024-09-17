package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("templates"))
	http.Handle("/templates/", http.StripPrefix("/templates", fs))
	http.HandleFunc("/", router)

	fmt.Println("http://localhost:8585")
	if err := http.ListenAndServe(":8585", nil); err != nil {
		log.Fatal(err)
	}
}
