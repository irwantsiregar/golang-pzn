package main

import (
	"context"
	"fmt"
	"testing"
)

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")
	contextG := context.WithValue(contextF, "g", "G")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextF)
	fmt.Println(contextE)
	fmt.Println(contextG)

	fmt.Println(contextF.Value("f")) // can
	fmt.Println(contextF.Value("c")) // can't have parent
	fmt.Println(contextF.Value("b")) // can't, parent is different
	fmt.Println(contextA.Value("b")) // can't get data child
}
