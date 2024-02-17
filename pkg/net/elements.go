package net

type Place struct {
	id string
}

type Transition struct {
	id string
}

type Arc struct {
	source string
	target string
}

type PetriNet struct {
	places      []Place
	transitions []Transition
	arcs        []Arc
}
