package poetry

import (
	"fmt"
	"strings"
)

//---------------------------------------------------------------------------------------
// Ports
//---------------------------------------------------------------------------------------

// IObtainPoem obtain poems from a library or storage (right-side)
type IObtainPoem interface {
	GetAPoem() string
}

//---------------------------------------------------------------------------------------
// Reader
//---------------------------------------------------------------------------------------

// Reader is the object that read the poetries
type Reader struct {
	library IObtainPoem
}

// NewReader returs a new Reader of poems
func NewReader(library IObtainPoem) Reader {
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
