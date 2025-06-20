package helper

import (
	"testing"

	"github.com/stretchr/testify/require"
)


func TestSubTest(t *testing.T){
	t.Run("Irwan", func (t *testing.T)  {
		result := HelloWorld("Irwan")

		require.Equal(t, "Hello Irwan", result);	
	})

	t.Run("Siregar", func (t *testing.T)  {
		result := HelloWorld("Irwan")

		require.Equal(t, "Hello Siregar", result);	
	})
}