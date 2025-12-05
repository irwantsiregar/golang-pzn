package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type AgeA int

type NumberA interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
	}

func MinA[T NumberA](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestMinA(t *testing.T) {
	assert.Equal(t, 100, MinA[int](100, 200))
	assert.Equal(t, int64(100), MinA[int64](int64(100), int64(200)))
	assert.Equal(t, float64(100), MinA[float64](float64(100), float64(200)))
	// assert.Equal(t, AgeA(100), MinA[AgeA](AgeA(100), AgeA(200)))
}