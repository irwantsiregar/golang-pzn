package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Eko")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // tidak bisa diakses
	// fmt.Println(helper.sayGoodBye("Eko")) // tidak bisa diakses
}

/*
# [ Access Modifier ]
- Di bahasa pemrograman lain (Ex: public, protected, private), biasanya ada kata kunci yang bisa digunakan untuk menentukan access modifier terhadap suatu function atau variable.
- Di Go-Lang, untuk menentukan access modifier, cukup dengan nama function atau variable.
- Jika nama nya diawali dengan hurup besar, maka artinya bisa diakses dari package lain, jika dimulai dengan hurup kecil, artinya tidak bisa diakses dari package lain.
*/
