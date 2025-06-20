package helper

import (
	"fmt"
	"testing"
)


func TestHelloWorldError(t *testing.T){
	result := HelloWorld("Irwands")

	if result != "Hello Irwan" {
		t.Error("Harusnya Hello Irwan")
	}

	fmt.Println("Dieksekusi")
}

func TestHelloWorldFatal(t *testing.T){
	result := HelloWorld("Irwancs")

	if result != "Hello Irwan" {
		t.Fatal("Harusnya Hello Irwan")
	}

	fmt.Println("Tidak dieksekusi")
}