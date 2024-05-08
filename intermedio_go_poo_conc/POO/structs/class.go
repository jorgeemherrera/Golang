package main

import "fmt"

// Structs vs. clases
type Employee struct {
	id   int
	name string
}

// Métodos y funciones
// receiver functions

// Setter
func (emp *Employee) SetId(id int) {
	emp.id = id
}

func (emp *Employee) SetName(name string) {
	emp.name = name
}

// Getter
func (emp Employee) GetId() int {
	return emp.id
}

func (emp Employee) GetName() string {
	return emp.name
}

func main() {
	// Instancia del empleado
	employee := Employee{}
	fmt.Printf("%v", employee)
	employee.id = 1
	employee.name = "jorge"
	fmt.Printf("%v", employee)
	employee.SetId(5)
	fmt.Printf("%v", employee)
	employee.SetName("Eliecer")
	fmt.Printf("%v", employee)
	fmt.Println(employee.GetId())
	fmt.Println(employee.GetName())
}
