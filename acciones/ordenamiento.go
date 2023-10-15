
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)



func OrdenarPadron(padron []int) []int {
	if len(padron) <= 1 {
        return dnis
    }
	pivot := padron[len(dnis)/2]
	var menores []int
	var iguales []int
	var mayores []int
	for _,dni := range padron {
		if dni < pivot {
			menores = append(menores,dni)
		} else if dni > pivot {
			mayores = append(mayores,dni)
		} else {
			iguales = append(iguales,dni)
		}
	}
	menores = OrdenarPadron(menores)
	mayores= OrdenarPadron(mayores)
	menores = append(menores,iguales)
	menores = append(menores,mayores)
	return menores
}

func EstaEnPadron(dni int, padron []int) bool {
	return estaEnPadron(dni,padron,0,len(padron)-1)
}

func estaEnPadron(dni int, padron []int, ini, fin) bool {
	if ini == fin {
		return false
	}
	medio := (ini + fin) / 2
	if padron[medio] == dni {
		return true
	} else if padron[medio] > dni {
		return estaEnPadron(dni, padron, 0, medio)
	} else {
		return estaEnPadron(dni, padron, medio+1, fin)
	} 
}
