package main

import (
	_ "embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

/*
* Using TEMPLATE DATA
 */

// Data with MAP
func TemplateDataMap(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))

	t.ExecuteTemplate(writer, "name.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"Name": "Irwan",
		"Address": map[string]interface{}{
			"Street": "Jalan Belum Ada Lagi",
		},
	})
}
func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil);
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)
	
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)	
	
	fmt.Println(string(body))
}


// Data with STRUCT
type Address struct {
	Street string
}
type Page struct {
	Title string
	Name string
	Address Address
}

func TemplateDataStruct(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))

	t.ExecuteTemplate(writer, "name.gohtml", Page{
		Title: "Template Data Map",
		Name: "Irwan",
		Address: Address{
			Street: "Jl. Developer Stuck",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil);
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)
	
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)	
	
	fmt.Println(string(body))
}



