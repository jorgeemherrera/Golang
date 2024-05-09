package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("Este es un fmt.Println")

	// Tipo info
	log.Println("Este es un log.Println")

	// Logs infinitos
	i := 1

	for {
		err := fmt.Errorf(fmt.Sprintf("este es el error %v", i))
		log.Println(err)
		i++
		time.Sleep(time.Second * 1)
	}

	// Tipo fatal
	/* log.Fatalln("Errooor!, cierra el programa")

	log.Println("Esto ya no se ejecuta/sucede") */

}
