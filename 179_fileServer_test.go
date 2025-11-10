package main

import (
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	directory := http.Dir("./resources")
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()
	
	// mux.Handle("/static/", fileServer) // Result: 404 page not found

	mux.Handle("/static/", http.StripPrefix("/static", fileServer)) // Used '/static' as new prepix

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

