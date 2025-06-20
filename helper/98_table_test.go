package helper

import (
	"testing"

	"github.com/stretchr/testify/require"
)


func TestHelloWorldTable(t *testing.T){
	tests := []struct{
		name string
		request string
		expected string
	}{
		{
			name: "HelloWorld(Irwan)",
			request: "Irwan",
			expected: "Hello Irwan",
		},
		{
			name: "HelloWorld(Siregar)",
			request: "Siregar",
			expected: "Hello Siregar",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func (t *testing.T)  {
			result := HelloWorld(test.request)
			
			require.Equal(t, test.expected, result);	
		})
	}

}