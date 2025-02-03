package main

import (
	"testing"

	"github.com/arrno/gliter"
)

const PAGES int = 10

// baseline
func TestSequential(t *testing.T) {

	sim := NewSimulator(makeMockDB())
	defer sim.Clear()
	sim.LoadFakeData()

	// sim ten pages of work
	for _ = range PAGES {
		data := sim.FetchData()
		newData := sim.DeriveNew(data)
		sim.SetData(newData)
	}

}

// with pipeline
func TestPipeline(t *testing.T) {

	sim := NewSimulator(makeMockDB())
	defer sim.Clear()
	sim.LoadFakeData()

	gen := func() func() ([]DocBundle, bool, error) {
		idx := 0
		return func() ([]DocBundle, bool, error) {
			if idx >= PAGES {
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
