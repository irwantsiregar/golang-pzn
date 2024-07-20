package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// A function assigned to in Struct.
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var eko Customer
	fmt.Println(eko)

	eko.Name = "Eko_Kurniawan"
	eko.Address = "Indonesia"
	eko.Age = 30
	fmt.Println(eko)
	fmt.Println(eko.Name)
	fmt.Println(eko.Address)
	fmt.Println(eko.Age)

	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     30,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Indonesia", 30}
	fmt.Println(budi)

	budi.sayHello("Agus")
	eko.sayHello("Agus")
	joko.sayHello("Agus")
}

/*
=> Result from Struct when running the background:
	budi: {
		Name:    "Eko",
		Address: "Indonesia",
		Age:     30,
		sayHello: ("Agus") => {
			fmt.Println("Hello", <"Agus">, "my name is", <"Ekoo">)
		}
	}
*/

/*
# [ Struct Method ]
- Struct adalah tipe data seperti tipe data lainnya, dia bisa digunakan sebagai parameter untuk function
- Namun jika kita ingin menambahkan method ke dalam structs, sehingga seakan-akan sebuah struct memiliki function
- Method adalah function
*/
