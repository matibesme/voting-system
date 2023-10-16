package acciones

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tp1/diseno_alumnos/votos"
)



func OrdenarPadron(padron []votos.Votante) []votos.Votante {
	if len(padron) <= 1 {
        return dnis
    }
	pivot := padron[len(dnis)/2].LeerDNI()
	var menores,iguales,mayores []votos.Votante

	for i := 0; i < len(padron); i++ {
		if padron[i].LeerDNI() < pivot {
			menores = append(menores,padron[i])
		} else if padron[i].LeerDNI() > pivot {
			mayores = append(mayores,padron[i])
		} else {
			iguales = append(iguales,padron[i])
		}
	}

	menores = OrdenarPadron(menores)
	mayores= OrdenarPadron(mayores)
	menores = append(menores,iguales)
	menores = append(menores,mayores)
	return menores
}

func EstaEnPadron(dni int, padron []votos.Voto) votos.Votante {
	return estaEnPadron(dni,padron,0,len(padron)-1)
}

func estaEnPadron(dni int, padron []int, ini, fin)  {
	if ini == fin {
		return false,nil
	}
	medio := (ini + fin) / 2
	if padron[medio].LeerDNI() == dni {
		return true,padron[medio]
	} else if padron[medio].LeerDNI() > dni {
		return estaEnPadron(dni, padron, 0, medio)
	} else {
		return estaEnPadron(dni, padron, medio+1, fin)
	} 
}
