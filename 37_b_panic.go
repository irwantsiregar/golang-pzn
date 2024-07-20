package main

import "fmt"

// Fungsionalitas mirip seperti 'try catch' jika di Javascript atau lainnya.

func endApp() {
	fmt.Println("End app")
	message := recover()
	fmt.Println("terjadi panic", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Ups Error")
	}
}

func main() {
	runApp(true)
	fmt.Println("Eko Kurniawan Khannedy")
}

/*
[ Panic ]
- Panic function adalah function yang bisa kita gunakan untuk menghentikan program.
- Panic function biasanya dipanggil ketika terjadi panic pada saat program kita berjalan.
- Saat panic function dipanggil, program akan terhenti, namun defer function tetap akan dieksekusi.

[ Recover ]
- Recover adalah function yang bisa kita gunakan untuk menangkap data panic.
- Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan.
*/
