package test

import "testing"

func BenchmarkSub(b *testing.B){
	b.Run("Irwan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Irwan")
		}
	})

	b.Run("Siregar", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Siregar")
		}
	})
}