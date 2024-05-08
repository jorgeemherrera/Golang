package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var x int
	x = 8
	y := 7
	fmt.Println(x, y)

	// Errores implicitos
	myValue, err := strconv.ParseInt("8", 0, 64)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("The value of myValue is %d\n", myValue)
	}

	//Map
	m := make(map[string]int)
	m["Key"] = 6
	fmt.Println(m["Key"])

	//Slice
	sl := []int{1, 2, 3}
	for index, value := range sl {
		fmt.Println(index)
		fmt.Println(value)
	}
	sl = append(sl, 16)
	for index, value := range sl {
		fmt.Println(index)
		fmt.Println(value)
	}
	//Channel
	//c := make(chan int)
	//go doSomething(c)
	//<-c //Espera por el mensaje en el canal c
	g := 25
	fmt.Println(g)
	h := &g // solo la referencia
	fmt.Println(h)
	fmt.Println(*h) // acceder al valor
}

func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
	c <- 1
}
