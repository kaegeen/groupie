package main

import (
	"fmt"
	"io"
	"net/http"
)

var Api = map[string]string{
	"artists":   "https://groupietrackers.herokuapp.com/api/artists",
	"locations": "https://groupietrackers.herokuapp.com/api/locations",
	"dates":     "https://groupietrackers.herokuapp.com/api/dates",
	"relation":  "https://groupietrackers.herokuapp.com/api/relation",
}

func callApi(url string) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {

		fmt.Printf("Error fetching data from %s: %v\n", url, err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		fmt.Printf("Unexpected response status from %s: %s\n", url, resp.Status)
		return nil, fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {

		fmt.Printf("Error reading response body from %s: %v\n", url, err)
		return nil, err
	}

	return body, nil
}
