package acciones

import (
	"bufio"
	"fmt"
	"os"
	"rerepolez/diseno_alumnos/errores"
	"rerepolez/diseno_alumnos/votos"
	"strconv"
	"strings"
)

func PartidosEnArchivo(archivo_lista string) []string {
	var partidos []string
	archivo, err := os.Open(archivo_lista)

	if err != nil {
		return nil
	}
	defer archivo.Close()
	lector := bufio.NewScanner(archivo)

	for lector.Scan() {
		partidos = append(partidos, lector.Text())
	}

	return partidos
}

func PadronesEnArchivo(archivo_lista string) []votos.Votante {
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

func verificoCargoAVotar(cargo string) (votos.TipoVoto, error) {
	switch cargo {
	case CARGOS[0]:
		return votos.PRESIDENTE, nil
	case CARGOS[1]:
		return votos.GOBERNADOR, nil
	case CARGOS[2]:
		return votos.INTENDENTE, nil
	default:
		return votos.CANT_VOTACION, errores.ErrorTipoVoto{}
	}
}

func partidoValido(partidos []string, cantPartido int) bool {
	return len(partidos) >= cantPartido
}
