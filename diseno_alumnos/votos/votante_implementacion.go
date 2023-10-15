package votos

import (
	TDAPila "tdas/pila"
	"tp1/diseno_alumnos/errores"
)

type votanteImplementacion struct {
	dni         int
	pilaVotos   TDAPila.Pila[Voto]
	fraudulento bool
	voto       Voto
}

func CrearVotante(dni int) Votante {
	return &votanteImplementacion{
		dni:         dni,
		pilaVotos:   TDAPila.CrearPilaDinamica[Voto](),
		fraudulento: false,
	}
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.dni
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) error {

	if votante.fraudulento {
		return errores.ErrorVotanteFraudulento{Dni: votante.dni}
	}
	if alternativa == LISTA_IMPUGNA {
		votante.voto.Impugnado = true
	} else {
		votante.voto.VotoPorTipo[tipo] = alternativa
	}
	votante.pilaVotos.Apilar(votante.voto)
	return nil

}

func (votante *votanteImplementacion) Deshacer() error {

	if votante.fraudulento {
		return errores.ErrorVotanteFraudulento{Dni: votante.dni}
	}

	if votante.pilaVotos.EstaVacia() {
		return errores.ErrorNoHayVotosAnteriores{}
	}

	votante.pilaVotos.Desapilar()

	return nil
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {
	if votante.fraudulento {
		return votante.voto, errores.ErrorVotanteFraudulento{Dni: votante.dni}
	}
	
	votante.fraudulento = true
	return votante.voto, nil
}
