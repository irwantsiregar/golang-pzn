package test

import "testing"

func BenchmarkHelloWorld(b *testing.B){
	for i := 0; i < b.N; i++ {
		HelloWorld("Irwan")
	}
}

func BenchmarkHelloWorldSiregar(b *testing.B){
	for i := 0; i < b.N; i++ {
		HelloWorld("Siregar")
	}
}