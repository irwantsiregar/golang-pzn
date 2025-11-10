package main

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFileHttp(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		http.ServeFile(writer, request, "./resources/oke.html")
	} else {
		http.ServeFile(writer, request, "./resources/notfound.html")
	}
}

/*
*
* Using GoLang Embed
*/

//go:embed resources/ok.html
var resourcesOk string

//go:embed resources/notfound.html
var resourcesNotfound string

func ServeFileEmbed(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		fmt.Fprint(writer, resourcesOk)
	} else {
		fmt.Fprint(writer, resourcesNotfound)
	}
}

func TestServeFileServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(ServeFileHttp),
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func TestServeFileServerEmbed(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

