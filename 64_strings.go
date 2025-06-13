package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Eko Kurniawan", "Eko"))
	fmt.Println(strings.Split("Eko Kurniawan", " "))
	fmt.Println(strings.ToLower("Eko Kurniawan"))
	fmt.Println(strings.ToUpper("Eko Kurniawan"))
	fmt.Println(strings.Trim("      Eko Kurniawan      ", " "))
	fmt.Println(strings.ReplaceAll("Eko Kurniawan Eko Khannedy", "Eko", "Budi"))
}


/*
[ Package strings ]

- Package strings adalah package yang berisikan function-function untuk memanipulasi tipe data String
- Ada banyak sekali function yang bisa kita gunakan
- https://golang.org/pkg/strings/
*/
