package structs

import "fmt"

func Demo2() {
	c := customer{firstName: "Samed", lastName: "Kütahyalı", age: 21}
	c.save()
	c.update()
}

type customer struct {
	firstName string
	lastName  string
	age       int
}

func (c customer) save() {
	fmt.Println("Eklendi. ", c)
}

func (c customer) update() {
	fmt.Println("Güncellendi. ", c)
}
