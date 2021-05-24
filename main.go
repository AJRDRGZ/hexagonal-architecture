package main

import (
	"github.com/AJRDRGZ/hexagonal-architecture/poetry"
	"github.com/AJRDRGZ/hexagonal-architecture/storage"
)

func main() {
	library := storage.New()
	poetryReader := poetry.NewReader(library)
	poetryReader.Ask()
}
