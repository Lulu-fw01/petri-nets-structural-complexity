package net

type Place struct {
	Id string
}

type Transition struct {
	Id      string
	IsBlack bool
}

type Arc struct {
	Source string
	Target string
}

type PetriNet struct {
	Places      []*Place
	Transitions []*Transition
	Arcs        []*Arc
}
