package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

// XSS: Cross Site Scripting
// # In Latest Version: For now, AutoEscape no built in Go Lang (go1.25.4 linux/amd64)

/*
* Using TEMPLATE Escape
 */
func TemplateAutoEscape(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
			"Title": "Template Auto Escape",
			"Body": template.HTMLEscapeString("<p>Ini Adalah Body</p>"),
	})
}
func TestTemplateAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil);
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)
	
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)	
	
	fmt.Println(string(body))
}

func TestTemplateAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

/*
* Using TEMPLATE AutoEscape
 */
func TemplateAutoEscapeDisabled(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
			"Title": "Template Auto Escape",
			"Body": "<p>Ini Adalah Body</p>",
	})
}
func TestTemplateAutoEscapeDisabled(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil);
	recorder := httptest.NewRecorder()

	TemplateAutoEscapeDisabled(recorder, request)
	
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)	
	
	fmt.Println(string(body))
}

func TestTemplateAutoEscapeDisabledServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscapeDisabled),
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

/*
* Using TEMPLATE XSS
 */
func TemplateXSS(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
			"Title": "Template Auto Escape",
			"Body": request.URL.Query().Get("body"),
	})
}
func TestTemplateXSS(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?body=<p>alert</p>", nil);
	recorder := httptest.NewRecorder()

	TemplateXSS(recorder, request)
	
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)	
	
	fmt.Println(string(body))
}

func TestTemplateXSSServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(TemplateXSS),
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}