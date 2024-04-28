package woped

import (
	"complexity/pkg/net"
	"complexity/pkg/settings"
	"encoding/xml"
)

type Place struct {
	XmlName xml.Name `xml:"place"`
	Id      string   `xml:"id,attr"`
}

type Transition struct {
	XMLName xml.Name `xml:"transition"`
	Id      string   `xml:"id,attr"`
	Name    Name     `xml:"name"`
}

type Name struct {
	Text string `xml:"text"`
}

type Arc struct {
	XMLName xml.Name `xml:"arc"`
	Source  string   `xml:"source,attr"`
	Target  string   `xml:"target,attr"`
}

type PetriNet struct {
	XMLName     xml.Name      `xml:"net"`
	Places      []*Place      `xml:"place"`
	Transitions []*Transition `xml:"transition"`
	Arcs        []*Arc        `xml:"arc"`
}

type Pnml struct {
	XMLName xml.Name  `xml:"pnml"`
	Net     *PetriNet `xml:"net"`
}

func (pnml Pnml) Convert(netSettings settings.Settings) *net.PetriNet {
	pipeNet := pnml.Net
	// todo rewrite!!!
	return &net.PetriNet{
		Places:      convertPipePlacesToPlaces(pipeNet.Places),
		Transitions: convertPipeTransitionsToTransitions(pipeNet.Transitions, netSettings),
		Arcs:        convertPipeArcsToArcs(pipeNet.Arcs),
	}
}

func convertPipePlacesToPlaces(pipePlaces []*Place) []*net.Place {
	var places []*net.Place
	for _, pipePlace := range pipePlaces {
		places = append(places, &net.Place{Id: pipePlace.Id})
	}
	return places
}

func convertPipeTransitionsToTransitions(pipeTransitions []*Transition, netSettings settings.Settings) []*net.Transition {
	var transitions []*net.Transition
	for _, pipeTransition := range pipeTransitions {
		transitions = append(transitions, &net.Transition{Id: pipeTransition.Id, IsSilent: netSettings.IsSilentTransition(pipeTransition.Id)})
	}
	return transitions
}

func convertPipeArcsToArcs(pipeArcs []*Arc) []*net.Arc {
	var arcs []*net.Arc
	for _, pipeArc := range pipeArcs {
		arcs = append(arcs, &net.Arc{Source: pipeArc.Source, Target: pipeArc.Target})
	}
	return arcs
}
