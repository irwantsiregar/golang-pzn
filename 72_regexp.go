package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("edo"))
	fmt.Println(regex.MatchString("eKo"))

	fmt.Println(regex.FindAllString("eko edo egi ego e1o eto eKo", 10))
}


/*
[ Package regexp ]

- Package regexp adalah utilitas di Go-Lang untuk melakukan pencarian regular expression
- Regular expression di Go-Lang menggunakan library C yang dibuat Google bernama RE2
- https://github.com/google/re2/wiki/Syntax
- https://golang.org/pkg/regexp/

 Function							Kegunaan
regexp.MustCompile(string)			Membuat Regexp
Regexp.MatchString(string) bool		Mengecek apakah Regexp match dengan string
Regexp.FindAllString(string, max)	Mencari string yang match dengan maximum jumlah hasil

*/