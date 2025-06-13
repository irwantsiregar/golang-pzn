package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	fmt.Println(now)

	var utc time.Time = time.Date(2009, time.August, 17, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	formatter := "2006-01-02 15:04:05"

	value := "2020-10-10 10:10:10"
	// value := "ASAL"
	valueTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(valueTime)
	}

	fmt.Println(valueTime.Year())
	fmt.Println(valueTime.Month())
	fmt.Println(valueTime.Day())
	fmt.Println(valueTime.Hour())
}


/*
[ Package time ]
- Package time adalah package yang berisikan fungsionalitas untuk management waktu di Go-Lang

[ Beberapa Function di Package time ]
Function		Kegunaan
time.Now()		Untuk mendapatkan waktu saat ini
time.Date(...)	Untuk membuat waktu
time.Parse(layout, string)	Untuk memparsing waktu dari string


Good Know:
Pada GO-Lang untuk melakukan format terhadap date adalah dengan cara langsung melampirkan date string, seperti berikut: 
	- 2006-01-02 15:04:05

Berbeda dengan bahasa lain, seperti JavaScript yang melampirkan formate date,  seperti berikut: 
	- YYYY-MM-DD HH:ss:dd
*/