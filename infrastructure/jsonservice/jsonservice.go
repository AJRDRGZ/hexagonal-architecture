package jsonservice

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"
)

type poem struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Poem   string `json:"poem"`
}

type poems struct {
	Data []poem `json:"poems"`
}

// Storage from a json file
type Storage struct {
	poems poems
}

var defaultPoem = poem{Author: "Ajrdrgz", Poem: "my default poem"}

// New returns a new Storage JSON
func New(file string) Storage {
	storage := Storage{}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		storage.poems.Data = []poem{defaultPoem}
		return storage
	}

	if err := json.Unmarshal(data, &storage.poems); err != nil {
		storage.poems.Data = []poem{defaultPoem}
	}

	return storage
}

// GetAPoem returns a random poem
func (s Storage) GetAPoem() string {
	return s.poems.Data[s.random()].Poem
}

func (s Storage) random() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := len(s.poems.Data)
	if max == 1 {
		max = 2
	}

	return (rand.Intn(max-min) + min) - 1
}
