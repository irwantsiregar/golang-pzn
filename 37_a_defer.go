package main

import "fmt"

// Fungsionalitas mirip seperti 'try catch' jika di Javascript atau lainnya. di Go-Lang ada Defer, Panic & Recover.

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer logging()
	fmt.Println("Run application")
}

func main() {
	runApplication()
}

/*
[ Defer ]
- Defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai di eksekusi.
- Defer akan selalu dieksekusi walaupun terjadi error di function yang dieksekusi.
*/
