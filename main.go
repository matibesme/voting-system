package main

import (
	"fmt"
	"os"
	"tp1/diseno_alumnos/errores"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println(errores.ErrorParametros{})
		return
	}

	archivo_lista := os.Args[1]
	archivo_padron := os.Args[2]

	if archivo_lista==nil || archivo_padron==nil{
		fmt.Print(errores.ErrorLeerArchivo{})
		return
	}



}

func partidoValido(){

}



func dniEnPadron{

}

func Part