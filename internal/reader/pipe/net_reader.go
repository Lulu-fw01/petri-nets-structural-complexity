package pipe

import (
	"complexity/pkg/net"
	"encoding/xml"
	"golang.org/x/net/html/charset"
	"log"
	"os"
)

func ReadNet(path string, silentTransitions []string) (*net.PetriNet, error) {
	pipeNet, err := ReadPipeNet(path)
	if err != nil {
		return nil, err
	}
	return convertPipeNetToNet(pipeNet, silentTransitions), nil
}

func ReadPipeNet(path string) (*PetriNet, error) {
	netFile, err := os.Open(path)
	if err != nil {
		log.Fatalf("Could not open file with net. Error: %s", err)
		return nil, err
	}
	defer netFile.Close()

	xmlDecoder := xml.NewDecoder(netFile)
	xmlDecoder.CharsetReader = charset.NewReaderLabel

	var newNet Pnml
	err = xmlDecoder.Decode(&newNet)
	if err != nil {
		log.Fatalf("Error decoding xml file: %s", err)
		return nil, err
	}

	return newNet.Net, nil
}

func convertPipeNetToNet(pipeNet *PetriNet, silentTransitions []string) *net.PetriNet {
	return &net.PetriNet{
		Places:      convertPipePlacesToPlaces(pipeNet.Places),
		Transitions: convertPipeTransitionsToTransitions(pipeNet.Transitions, silentTransitions),
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

func convertPipeTransitionsToTransitions(pipeTransitions []*Transition, silentTransitions []string) []*net.Transition {
	var transitions []*net.Transition
	for _, pipeTransition := range pipeTransitions {
		transitions = append(transitions, &net.Transition{Id: pipeTransition.Id, IsSilent: isSilentTransition(pipeTransition.Id, silentTransitions)})
	}
	return transitions
}

func isSilentTransition(transitionId string, silentTransitions []string) bool {
	for _, t := range silentTransitions {
		if t == transitionId {
			return true
		}
	}
	return false
}

func convertPipeArcsToArcs(pipeArcs []*Arc) []*net.Arc {
	var arcs []*net.Arc
	for _, pipeArc := range pipeArcs {
		arcs = append(arcs, &net.Arc{Source: pipeArc.Source, Target: pipeArc.Target})
	}
	return arcs
}
