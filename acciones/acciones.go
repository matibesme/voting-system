package acciones

import (
	"fmt"
	"strconv"
	"tdas/cola"
	"tp1/diseno_alumnos/errores"
	"tp1/diseno_alumnos/votos"
)

func IngresarVotante(dni_string string,cola cola.Cola[int], padron []votos.Votante) {

	dni, err := strconv.Atoi(numeroStr)

	if dni <= 0 || len(dni_str) != 8 {
		fmt.Println(errores.DNIError{})
	} else {
		//buscar dni
		votante := busquedaDni(padron, dni)

		if votante == nil {
			fmt.Println(errores.DNIFueraPadron{})
		} else {
			cola.Encolar(votante)
			fmt.Printf("OK")
		}
	}

}

func Votar(entrada []string, cola.Cola[int]) {
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

func Deshacer(cola.Cola[int]) {

	if cola.EstaVacia(){
		fmt.Println(errores.FilaVacia{})
	} else{

		
	}



}

func FinVotar() {}

func Fin() {}



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
