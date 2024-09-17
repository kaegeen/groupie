package main

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Relations    string   `json:"relations"`
	Concerts     []Concert
	AuthorMemb   string
	AllNames     string
}

type Concert struct {
	Country string
	City    string
	Dates   []string
}
