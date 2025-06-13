package main

import (
	"fmt"
	"time"
)

func main() {
	var duration1 time.Duration = 100 * time.Second
	var duration2 time.Duration = 10 * time.Millisecond
	var duration3 time.Duration = duration1 - duration2

	fmt.Println(duration3)
	fmt.Printf("duration : %d\n", duration3)
}

/*
[ Time: Duration ]

- Saat menggunakan tipe data waktu, kadang kita butuh data berupa durasi
- Package tipe memiliki type Duration, yang sebenarnya adalah alias untuk int64
- Namun terdapat banyak method yang bisa kita gunakan untuk memanipulasi Duration

Good Know:
 Time atau duration biasanya digunakan ketika ingin mengetahui waktu yg diperlukan untuk melakukan koneksi ke database.
*/