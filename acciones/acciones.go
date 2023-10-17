package acciones

import (
	"fmt"
	"strconv"
	"tdas/cola"
	"tp1/diseno_alumnos/errores"
	"tp1/diseno_alumnos/votos"
)

var contador_impugnados = 0

func IngresarVotante(dni_string string, cola cola.Cola[int], padron []votos.Votante) {
	dni, err := strconv.Atoi(dni_string)

	if err != nil || dni <= 0 || len(dni_string) != 8 {
		fmt.Println(errores.DNIError{})
	} else {
		votante := EstaEnPadron(dni, padron)
		if votante == -1 {
			fmt.Println(errores.DNIFueraPadron{})
		} else {
			cola.Encolar(votante)
			fmt.Println("OK")
		}
	}
}

func Votar(entrada []string, cola cola.Cola[int], padron []votos.Votante, crear_partidos []votos.Partido, lista_partidos []string) {
	//evaluo errores
	cargo, err2 := verificoCargoAVotar(entrada[1])
	candidato, err := strconv.Atoi(entrada[2])

	if cola.EstaVacia() {
		fmt.Println(errores.FilaVacia{})
	} else if err2 != nil || cargo == 3 {
		fmt.Println(errores.ErrorTipoVoto{})
	} else if err != nil || !partidoValido(lista_partidos, candidato) {
		//CONDICION
		fmt.Println(errores.ErrorAlternativaInvalida{})
	} else {
		votar := (padron[cola.VerPrimero()]).Votar(cargo, candidato)
		if votar == nil {
			fmt.Println("OK")
		} else {
			//error de votar
			fmt.Println(votar)
			cola.Desencolar()
		}

	}

}

func Deshacer(cola cola.Cola[int], padrones []votos.Votante) {

	if cola.EstaVacia() {
		fmt.Println(errores.FilaVacia{})
	} else {
		test := padrones[cola.VerPrimero()].Deshacer()
		if test == nil {
			fmt.Println("OK")
		} else if test == (errores.ErrorNoHayVotosAnteriores{}) {
			fmt.Println(test)
		} else {
			fmt.Println(test)
			cola.Desencolar()
		}

	}

}

func FinVoto(cola cola.Cola[int], padrones []votos.Votante, partidos []votos.Partido) {
	if cola.EstaVacia() {
		fmt.Println(errores.FilaVacia{})
	} else {
		votante_actual := padrones[cola.VerPrimero()]
		datos, err := votante_actual.FinVoto()
		if err != nil {
			fmt.Println(err)
		} else if datos.Impugnado {
			contador_impugnados++
		} else {
			fmt.Println("OK")
			for i := 0; i < 3; i++ {
				partidos[datos.VotoPorTipo[i]].VotadoPara(votos.TipoVoto(i))
			}
		}
	}
}

func ResultadosElectorales(partidosCreados []votos.Partido, cola_voto cola.Cola[int], padrones []votos.Votante) {
	if !cola_voto.EstaVacia() {
		fmt.Println(errores.ErrorCiudadanosSinVotar{})
	}
	for tipo_voto := votos.PRESIDENTE; tipo_voto < votos.CANT_VOTACION; tipo_voto++ {
		for _, partido := range partidosCreados {
			fmt.Println(partido.ObtenerResultado(tipo_voto))
		}
		fmt.Println()
	}
	fmt.Println("Votos Impugnados:", contador_impugnados)
}

func verificoCargoAVotar(cargo string) (votos.TipoVoto, error) {
	switch cargo {
	case "Presidente":
		return votos.PRESIDENTE, nil
	case "Gobernador":
		return votos.GOBERNADOR, nil
	case "Intendente":
		return votos.INTENDENTE, nil
	default:
		return votos.CANT_VOTACION, errores.ErrorTipoVoto{}
	}
}

func partidoValido(partidos []string, cantPartido int) bool {
	return len(partidos) >= cantPartido
}
