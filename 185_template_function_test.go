package main

import (
	_ "embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"text/template"
)

/*
* Using TEMPLATE FUNCTION
 */

 type MyPage struct {
	Name string
 }

 func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", My Name Is " + myPage.Name
 }

 // With 'Function'
func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "Budi"}}`))

	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Irwan",
	})
}
func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil);
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)
	
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)	
	
	fmt.Println(string(body))
}

// With 'Global Function'
func TemplateFunctionGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))

	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Irwan",
	})
}
func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil);
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)
	
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)	
	
	fmt.Println(string(body))
}

// With 'Added: Create Global Function'
func TemplateFunctionCreateGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION")
	
	t = t.Funcs(map[string]interface{}{
		"upper": func (value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse(`{{ upper .Name}}`))

	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Irwan Siregar",
	})
}
func TestTemplateFunctionCreateGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil);
	recorder := httptest.NewRecorder()

	TemplateFunctionCreateGlobal(recorder, request)
	
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)	
	
	fmt.Println(string(body))
}


// With 'Added: with Function PipeLines'
// The result from Function can be return value in between function to function
func TemplateFunctionCreateGlobalPipeline(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION")
	
	t = t.Funcs(map[string]interface{}{
		"sayHello": func (name string) string {
			return "Hello " + name
		},
		"upper": func (value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse(`{{ sayHello .Name | upper }}`))

	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Irwan Siregar",
	})
}
func TestTemplateFunctionCreateGlobalPipeline(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil);
	recorder := httptest.NewRecorder()

	TemplateFunctionCreateGlobalPipeline(recorder, request)
	
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)	
	
	fmt.Println(string(body))
}


