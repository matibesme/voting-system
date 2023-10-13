package votos

import "fmt"

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
		nombre:     nombre,
		candidatos: candidatos,
	}
}

func CrearVotosEnBlanco() Partido {
	return &partidoEnBlanco{}
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {
	partido.votosCandidatos[tipo]++

}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	cadena := fmt.Sprintf("", partido.nombre, partido.candidatos[tipo])
	return partido.nombre
}

func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {
	blanco.votos[tipo]++
}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	return ""
}
