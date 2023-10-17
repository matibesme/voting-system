package votos

import (
	"fmt"
)

type partidoImplementacion struct {
	nombre          string
	candidatos      [CANT_VOTACION]string
	votosCandidatos [CANT_VOTACION]int
}

type partidoEnBlanco struct {
	votos [CANT_VOTACION]int
}

func CrearPartido(nombre string, candidatos [CANT_VOTACION]string) Partido {
	return &partidoImplementacion{
		nombre:          nombre,
		candidatos:      candidatos,
		votosCandidatos: [CANT_VOTACION]int{0, 0, 0},
	}
}

func CrearVotosEnBlanco() Partido {
	return &partidoEnBlanco{}
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {
	partido.votosCandidatos[tipo]++

}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	formato := "%s - %s: %d voto"
	if partido.votosCandidatos[tipo] != 1 {
		formato = "%s - %s: %d votos"
	}
	return fmt.Sprintf(formato, partido.nombre, partido.candidatos[tipo], partido.votosCandidatos[tipo])
}

func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {
	blanco.votos[tipo]++
}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	if blanco.votos[tipo] == 1 {
		return fmt.Sprintf("Votos en Blanco: %d voto", blanco.votos[tipo])
	}
	return fmt.Sprintf("Votos en Blanco: %d votos", blanco.votos[tipo])
}
