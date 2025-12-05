package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestSample(t *testing.T) {
	var result string = Length[string]("Eko")
	assert.Equal(t, "Eko", result)

	var resultNumber int = Length[int](100)
	assert.Equal(t, 100, resultNumber)
}
