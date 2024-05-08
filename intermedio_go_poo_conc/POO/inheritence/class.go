package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	// manera anonima
	Person
	Employee
}

/* func GetMessage(p Person) {
	fmt.Printf("%s with age %d\n", p.name, p.age)
} */

func main() {
	// Composicion
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Jorge Eliecer"
	ftEmployee.age = 33
	ftEmployee.id = 1020
	fmt.Printf("%v\n", ftEmployee)

	//GetMessage(ftEmployee)
}
