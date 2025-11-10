package main

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

/*
* Using TEMPLATE CACHING
 */

//go:embed templates/*.gohtml
var _templates embed.FS

// Initialize as global variabel for the caching template.
var myTemplates = template.Must(template.ParseFS(_templates, "templates/*.gohtml"));

 // With 'Function'
func TemplateCaching(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "simple.gohtml", "Hello Template Caching")
}
func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil);
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder, request)
	
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)	
	
	fmt.Println(string(body))
}
