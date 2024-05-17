package prom

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

type PageWithNet struct {
	XMLName     xml.Name      `xml:"page"`
	Places      []*Place      `xml:"place"`
	Transitions []*Transition `xml:"transition"`
	Arcs        []*Arc        `xml:"arc"`
}

type Net struct {
	XMLName xml.Name     `xml:"net"`
	Page    *PageWithNet `xml:"page"`
}

type Pnml struct {
	XMLName xml.Name `xml:"pnml"`
	Net     *Net     `xml:"net"`
}

func (pnml Pnml) Convert(netSettings settings.Settings) *net.PetriNet {
	wopedNet := pnml.Net.Page
	transitionIdToText := make(map[string]string)
	for _, t := range wopedNet.Transitions {
		transitionIdToText[t.Id] = t.Name.Text
	}
	return &net.PetriNet{
		Places:      convertPipePlacesToPlaces(wopedNet.Places),
		Transitions: convertPipeTransitionsToTransitions(wopedNet.Transitions, netSettings),
		Arcs:        convertPipeArcsToArcs(wopedNet.Arcs, transitionIdToText),
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

func convertPipeArcsToArcs(wopedArcs []*Arc, transitionIdToText map[string]string) []*net.Arc {
	var arcs []*net.Arc
	for _, pipeArc := range wopedArcs {
		source, target := pipeArc.Source, pipeArc.Target
		if actualSource, exist := transitionIdToText[source]; exist {
			source = actualSource
		}
		if actualTarget, exist := transitionIdToText[source]; exist {
			target = actualTarget
		}
		arcs = append(arcs, &net.Arc{Source: source, Target: target})
	}
	return arcs
}
