package poetry

import (
	"fmt"
	"strings"

	"github.com/AJRDRGZ/hexagonal-architecture/infrastructure/postgres"
)

// Reader is the object that read the poetries
type Reader struct {
	library postgres.Storage
}

// NewReader returs a new Reader of poems
func NewReader(library postgres.Storage) Reader {
	return Reader{library}
}

// GiveMeSomePoetry returns a poem from the library in uppercase
func (r Reader) GiveMeSomePoetry() string {
	poem := r.library.GetAPoem()
	return strings.ToUpper(poem)
}

// Ask print a random poem
func (r Reader) Ask() {
	fmt.Println("Here is some poetry:")
	fmt.Println(r.GiveMeSomePoetry())
}
