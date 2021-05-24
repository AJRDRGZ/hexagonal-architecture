package main

import (
	"github.com/AJRDRGZ/hexagonal-architecture/domain/poetry"
	"github.com/AJRDRGZ/hexagonal-architecture/infrastructure/echo"
	"github.com/AJRDRGZ/hexagonal-architecture/infrastructure/jsonservice"
)

func main() {
	// 1. Instantiate right-side adapters
	poetryLibrary := jsonservice.New("./poems.json")

	// 2. Instantiate the hexagon
	poetryReader := poetry.NewReader(poetryLibrary)

	// 3. Instantiate the left-side adapter(s)
	apiAdapter := echo.New(poetryReader)

	apiAdapter.Start(":1234")

	// NOTE:
	// El adaptador de Izquierda (consoleAdapter) MANEJA la lógica de negocio (poetryReader),
	// y el adaptador de derecha (poetryLibrary) es MANEJADO por la lógica de negocio.
}
