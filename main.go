package main

import "fmt"

type animal interface {
	// methods
	speak() string
}

func speaker(s animal) {
	println(s.speak())
}

type cat struct {
	name string
}

func (c *cat) speak() string {
	return fmt.Sprintf("[%s] meow", c.name)
}

type dog struct {
	name string
}

func (d *dog) speak() string {
	return fmt.Sprintf("[%s] woof", d.name)
}

func main() {
	println("main is running example of interface in GO")

	mitzi := &cat{name: "Mitzi"}
	rover := &dog{name: "Rover"}

	animals := []animal{mitzi, rover}
	for _, a := range animals {
		speaker(a)
	}
}
