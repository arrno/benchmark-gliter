package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

const COL_SIZE int = 100
const VAL_SIZE int = 10
const COL_NAME string = "mockdata"

type MyData struct {
	ID     string
	Name   string
	Values []float64
}

type DocBundle struct {
	Data MyData
	Path string
}
type Database interface {
	SetBatch([]DocBundle)
	FetchBatch(path string) []DocBundle
	DeleteCollection(path string)
}

type Simulator struct {
	db  Database
	uid int
}

func NewSimulator(db Database) *Simulator {
	return &Simulator{db, 0}
}

func (s *Simulator) MakeMyData() MyData {
	names := names()
	values := make([]float64, VAL_SIZE)
	for i := range VAL_SIZE {
		values[i] = s.jumbleFloat(rand.Float64())
	}
	d := MyData{
		ID:     strconv.Itoa(s.uid),
		Name:   names[s.uid%len(names)],
		Values: values,
	}
	s.uid++
	return d
}

func (s *Simulator) Clear() {
	s.db.DeleteCollection(COL_NAME)
}

func (s *Simulator) LoadFakeData() {
	docSet := make([]DocBundle, COL_SIZE)
	for i := range COL_SIZE {
		id, err := gonanoid.New()
		expectNil(err)
		docSet[i] = DocBundle{Path: fmt.Sprintf("%s/%s", COL_NAME, id), Data: s.MakeMyData()}
	}
	s.db.SetBatch(docSet)
}

func (s *Simulator) FetchData() []DocBundle {
	return s.db.FetchBatch(COL_NAME)
}

func (s *Simulator) DeriveNew(data []DocBundle) []DocBundle {
	newSet := make([]DocBundle, len(data))
	for i, db := range data {
		newBundle := DocBundle{
			Path: db.Path,
			Data: MyData{
				ID:     db.Data.ID,
				Name:   db.Data.Name,
				Values: make([]float64, len(db.Data.Values)),
			},
		}
		for j, val := range db.Data.Values {
			newBundle.Data.Values[j] = s.jumbleFloat(val)
		}
		newSet[i] = newBundle
	}
	return newSet
}

func (s *Simulator) SetData(data []DocBundle) {
	s.db.SetBatch(data)
}

func (s *Simulator) jumbleFloat(x float64) float64 {
	x = math.Mod(x*1.61803398875, 1.0) // Multiply by golden ratio and mod 1
	x = math.Sin(x * 2 * math.Pi)      // Keep it bounded
	return x
}
