package poetry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type FakeLibrary struct {
	Poem string
}

func (f FakeLibrary) GetAPoem() string { return f.Poem }

func TestReader_GiveMeSomePoetry(t *testing.T) {
	// 1. Instantiate right-side adapters
	poetryStub := FakeLibrary{
		Poem: "This is my fake Poem by Ajrdrgz",
	}

	// 2. Instantiate the hexagon
	poetryReader := NewReader(poetryStub)

	// 3. Driving from test (left-side)
	got := poetryReader.GiveMeSomePoetry()

	assert.Equal(t, "THIS IS MY FAKE POEM BY AJRDRGZ", got)
}
