package acciones

import (
	"tp1/diseno_alumnos/votos"
)

//Usamos el algoritmo de RadixSort y adentro se usa CountingSort
func OrdenarPadron(padron []votos.Votante) []votos.Votante {
	max := maximo(padron)

    for exp := 1; max/exp > 0; exp *= 10 {
        padron = countingSort(padron, exp)
    }

    return padron
}

func maximo(arr []votos.Votante) int {
    max := arr[0].LeerDNI()
    for i := 1; i < len(arr); i++ {
        dni := arr[i].LeerDNI()
        if dni > max {
            max = dni
        }
    }
    return max
}

func countingSort(arr []votos.Votante, exp int) []votos.Votante {
    largo := len(arr)
    salida := make([]votos.Votante, largo)
    frecuencias := make([]int, 10) 

    for i := 0; i < largo; i++ {
        indice := (arr[i].LeerDNI() / exp) % 10
        frecuencias[indice]++
    }

    for i := 1; i < 10; i++ {
        frecuencias[i] += frecuencias[i-1]
    }

    for i := largo - 1; i >= 0; i-- {
        indice := (arr[i].LeerDNI() / exp) % 10
        salida[frecuencias[indice]-1] = arr[i]
        frecuencias[indice]--
    }

    return salida
}

func EstaEnPadron(dni int, padron []votos.Votante) int {

	return estaEnPadron(dni, padron, 0, len(padron))
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
