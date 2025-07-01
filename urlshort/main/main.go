package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"Gophercises/urlshort"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	jsonContent, err := os.ReadFile("urlshort/resources/redirects.json")
	if err != nil {
		fmt.Errorf("failed to read json file", err)
	}

	var pathsToUrls map[string]string

	if err := json.Unmarshal(jsonContent, &pathsToUrls); err != nil {
		fmt.Errorf("failed to parse json content", err)
	}

	// pathsToUrls := map[string]string{
	// 	"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
	// 	"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	// }
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yamlContent, err := os.ReadFile("urlshort/resources/redirects.yml")
	if err != nil {
		panic(err)
	}

	yamlHandler, err := urlshort.YAMLHandler(yamlContent, mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
