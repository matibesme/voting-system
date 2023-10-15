package main

Implementacion) VotadoPara(tipo TipoVoto) {
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
	if blanco.votosBlancos[tipo] == 1 {
		return fmt.Sprintf("Votos en Blanco: %d voto", blanco.votos[tipo])
	}
	return fmt.Sprintf("Votos en Blanco: %d votos", blanco.votosBlancos[tipo])
}


func main() {

	if len(os.Args) != 3 {
		fmt.Println(errores.ErrorParametros{})
		return
	}

	lista_partidos := partidosEnArchivo(os.Args[1])
	lista_padrones := partidosEnArchivo(os.Args[2])

	if lista_partidos == nil || lista_padrones == nil {
		fmt.Print(errores.ErrorLeerArchivo{})
		return
	}

	partidos := CrearPartidos(lista_partidos)
	cola_votantes := cola.CrearColaEnlazada[int]()

	texto_ingresado := bufio.NewScanner(os.Stdin)
	for texto_ingresado.Scan() {
		entrada := strings.Split(texto_ingresado.Text(), " ")

		switch entrada[0] {
		case "ingresar":
			acciones.IngresarVotante(entrada[1])

		case "votar":
			acciones.Votar()

		case "deshacer":
			acciones.Deshacer()

		case "fin-votar":
			acciones.FinVotar()

		}

	}

	acciones.Fin()

}

//Aca una par de funciones extras
//dsp reutilizar las dos lecturas de archivo

func partidosEnArchivo(archivo_lista string) []string {
	var partidos []string
	archivo, err := os.Open(archivo_lista)
	if err != nil {
		return nil
	}

	defer archivo.Close()
	lector := bufio.NewScanner(archivo)
	for lector.Scan() {

		partidos = append(partidos, lector.Text())
	}
	return partidos
}

func padronesEnArchivo(archivo_lista string) []votos.Votante {
	// leer archivo de padrones
	var padron []votos.Votante
	archivo, err := os.Open(archivo_lista)
	if err != nil {
		return nil
	}

	defer archivo.Close()
	lector := bufio.NewScanner(archivo)
	for lector.Scan() {
		//dni deberia ser string?
		padron = append(padron, votos.CrearVotante(lector.Text()))
	}
	return padron

}

func partidoValido(partidos []string, cantPartido int) bool {
	return len(partidos) >= cantPartido
}

// algun algoritmo de ordenamiento para fijarse que este el dni
// llamado desde acciones
func dniEnPadron() {

}

func CrearPartidos(partidos []string) []votos.Partido {}
