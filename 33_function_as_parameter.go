package main

import "fmt"

// type declarations
type Filter func(string) string

// ## Function As Parameters:
/*
func sayHelloWithFilter(name string, filterFn func(string) string) {
	filteredName := filterFn(name)
	fmt.Println("Hello", filteredName)
}
*/

func sayHelloWithFilter(name string, filterFn Filter) {
	filteredName := filterFn(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Eko", spamFilter)

	filterFn := spamFilter
	sayHelloWithFilter("Anjing", filterFn)
}
