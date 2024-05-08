package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

// Better way
func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	// forma #1
	employee := Employee{}
	fmt.Printf("%v\n", employee)

	// forma #2
	employee2 := Employee{
		id:       11,
		name:     "Jorge",
		vacation: true,
	}
	fmt.Printf("%v\n", employee2)

	// forma #3
	employee3 := new(Employee)
	fmt.Println(*employee3)
	employee3.id = 33
	employee3.name = "Eliecer"
	employee3.vacation = true
	fmt.Println(*employee3)

	// forma #4
	employee4 := NewEmployee(12, "Jorge Eliecer", false)
	fmt.Printf("%v\n", *employee4)
}
