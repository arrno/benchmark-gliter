package main

import (
	"log"
	"testing"

	"github.com/arrno/gliter"
)

func setupTest(tb testing.TB) (*Simulator, func(tb testing.TB)) {
	log.Println("setup test")

	// make and start fresh
	sim := NewSimulator(makeDB())
	sim.Clear()
	// seed
	sim.LoadFakeData()

	return sim, func(tb testing.TB) {
		log.Println("teardown test")
		sim.Clear()
	}
}

// baseline
func TestSequential(t *testing.T) {

	sim, teardownTest := setupTest(t)
	defer teardownTest(t)

	// sim ten pages of work
	for _ = range 10 {
		data := sim.FetchData()
		newData := sim.DeriveNew(data)
		sim.SetData(newData)
	}

}

// with pipeline
func TestPipeline(t *testing.T) {

	sim, teardownTest := setupTest(t)
	defer teardownTest(t)

	gen := func() func() ([]DocBundle, bool, error) {
		idx := 0
		return func() ([]DocBundle, bool, error) {
			if idx >= 10 {
				return nil, false, nil
			}
			idx++
			return sim.FetchData(), true, nil
		}
	}

	work := func(data []DocBundle) ([]DocBundle, error) {
		return sim.DeriveNew(data), nil
	}

	store := func(data []DocBundle) ([]DocBundle, error) {
		sim.SetData(data)
		return nil, nil
	}

	gliter.NewPipeline(gen()).
		Stage(work).
		Stage(store).
		Run()
}
