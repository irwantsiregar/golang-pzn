package helper

import (
	"fmt"
	"testing"
)


func TestBeforeAfter(t *testing.M){
	// before
	fmt.Println("Before Unit Test")

	m.run()

	// after
	fmt.Println("After Unit Test")
}