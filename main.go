package main

import (
	"fmt"
	"os"
	Errores "tp1/diseno_alumnos/errores"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println(Errores.ErrorParametros{})
		return
	}

	archivo_lista := os.Args[1]
	archivo_padron := os.Args[2]

}
