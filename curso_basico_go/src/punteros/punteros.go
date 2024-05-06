package punteros

import "fmt"

type Pc struct {
	Ram   int
	Disk  int
	Brand string
}

func (myPC Pc) Ping() {
	fmt.Println(myPC.Brand, "Pong")
}

func (myPc *Pc) DuplicateRAM() {
	myPc.Ram = myPc.Ram * 2
}

/*
*
La estructura de datos " Struct " tiene un método llamado " String " , que podemos sobrescribir para personalizar la salida a consola de los datos del struct.
*
*/
func (myPc Pc) String() string {
	return fmt.Sprintf("tengo %d GB RAM , %d GB Disco y es una %s", myPc.Ram, myPc.Disk, myPc.Brand)
}
