package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


type AgeC int

type NumberC interface {
	~int | int8 | int16 | int32 | int64 |
		float32 | float64
}

func MinC[T NumberC](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestMinC(t *testing.T) {
	assert.Equal(t, 100, MinC[int](100, 200))
	assert.Equal(t, int64(100), MinC[int64](int64(100), int64(200)))
	assert.Equal(t, float64(100), MinC[float64](float64(100), float64(200)))
	assert.Equal(t, AgeC(100), MinC[AgeC](AgeC(100), AgeC(200)))
}

func TestMinTypeInferenceC(t *testing.T) {
	assert.Equal(t, 100, MinC(100, 200))
	assert.Equal(t, int64(100), MinC(int64(100), int64(200)))
	assert.Equal(t, float64(100), MinC(float64(100), float64(200)))
	assert.Equal(t, AgeC(100), MinC(AgeC(100), AgeC(200)))
}