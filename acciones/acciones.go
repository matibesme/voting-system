package acciones

import (
	"fmt"
	"strconv"
	"tdas/cola"

	"tp1/diseno_alumnos/errores"
	"tp1/diseno_alumnos/votos"
)

func IngresarVotante(dni_string string,cola cola.Cola[int], padron []votos.Votante) {
	//convierte string a int
	dni, err := strconv.Atoi(numeroStr)


	if dni <= 0 || len(dni_str) != 8 {
		fmt.Println(errores.DNIError{})
	} else {
		//buscar dni
		valor,votante := acciones.EstaEnPadron(dni, padron)

		if valor == false {
			fmt.Println(errores.DNIFueraPadron{})
		} else {
			cola.Encolar(votante)
			fmt.Printf("OK")
		}
	}

}

func Votar(entrada []string,cola cola.Cola[int],padron []votos.Votante,crear_partidos []votos.Partido,lista_partidos []string) {
	//evaluo errores
	cargo := verificoCargoAVotar(entrada[1])

	if cola.EstaVacia(){
		fmt.Println(errores.FilaVacia{})
	} else if cargo == "INVALIDO"{
		fmt.Println(errores.ErrorTipoVoto{})
	}else if  {
		//CONDICION
		fmt.Println(errores.ErrorAlternativaInvalida{})
	}
	else{
		votar:=(padropadron[cola.VerPrimero()]).Votar()
		if votar==nil{
			fmt.Println("OK")
		}else{
			//error de votar
			fmt.Println(votar)
			cola.Desencolar()
		}

	}

}

func Deshacer(cola cola.Cola[int],padrones []votos.Votante) {

	if cola.EstaVacia(){
		fmt.Println(errores.FilaVacia{})
	} else{
		test:=padron[cola.VerPrimero()].Deshacer()
		if test == nil{
			fmt.Println("OK")
		}else if test == errores.ErrorNoHayVotosAnteriores{} {
			fmt.Println(test)
		}else{
			fmt.Println(test)
			cola.Desencolar()
		}
		
	}



}

func FinVoto(cola cola.Cola[int],padrones []votos.Votante,crear_partidos []votos.Partido) {
	if cola.EstaVacia(){
		fmt.Println(errores.FilaVacia{})
	} else{
		votante_actual:=padrones[cola.Desencolar()]
		datos,erL:=votante.FinVoto()
		if err!=nil{
			fmt.Println((err))
		} else{
			if datos.Impugnado{
				crear_partidos[0].VotadoPara(votos.TipoVoto())
				fmt.Println("OK")
			}else{
				fmt.Println("OK")
			}
		}
		
	}


}

func ResultadosElectorales(partidosCreados []votos.Partido, cola_voto cola.Cola[int], padrones []votos.Votante) {
	if !cola_voto.EstaVacia() {
		fmt.Println(errores.ErrorCiudadanosSinVotar{})
	}
	for i := 0; i < 3; i++ {
		tipo_voto := votos.TipoVoto(i)
		for j := 0; j < len(partidosCreados); j++ {
			fmt.Println()
		}
		fmt.Println()
	}
	fmt.Println("Votos Impugnados:", partidosCreados[0].ObtenerResultado())
}








func verificoCargoAVotar(cargo string) votos.TipoVoto{
	switch cargo {
	case "Presidente":
		return votos.PRESIDENTE

	case "Gobernador":
		return votos.GOBERNADOR

	case "Intendente":
		return votos.INTENDENTE

	default: 
		return "INVALIDO"

	}

}
