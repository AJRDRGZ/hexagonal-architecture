package main

import (
	"github.com/AJRDRGZ/hexagonal-architecture/domain/poetry"
	"github.com/AJRDRGZ/hexagonal-architecture/infrastructure/console"
	"github.com/AJRDRGZ/hexagonal-architecture/infrastructure/jsonservice"
)

func main() {
	// 1. Instantiate right-side adapters
	poetryLibrary := jsonservice.New("./poems.json")

	// 2. Instantiate the hexagon
	poetryReader := poetry.NewReader(poetryLibrary)

	// 3. Instantiate the left-side adapter(s)
	consoleAdapter := console.New(poetryReader)

	consoleAdapter.Ask()

	// NOTE:
	// El adaptador de Izquierda (consoleAdapter) MANEJA la lógica de negocio (poetryReader),
	// y el adaptador de derecha (poetryLibrary) es MANEJADO por la lógica de negocio.
}
