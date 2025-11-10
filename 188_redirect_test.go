package main

import (
	"fmt"
	"net/http"
	"testing"
)

func RedirectTo(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello Redirect")
}

func RedirectForm(writer http.ResponseWriter, request *http.Request) {
	http.Redirect(writer, request, "/redirect-to", http.StatusTemporaryRedirect)
}


func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()
	
	mux.HandleFunc("/redirect-from", RedirectForm)
	mux.HandleFunc("/redirect-to", RedirectTo)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

