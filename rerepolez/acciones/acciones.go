package acciones

import (
	"fmt"
	"rerepolez/diseno_alumnos/errores"
	"rerepolez/diseno_alumnos/votos"
	"strconv"
	"tdas/cola"
)

var CONTADOR_IMPUGNADOS = 0
var CARGOS = []string{"Presidente", "Gobernador", "Intendente"}

func AccionIngresarVotante(dni_string string, cola cola.Cola[int], padron []votos.Votante) {
	dni, err := strconv.Atoi(dni_string)

	if err != nil || dni <= 0 || len(dni_string) != 8 {
		fmt.Println(errores.DNIError{})
	} else {
		votante := EstaEnPadron(dni, padron)

		if votante == -1 {
			fmt.Println(errores.DNIFueraPadron{})
		} else {
			cola.Encolar(votante)
			fmt.Println(padron[votante].LeerDNI())
			fmt.Println("OK")
		}
	}
}

func AccionVotar(entrada []string, cola cola.Cola[int], padron []votos.Votante, crear_partidos []votos.Partido, lista_partidos []string) {
	//evaluo errores

	if len(entrada) != 3 {
		fmt.Println(errores.ErrorParametros{})

	} else {

		cargo, err2 := verificoCargoAVotar(entrada[1])
		candidato, err := strconv.Atoi(entrada[2])

		if cola.EstaVacia() {
			fmt.Println(errores.FilaVacia{})

		} else if err2 != nil || cargo == 3 {
			fmt.Println(errores.ErrorTipoVoto{})
		} else if err != nil || !partidoValido(lista_partidos, candidato) {
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
}

func AccionDeshacer(cola cola.Cola[int], padrones []votos.Votante) {

	if cola.EstaVacia() {
		fmt.Println(errores.FilaVacia{})
	} else {
		fmt.Println(padrones[cola.VerPrimero()])
		deshacer := padrones[cola.VerPrimero()].Deshacer()
		if deshacer == (errores.ErrorVotanteFraudulento{}) {
			fmt.Println(deshacer)
		} else if deshacer == (errores.ErrorNoHayVotosAnteriores{}) {
			fmt.Println(deshacer)
		} else {

			fmt.Println("OK")
		}
	}

}

func AccionFinVotante(cola cola.Cola[int], padrones []votos.Votante, partidos []votos.Partido) {
	if cola.EstaVacia() {
		fmt.Println(errores.FilaVacia{})
	} else {
		votante_actual := padrones[cola.VerPrimero()]
		datos, err := votante_actual.FinVoto()
		if err != nil {
			fmt.Println(err)
		} else if datos.Impugnado {
			CONTADOR_IMPUGNADOS++
		} else {
			fmt.Println("OK")
			for i := 0; i < 3; i++ {
				partidos[datos.VotoPorTipo[i]].VotadoPara(votos.TipoVoto(i))
			}
			cola.Desencolar()
		}
	}
}

func AccionResultadosElectorales(partidosCreados []votos.Partido, cola_voto cola.Cola[int], padrones []votos.Votante) {
	if !cola_voto.EstaVacia() {
		fmt.Println(errores.ErrorCiudadanosSinVotar{})
	}
	for tipo_voto := votos.PRESIDENTE; tipo_voto < votos.CANT_VOTACION; tipo_voto++ {
		fmt.Printf("%s:\n", CARGOS[tipo_voto])
		for _, partido := range partidosCreados {
			fmt.Println(partido.ObtenerResultado(tipo_voto))
		}
		fmt.Println()
	}
	fmt.Println("Votos Impugnados:", CONTADOR_IMPUGNADOS)
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
