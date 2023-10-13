package votos

import (
	"tp1/diseno_alumnos/errores"
	TDAPila "tdas/pila"
)

type votanteImplementacion struct {
	dni int
	pilaVotos TDAPila.Pila[Voto]
	fraudulento bool
	voto Voto

}

func CrearVotante(dni int) Votante {
	return &votanteImplementacion{
		dni:  dni,
		votos: TDAPila.CrearPilaDinamica[Voto]()
		fraudulento: false,

	}
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.dni
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) error {
	
	if votante.fraudulento{
		return errores.ErrorVotanteFraudulento{dni: votante.dni}
	}

	votante.voto.VotoPorTipo[tipo]=alternativa
	votante.pilaVotos.Apilar(votante.voto)
	return nil

}

func (votante *votanteImplementacion) Deshacer() error {
	
	if votante.pilaVotos.EstaVacia(){
		return errores.ErrorNoHayVotosAnteriores{}
	}

	if votante.fraudulento{
		return errores.ErrorVotanteFraudulento{dni: votante.dni}
	}
	


	voto = votante.pilaVotos.Desapilar()
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {
	return Voto{}, nil
}
