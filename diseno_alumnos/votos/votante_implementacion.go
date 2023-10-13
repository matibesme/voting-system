package votos

import (
	"tp1/diseno_alumnos/errores"
	TDAPila "tdas/pila"
)

type votanteImplementacion struct {
	dni int
	pilaVotos TDAPila.Pila[]
	fraudulento bool

}

func CrearVotante(dni int) Votante {
	return votanteImplementacion{
		dni:  dni,
		votos: TDAPila.CrearPilaDinamica[]()
		fraudulento: false,

	}
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.dni
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) error {
	
}

func (votante *votanteImplementacion) Deshacer() error {
	
	if votante.pilaVotos.EstaVacia(){
		return errores.ErrorNoHayVotosAnteriores{}
	}
	voto = votante.pilaVotos.Desapilar()
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {
	return Voto{}, nil
}
