package acciones

import (
	"fmt"
	"strconv"
	"tdas/cola"
	"tp1/acciones"
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

func Votar(entrada []string, cola.Cola[int],padron []votos.Votante,crear_partidos []votos.Partido,lista_partidos []string) {
	//evaluo errores
	cargo=verificoCargoAVotar(entrada[1])

	if cola.EstaVacia(){
		fmt.Println(errores.FilaVacia{})
	} else if cargo=="INVALIDO"{
		fmt.Println(errores.ErrorTipoVoto{})
	}else if  {
		//CONDICION
		fmt.Println(errores.ErrorAlternativaInvalida{})
	}
	else{
		


	}

}

func Deshacer(cola cola.Cola[int],padrones []votos.Votante) {

	if cola.EstaVacia(){
		fmt.Println(errores.FilaVacia{})
	} else{
		test:=padron[cola.CrearColaEnlazada().VerPrimero()].Deshacer()
		if test ==nil{
			fmt.PrintIn()
		}else if{
		
		}else{

		}
		
	}



}

func FinVoto(cola.Cola[int],padrones []votos.Votante,crear_partidos []votos.Partido) {
	if cola.EstaVacia(){
		fmt.Println(errores.FilaVacia{})
	} else{
		votante_actual:=padrones[cola.CrearColaEnlazada().Desencolar()]
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

func Fin(cola.Cola[int]) {
	if cola.EstaVacia(){
		fmt.Println(errores.FilaVacia{})
	} 

	for i := 0; i < 3; i++{

	}	
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
