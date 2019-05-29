package main

import (
	"flag"
	"fmt"
	"net/http"

	// "github.com/guilledipa/urlshort"
	"proyectos/urlshort"
)

var (
	targets = flag.String("targets", "./targets.yaml", "Quiz filename.")
	port    = flag.Int("port", 8080, "Port to listen on.")
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	fallBackHandler := urlshort.MapHandler(pathsToUrls, mux)

	yamlHandler, err := urlshort.YAMLHandler(*targets, fallBackHandler)
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
