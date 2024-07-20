package main

import "fmt"

// parameters with var_args
func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	// arguments with single value <int>
	fmt.Println(sumAll(10, 10, 10))
	fmt.Println(sumAll(10, 10, 10, 10, 10, 10))
	fmt.Println(sumAll(10, 10, 10, 10, 10, 10, 10, 10, 10))

	// arguments with slice method
	numbers := []int{10, 10, 10, 10}
	fmt.Println(sumAll(numbers...))
}
