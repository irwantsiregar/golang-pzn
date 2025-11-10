package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name");

	if name == "" {
		fmt.Fprint(writer, "Hello")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func MultipleParameter(writer http.ResponseWriter, request *http.Request) {
	firstName := request.URL.Query().Get("first_name");
	lastName := request.URL.Query().Get("last_name");

	fmt.Fprintf(writer, "%s %s", firstName, lastName)
}

func MultipleParameter2(writer http.ResponseWriter, request *http.Request) {
	var query url.Values = request.URL.Query();
	var names []string = query["name"]

	fmt.Fprintln(writer, strings.Join(names, ","))
}

func TestQueryParams(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello?name=Irwan", nil);
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	MultipleParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)

	fmt.Println(string(body))
}

func TestMultipleValueQueryParams(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/?name=Irwan&name=Siregar", nil);
	recorder := httptest.NewRecorder()

	MultipleParameter2(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)

	fmt.Println(string(body))
}
