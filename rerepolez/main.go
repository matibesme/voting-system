package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"tdas/cola"
	"tp1/acciones"
	"tp1/diseno_alumnos/errores"
	"tp1/diseno_alumnos/votos"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println(errores.ErrorParametros{})
		return
	}

	lista_partidos := partidosEnArchivo(os.Args[1])
	lista_padrones := padronesEnArchivo(os.Args[2])

	if lista_partidos == nil || lista_padrones == nil {
		fmt.Println(lista_partidos)
		fmt.Print(errores.ErrorLeerArchivo{})
		return
	}

	crear_partidos := CrearPartidos(lista_partidos)
	padrones := acciones.OrdenarPadron(lista_padrones)
	cola_votantes := cola.CrearColaEnlazada[int]()

	texto_ingresado := bufio.NewScanner(os.Stdin)
	for texto_ingresado.Scan() {
		entrada := strings.Split(texto_ingresado.Text(), " ")

		switch entrada[0] {
		case "ingresar":
			acciones.IngresarVotante(entrada[1], cola_votantes, lista_padrones)

		case "votar":
			acciones.Votar(entrada, cola_votantes, padrones, crear_partidos, lista_partidos)

		case "deshacer":
			acciones.Deshacer(cola_votantes, padrones)

		case "fin-votar":
			acciones.FinVoto(cola_votantes, padrones, crear_partidos)

		}

	}

	acciones.ResultadosElectorales(crear_partidos, cola_votantes, padrones)

}

//Aca una par de funciones extras
//dsp reutilizar las dos lecturas de archivo

func partidosEnArchivo(archivo_lista string) []string {
	var partidos []string
	archivo, err := os.Open(archivo_lista)

	if err != nil {
		return nil
	}
	fmt.Println("hola")
	defer archivo.Close()
	lector := bufio.NewScanner(archivo)

	for lector.Scan() {

		partidos = append(partidos, lector.Text())
	}

	return partidos
}

func padronesEnArchivo(archivo_lista string) []votos.Votante {
	// leer archivo de padrones
	var padron []votos.Votante
	archivo, err := os.Open(archivo_lista)
	if err != nil {
		return nil
	}

	defer archivo.Close()
	lector := bufio.NewScanner(archivo)
	for lector.Scan() {
		dni_en_num, err := strconv.Atoi(lector.Text())
		if err != nil {
			fmt.Println(errores.DNIError{})
		}
		padron = append(padron, votos.CrearVotante(dni_en_num))
	}
	return padron

}

func CrearPartidos(lista_partidos []string) []votos.Partido {

	var partidos []votos.Partido

	partidos = append(partidos, votos.CrearVotosEnBlanco())
	for _, partido := range lista_partidos {
		cargo_partidos := strings.Split(partido, ",")
		postulantes := [3]string{cargo_partidos[1], cargo_partidos[2], cargo_partidos[3]}
		partidos = append(partidos, votos.CrearPartido(cargo_partidos[0], postulantes))

	}

	return partidos

}
