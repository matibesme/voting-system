package acciones

import (
	"rerepolez/diseno_alumnos/votos"
)

func OrdenarPadron(padron []votos.Votante) []votos.Votante {
	if len(padron) <= 1 {
		return padron
	}
	pivot := padron[len(padron)/2].LeerDNI()
	var menores, iguales, mayores []votos.Votante

	for i := 0; i < len(padron); i++ {
		if padron[i].LeerDNI() < pivot {
			menores = append(menores, padron[i])
		} else if padron[i].LeerDNI() > pivot {
			mayores = append(mayores, padron[i])
		} else {
			iguales = append(iguales, padron[i])
		}
	}

	menores = OrdenarPadron(menores)
	mayores = OrdenarPadron(mayores)
	for _, votante := range iguales {
		menores = append(menores, votante)
	}
	for _, votante := range mayores {
		menores = append(menores, votante)
	}

	return menores
}

func EstaEnPadron(dni int, padron []votos.Votante) int {

	return estaEnPadron(dni, padron, 0, len(padron)-1)
}

func estaEnPadron(dni int, padron []votos.Votante, ini, fin int) int {
	if ini == fin {
		return -1
	}
	medio := (ini + fin) / 2
	if padron[medio].LeerDNI() == dni {
		return medio
	} else if padron[medio].LeerDNI() > dni {
		return estaEnPadron(dni, padron, 0, medio)
	} else {
		return estaEnPadron(dni, padron, medio+1, fin)
	}
}
