package helper

import (
	"fmt"
	"testing"

	// "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)


// func TestHelloWorlAssertion(t *testing.T){
// 	result := HelloWorld("Irwan")
// 	fmt.Println("Dieksekusi")

// 	assert.Equal(t, "Hello Irwan", result);

// }

func TestHelloWorldRequire(t *testing.T){
	result := HelloWorld("Irwan")

	require.Equal(t, "Hello Irwan", result);

	fmt.Println("Tidak Dieksekusi")
}