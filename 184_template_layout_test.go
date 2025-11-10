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
* Using TEMPLATE LAYOUT
 */

func TemplateLayout(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles(
		"./templates/layout/header.gohtml",
		"./templates/layout/content.gohtml",
		"./templates/layout/footer.gohtml",
		))

	t.ExecuteTemplate(writer, "layout", map[string]interface{}{
			"Title": "Template Layout",
			"Name": "Irwan",
	})
}
func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil);
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)
	
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)	
	
	fmt.Println(string(body))
}


