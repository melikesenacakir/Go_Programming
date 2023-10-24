package structs

import "fmt"

func Demo2() {
	c := customer{firstname: "Melike", lastname: "Çakır", age: 22}
	c.save()
}

type customer struct {
	firstname string
	lastname  string
	age       int
}

func (c customer) save() { //bu struct metodudur
	fmt.Println("Eklendi: ", c.firstname)
}
