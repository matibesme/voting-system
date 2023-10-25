package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tdas/cola"
	"tp1/acciones"
	"tp1/diseno_alumnos/errores"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println(errores.ErrorParametros{})
		return
	}

	lista_partidos := acciones.PartidosEnArchivo(os.Args[1])
	lista_padrones := acciones.PadronesEnArchivo(os.Args[2])

	if lista_partidos == nil || lista_padrones == nil {

		fmt.Println(errores.ErrorLeerArchivo{})
		return
	}

	crear_partidos := acciones.CrearPartidos(lista_partidos)
	padrones_ordenados := acciones.OrdenarPadron(lista_padrones)
	cola_votantes := cola.CrearColaEnlazada[int]()
	texto_ingresado := bufio.NewScanner(os.Stdin)

	for texto_ingresado.Scan() {
		texto_ingresado := strings.Split(texto_ingresado.Text(), " ")

		switch texto_ingresado[0] {
		case acciones.ENTRADA[0]:
			acciones.AccionIngresarVotante(texto_ingresado[1], cola_votantes, padrones_ordenados)

		case acciones.ENTRADA[1]:
			acciones.AccionVotar(texto_ingresado, cola_votantes, padrones_ordenados, crear_partidos, lista_partidos)

		case acciones.ENTRADA[2]:
			acciones.AccionDeshacer(cola_votantes, padrones_ordenados)

		case acciones.ENTRADA[3]:
			acciones.AccionFinVotante(cola_votantes, padrones_ordenados, crear_partidos)
		}

	}
	acciones.AccionResultadosElectorales(crear_partidos, cola_votantes, padrones_ordenados)

}
