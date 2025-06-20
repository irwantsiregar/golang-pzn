package test

import (
	"testing"
)

// go test -v -run=TidakAda -bench=BenchmarkTable
func BenchmarkTable(b *testing.B){
	benchmarks := []struct{
		name string
		request string
	}{
		{
			name: "Irwan",
			request: "Irwan",
		},
	 	{
			name: "Siregar",
			request: "Siregar",
		},
		{
			name: "Budiono Siregar",
			request: "Budiono Siregar",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}