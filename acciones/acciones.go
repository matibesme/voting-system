package acciones

import (
	"fmt"
	"strconv"
	"tdas/cola"
	"tp1/diseno_alumnos/errores"
	"tp1/diseno_alumnos/votos"
)

var CARGOS = []string{"Presidente", "Gobernador", "Intendente"}
var ENTRADA = []string{"ingresar", "votar", "deshacer", "fin-votar"}

func AccionIngresarVotante(dni_string string, cola cola.Cola[int], padron []votos.Votante) error {
	dni, err := strconv.Atoi(dni_string)

	if err != nil || dni <= 0 {
		return errores.DNIError{}
	} else {
		votante := EstaEnPadron(dni, padron)
		if votante == -1 {
			return errores.DNIFueraPadron{}
		} else {
			cola.Encolar(votante)
			return nil
		}
	}
}

func AccionVotar(entrada []string, cola cola.Cola[int], padron []votos.Votante, crear_partidos []votos.Partido, lista_partidos []string) error {

	if len(entrada) != 3 {
		return errores.ErrorParametros{}

	} else {

		cargo, err2 := verificoCargoAVotar(entrada[1])
		candidato, err := strconv.Atoi(entrada[2])

		if cola.EstaVacia() {
			return errores.FilaVacia{}

		} else if err2 != nil || cargo == 3 {
			return errores.ErrorTipoVoto{}
		} else if err != nil || !partidoValido(lista_partidos, candidato) {
			return errores.ErrorAlternativaInvalida{}
		} else {
			votar := (padron[cola.VerPrimero()]).Votar(cargo, candidato)

			if votar == nil {
				return nil
			} else {
				//error de votar
				cola.Desencolar()
				return votar
			}

		}
	}
}

func AccionDeshacer(cola cola.Cola[int], padrones []votos.Votante) error {

	if cola.EstaVacia() {
		return errores.FilaVacia{}
	} else {
		deshacer := padrones[cola.VerPrimero()].Deshacer()
		if deshacer == (errores.ErrorNoHayVotosAnteriores{}) {
			return deshacer
		} else if deshacer == nil {
			return nil
		} else {
			cola.Desencolar()
			return deshacer
		}
	}
}

func AccionFinVotante(cola cola.Cola[int], padrones []votos.Votante, partidos []votos.Partido, impugnados *int) error {
	if cola.EstaVacia() {
		return errores.FilaVacia{}
	} else {
		votante_actual := padrones[cola.Desencolar()]
		datos, err := votante_actual.FinVoto()
		if err != nil {
			return err
		} else if datos.Impugnado {
			*impugnados++
			return nil
		} else {

			for i := 0; i < 3; i++ {
				partidos[datos.VotoPorTipo[i]].VotadoPara(votos.TipoVoto(i))
			}
			return nil

		}
	}
}

func AccionResultadosElectorales(partidosCreados []votos.Partido, cola_voto cola.Cola[int], padrones []votos.Votante, impugnados *int) {
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
	if *impugnados == 1 {
		fmt.Println("Votos Impugnados:", *impugnados, "voto")
	} else {
		fmt.Println("Votos Impugnados:", *impugnados, "votos")
	}
}
