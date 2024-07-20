package main

import "fmt"

// Contract Interface
type HasName interface {
	GetName() string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

// Implement interface with Struct.
type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

// Implement from  interface HashName because it a have GetName method.
func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "Eko"}
	SayHello(person)

	animal := Animal{Name: "Kucing"}
	SayHello(animal)
}

/*
[Interface]
- Interface adalah tipe data Abstract, dia tidak memiliki implementasi langsung.
- Sebuah interface berisikan definisi-definisi method.
- Biasanya interface digunakan sebagai kontrak.

[#Implementasi Interface]
- Setiap tipe data yang sesuai dengan kontrak interface, secara otomatis dianggap sebagai interface tersebut.
- Sehingga kita tidak perlu mengimplementasikan interface secara manual.
- Hal ini agak berbeda dengan bahasa pemrograman lain yang ketika membuat interface, kita harus menyebutkan secara eksplisit akan menggunakan interface mana.
*/
