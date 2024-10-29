package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) getName() string {
	return p.name
}

func (p Person) getAge() int {
	return p.age
}

func main() {
	p1 := Person{"Peter Griffin", 48}

	p2 := Person{
		name: "Louis",
		age:  43,
	}

	fmt.Printf("name: %s, age: %d\n", p1.getName(), p1.getAge())
	fmt.Printf("name: %s, age: %d\n", p2.getName(), p2.getAge())
}
