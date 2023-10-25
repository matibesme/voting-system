package acciones

import (
	"fmt"
	"strconv"
	"tdas/cola"
	"tp1/diseno_alumnos/errores"
	"tp1/diseno_alumnos/votos"
)

var CONTADOR_IMPUGNADOS = 0
var CARGOS = []string{"Presidente", "Gobernador", "Intendente"}
var ENTRADA = []string{"ingresar", "votar", "deshacer", "fin-votar"}


func AccionIngresarVotante(dni_string string, cola cola.Cola[int], padron []votos.Votante) {
	dni, err := strconv.Atoi(dni_string)

	if err != nil || dni <= 0 {
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

func AccionVotar(entrada []string, cola cola.Cola[int], padron []votos.Votante, crear_partidos []votos.Partido, lista_partidos []string) {

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
		deshacer := padrones[cola.VerPrimero()].Deshacer()
		if deshacer == (errores.ErrorNoHayVotosAnteriores{}) {
			fmt.Println(deshacer)
		} else if deshacer == nil {
			fmt.Println("OK")
		} else {
			fmt.Println(deshacer)
			cola.Desencolar()

		}
	}

}

func AccionFinVotante(cola cola.Cola[int], padrones []votos.Votante, partidos []votos.Partido) {
	if cola.EstaVacia() {
		fmt.Println(errores.FilaVacia{})
	} else {
		votante_actual := padrones[cola.Desencolar()]
		datos, err := votante_actual.FinVoto()
		if err != nil {
			fmt.Println(err)
		} else if datos.Impugnado {
			CONTADOR_IMPUGNADOS++
			fmt.Println("OK")
		} else {
			fmt.Println("OK")
			for i := 0; i < 3; i++ {
				partidos[datos.VotoPorTipo[i]].VotadoPara(votos.TipoVoto(i))
			}

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
	if CONTADOR_IMPUGNADOS == 1 {
		fmt.Println("Votos Impugnados:", CONTADOR_IMPUGNADOS, "voto")
	} else {
		fmt.Println("Votos Impugnados:", CONTADOR_IMPUGNADOS, "votos")
	}
}
