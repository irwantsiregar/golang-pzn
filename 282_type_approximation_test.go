package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type AgeB int

type NumberB interface {
	~int | int8 | int16 | int32 | int64 |
		float32 | float64
}

func MinB[T NumberB](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestMinB(t *testing.T) {
	assert.Equal(t, 100, MinB[int](100, 200))
	assert.Equal(t, int64(100), MinB[int64](int64(100), int64(200)))
	assert.Equal(t, float64(100), MinB[float64](float64(100), float64(200)))
	assert.Equal(t, AgeB(100), MinB[AgeB](AgeB(100), AgeB(200)))
}
