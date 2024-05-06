package mypackage

import "fmt"

// CarPublic - car con acceso publico
type CarPublic struct {
	Brand string
	Year  int
}

// CarPublic - car con acceso privado
type carPrivate struct {
	brand string
	year  int
}

func PrintMessage(text string) {
	fmt.Println(text)
}

// funcion con acceso privado
func printMessageP(text string) {
	fmt.Println(text)
}
