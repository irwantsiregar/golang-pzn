package helper

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)


func TestSkip(t *testing.T){
	if runtime.GOOS == "darwin" {
		t.Skip("Unit test tidak bisa jalan di Mac OS")
	}

	result := HelloWorld("Irwan")

	require.Equal(t,"Hello Irwan", result)
}