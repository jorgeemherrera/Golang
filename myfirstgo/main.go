package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	//comments

	/* comments with
	double lines */

	// Variables
	var myNumber int = 9
	fmt.Println(myNumber)

	// := skip the var keyword, type inference
	otherNumber := 8
	fmt.Println(otherNumber)

	// Declare arrays
	var myArray []int = []int{1, 2, 3}
	fmt.Println(myArray)

	// Indexing
	fmt.Println(myArray[2])

	// Declare Structures (Objects)
	type myRecord struct {
		Name string
		Age  int
	}

	var oneStudent myRecord = myRecord{
		Name: "Jorge",
		Age:  33,
	}

	fmt.Println("student: ", oneStudent)

	// Conditions IF - ELSE
	if 2 > 5 {
		fmt.Println("this  is not true")
	}

	if 2 > 5 {
		fmt.Println("always false")
	} else if 4 > 6 {
		fmt.Println("two is not greater than 5")
	} else {
		fmt.Println("this is not true else")
	}

	// For Loop
	for i := 0; i < 10; i++ {
		fmt.Println("the values of i are ===: ", i)
	}

	i := 1
	for i < 10 {
		fmt.Println("the values of i are: ", i)
		i++
	}
}
