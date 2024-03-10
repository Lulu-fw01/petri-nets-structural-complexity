package algorithm

import (
	"complexity/pkg/net"
	"complexity/pkg/settings"
)

func FindCausalConnections(net *net.PetriNet, settings *settings.Settings) {
	transitionToAgent := getTransitionToAgentMap(settings)
	placesById := getPlacesMap(net.Places)
	transitionById := getTransitionsMap(net.Transitions)
	startPlaces := getStartPlaces(placesById, net.Arcs)
	transitionsInputArcsCount := getTransitionsInputArcsCount(net.Arcs, transitionById)

}

func getTransitionToAgentMap(settings *settings.Settings) map[string]string {
	transitionToAgent := make(map[string]string)
	for agent, transitions := range settings.AgentsToTransitions {
		for _, t := range transitions {
			transitionToAgent[t] = agent
		}
	}
	return transitionToAgent
}

func getTransitionsInputArcsCount(arcs []*net.Arc, transitionsById map[string]*net.Transition) map[string]int {
	transitionToInputArcsCount := make(map[string]int)
	for _, arc := range arcs {
		in := arc.Target
		// If in element is transition.
		if _, exists := transitionsById[in]; exists {
			transitionToInputArcsCount[in]++
		}
	}
	return transitionToInputArcsCount
}

func getTransitionsMap(transitions []*net.Transition) map[string]*net.Transition {
	transitionsById := make(map[string]*net.Transition)
	for _, t := range transitions {
		transitionsById[t.Id] = t
	}
	return transitionsById
}

func getPlacesMap(places []*net.Place) map[string]*net.Place {
	placesById := make(map[string]*net.Place)
	for _, p := range places {
		placesById[p.Id] = p
	}
	return placesById
}

func getStartPlaces(placesById map[string]*net.Place, arcs []*net.Arc) []string {
	for _, arc := range arcs {
		delete(placesById, arc.Target)
	}
	places := make([]string, 0, len(placesById))
	for k := range placesById {
		places = append(places, k)
	}
	return places
}

// result:
// Ai
// Ab
// metric result
