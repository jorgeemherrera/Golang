package main

import (
	"log"
	"os"
)

type CustomLogger struct {
	FilePath string
}

func (l *CustomLogger) Write(p []byte) (n int, err error) {
	file, err := os.OpenFile(l.FilePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	file.Write(p)

	return 0, nil

}

func main() {
	//SetOutput(io.Writer): definir la salida - io.Writer
	miLogger := new(CustomLogger)
	miLogger.FilePath = "webserver.log"
	log.SetOutput(miLogger)

	for i := 1; i <= 20; i++ {
		log.Printf("Error en la linea %v", i)
	}

	//Varios logger
	miSegundoLogger := new(CustomLogger)
	miSegundoLogger.FilePath = "system.log"

	log.SetOutput(miSegundoLogger)

	for i := 1; i < 20; i++ {
		log.Printf("Hay un error en la linea %v", i)
		log.Fatalln("Errooor!, close the program!!")
	}
}
