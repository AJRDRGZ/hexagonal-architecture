package console

import (
	"fmt"

	"github.com/AJRDRGZ/hexagonal-architecture/domain/poetry"
)

// Console allow print the poems
type Console struct {
	poetryReader poetry.IRequestVerses
}

// New returns a new Console
func New(poetry poetry.IRequestVerses) Console {
	return Console{poetry}
}

// Ask print a random poem
func (c Console) Ask() {
	fmt.Println("Here is some poetry:")
	fmt.Println(c.poetryReader.GiveMeSomePoetry())
}
