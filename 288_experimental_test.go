package main

import (
	"maps"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

func ExperimentalMin[T constraints.Ordered](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestExperimentalMin(t *testing.T) {
	assert.Equal(t, 100, ExperimentalMin(100, 200))
	assert.Equal(t, 100.0, ExperimentalMin(100.0, 200.0))
}

func TestExperimentalMaps(t *testing.T) {
	first := map[string]string{
		"Name": "Eko",
	}

	second := map[string]string{
		"Name": "Eko",
	}

	assert.True(t, maps.Equal(first, second))
}

func TestExperimantalSlice(t *testing.T) {
	first := []string{"Eko"}
	second := []string{"Eko"}

	assert.True(t, slices.Equal(first, second))
}