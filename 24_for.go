package main

import "fmt"

func main() {
	//counter := 1
	//
	//for counter <= 10 {
	//	fmt.Println("Perulangan ke ", counter)
	//	counter++
	//}

	// For Loop with Counter
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}
	fmt.Println("Selesai")

	// For Loop with Slice
	names := []string{"Eko", "Kurniawan", "Khannedy"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// For Loop with Range
	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
