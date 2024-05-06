package main

import (
	pk "curso_golang_platzi/src/mypackage"
	punt "curso_golang_platzi/src/punteros"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/labstack/echo"
)

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArguments(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

// Structs - la forma de hacer clases en Go
type car struct {
	brand string
	year  int
}

type figures2D interface {
	area() float64
}

type square struct {
	base float64
}

type rectangle struct {
	length, width float64
}

// Methods
func (c square) area() float64 {
	return c.base * c.base
}

// Methods
func (r rectangle) area() float64 {
	return r.length * r.width
}

func calculateArea(f figures2D) {
	fmt.Println("Area:", f.area())
}

func isPalindrom(text string) {
	var textReverse string
	text = strings.ToLower(text)
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}
	if text == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}

// GoRutines
func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

// Channel
func sayChan(text string, c chan<- string) {
	c <- text
}

func messageChan(text string, c chan string) {
	c <- text
}

func main() {
	fmt.Println("Hello World")
	fmt.Printf("2, %T", 2)

	// Declaracion de Constantes
	const PI float64 = 3.14
	const PI2 = 3.144

	fmt.Println("PI", PI)
	fmt.Println("PI2", PI2)

	// Declaracion de variables enteras
	base := 12
	var altura int = 8
	var area int

	fmt.Println(base, altura, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d) // 0 0  false

	// Area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("El area del cuadrado es: ", areaCuadrado)

	//Operadores aritméticos

	//Suma
	x := 10
	y := 50
	result := x + y
	fmt.Println("Suma:", result)

	//Resta
	result = y - x
	fmt.Println("Resta:", result)

	//Mutiplicacion
	z := 9
	result = x * z
	fmt.Println("Multiplicación:", result)

	//Division
	result = y / x
	fmt.Println("División:", result)

	// Modulo
	result = y % x
	fmt.Println("Modulo:", result)

	// Incrementar
	x++
	fmt.Println("Incrementar:", x)

	//Decrementar
	x--
	fmt.Println("Decrementar", x)

	//Area de Rectangulo, trapecio y circulo

	//Rectangulo
	largo := 10
	ancho := 20

	areaRectangulo := largo * ancho
	fmt.Println("El area del rectángulo es:", areaRectangulo)

	//Trapecio
	base = 3
	baseAbajo := 2
	altura = 2
	areaTrapecio := altura * (base + baseAbajo) / 2
	fmt.Println("El area del Trapecio es:", areaTrapecio)

	//Circulo
	radio := 5
	areaCirculo := math.Pi * (math.Pow(float64(radio), 2))
	fmt.Println("La Area del  Circulo es: ", areaCirculo)

	//Valores Primitivos

	//Numeros enteros

	//int = Depende del OS (32 o 64 bits)
	//int8 = 8bits = -128 a 127
	//int16 = 16bits = -2^15 a 2^15-1
	//int32 = 32bits = -2^31 a 2^31-1
	//int64 = 64bits = -2^63 a 2^63-1

	//Optimizar memoria cuando sabemos que el dato simpre va ser positivo

	//uint = Depende del OS (32 o 64 bits)
	//uint8 = 8bits = 0 a 127
	//uint16 = 16bits = 0 a 2^15-1
	//uint32 = 32bits = 0 a 2^31-1
	//uint64 = 64bits = 0 a 2^63-1

	//numeros decimales

	// float32 = 32 bits = +/- 1.18e^-38 +/- -3.4e^38
	// float64 = 64 bits = +/- 2.23e^-308 +/- -1.8e^308

	//textos y booleanos

	//string = ""
	//bool = true or false

	//numeros complejos

	//Complex64 = Real e Imaginario float32
	//Complex128 = Real e Imaginario float64
	//Ejemplo : c:=10 + 8i

	//Declaracion de variables
	helloMessage := "Hello"
	worldMessage := "World!"

	//Println - salto de linea
	fmt.Println(helloMessage, worldMessage) //Imprime Hello World!

	//Printf
	// %s --> string
	// %v --> si no sabes el dato
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)

	//Sprintf
	message := fmt.Sprintf("%s tiene mas de %d cursos", nombre, cursos)
	fmt.Println(message)

	//Tipo de dato de una variable
	// %T
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)

	fmt.Println("-----------------------")

	// funciones y funciones anonimas
	normalFunction("Hello World")
	tripleArguments(1, 2, "hola")

	value := returnValue(2)
	fmt.Println("Valor devuelve:", value)

	value1, value2 := doubleReturn(2)
	fmt.Println("Values: ", value1, value2)

	//Ciclos - solo tiene cilco FOR

	// for condicional
	for i := 0; i <= 10; i++ {
		fmt.Println("i:", i)
	}

	// for while
	counter := 0
	for counter < 10 {
		fmt.Println("counter: ", counter)
		counter++
	}

	// For forever
	/* counterForever := 0
	for {
		fmt.Println("counterForever", counterForever)
		counterForever++
	} */

	//Sentencias condicionales

	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// AND &&
	if valor1 == 1 && valor2 == 2 {
		println("es verdad")
	}

	// OR ||
	if valor1 == 0 || valor2 == 2 {
		println("es verdadera")
	}

	//Convertir string a numero
	value, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal(err)
	} else {
		println("Value", value)
	}

	//Switch
	//modulo := 5 % 2;
	switch modulo := 5 % 2; modulo {
	case 0:
		println("es par")
	default:
		println("es impar")
	}

	// Sin condicion
	value = 200
	switch {
	case value > 100:
		println("es mayor a 100")
	case value < 0:
		println("es menor a cero")
	default:
		println("No condicion")
	}

	// El uso de los keywords defer, break y continue

	//Defer
	//se utiliza para retrasar la ejecución de una función hasta que la función que la contiene haya terminado su ejecución
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	// Continue & Break
	for i := 0; i <= 11; i++ {
		println(i)

		//Continue
		if i == 2 {
			println("es dos")
			continue
		}

		//Break
		if i == 8 {
			println("Break")
			break
		}
	}

	// Array & Slices
	// len: legnth
	// cap: capacity
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	//Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Slicing
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	//Append
	slice = append(slice, 7)
	fmt.Println(slice)

	//Append new list
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)

	fmt.Println(slice)

	//Recorrido de Slices con Range
	sliceRange := []string{"Hola", "que", "pasa"}

	for i, valor := range sliceRange {
		fmt.Println(i, valor)
	}

	//Palindromo
	isPalindrom("Ama")

	// Llave valor con Maps
	m := make(map[string]int)
	m["Jose"] = 14
	m["Juan"] = 12

	fmt.Println(m)

	// Recorrer un map
	for i, valor := range m {
		fmt.Println(i, valor)
	}

	// Encontrar un valor
	valueFind, ok := m["Jose"]
	fmt.Println(valueFind, ok)

	// Structs
	myCar := car{brand: "Ford", year: 2024}
	fmt.Println(myCar)

	// instanciar objeto con un propiedad vacia
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)

	// Modificadores de Acceso
	// Mayuscula: publico
	// Minuscula: privado
	var myCarPublic pk.CarPublic
	myCarPublic.Brand = "Ferrari"
	myCarPublic.Year = 2020
	fmt.Println(myCarPublic)

	pk.PrintMessage("Hola Platzi")

	// Structs y Punteros
	punteroA := 50
	punteroB := &punteroA // direccion de memoria

	fmt.Println(punteroB)
	fmt.Println(*punteroB) // apuntando a la misma direccion de memoria

	*punteroB = 100
	fmt.Println(punteroA)

	myPc := punt.Pc{Ram: 16, Disk: 200, Brand: "msi"}
	fmt.Println(myPc)

	myPc.Ping()

	//Punteros
	fmt.Println(myPc)
	myPc.DuplicateRAM()
	fmt.Println(myPc)
	myPc.DuplicateRAM()
	fmt.Println(myPc)

	// Interfaces y listas de interfaces
	mySquare := square{base: 2}
	myRectangle := rectangle{length: 2, width: 4}

	calculateArea(mySquare)
	calculateArea(myRectangle)

	// Lista de interfaces
	myInterface := []interface{}{"Hola", 12, 4.90}
	fmt.Println(myInterface...)

	//Primer contacto con las Goroutines
	//sync.WaitGroup: agrupa un grupo de gorutines y la va liberando poco a poco
	var wg sync.WaitGroup
	fmt.Println("Hello")
	wg.Add(1)                // agregar un gorutine
	go say("gorutines", &wg) // puntero de wg

	wg.Wait() // La función wait() bloquea la rutina principal hasta que todas las demás rutinas del grupo hayan terminado. -- wg.Done()

	go func(text string) {
		fmt.Println(text)
	}("Adios")

	time.Sleep(time.Second * 1)

	//Channels: La forma de organizar las goroutines
	channel := make(chan string, 1)

	fmt.Println("Hello channel")

	go sayChan("Bye", channel)

	fmt.Println(<-channel)

	//Range, Close y Select en channels
	chann := make(chan string, 2)
	chann <- "Mensaje 1"
	chann <- "Mensaje 2"

	fmt.Println(len(chann), cap(chann))

	//Close
	close(chann)

	//chann <- "Mensaje 3"

	// Range
	for message := range chann {
		fmt.Println(message)
	}

	//Select
	email1 := make(chan string)
	email2 := make(chan string)

	go messageChan("mensaje 1", email1)
	go messageChan("mensaje 2", email2)

	for i := 0; i < 2; i++ {
		select {
		case mensaje1 := <-email1:
			fmt.Println("Recibido el correo electronico: ", mensaje1)
		case mensaje2 := <-email2:
			fmt.Println("Recibido el correo electrónico: ", mensaje2)
		}
	}

	//Go get: El manejador de paquetes

	//instanciar ruta
	e := echo.New()
	//Ruta
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	e.Logger.Fatal(e.Start(":1323"))

}
