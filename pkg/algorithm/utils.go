package algorithm

import "complexity/pkg/net"

// return map where key is place id and Place as value.
func getPlacesMap(places []*net.Place) map[string]*net.Place {
	placesById := make(map[string]*net.Place)
	for _, p := range places {
		placesById[p.Id] = p
	}
	return placesById
}

func getTransitionsMap(transitions []*net.Transition) map[string]*net.Transition {
	transitionsById := make(map[string]*net.Transition)
	for _, t := range transitions {
		transitionsById[t.Id] = t
	}
	return transitionsById
}

// getInputAndOutputArcs returns two maps. The first map contains elements as keys and slices of arcs that end at these elements as values.
// The second map contains elements as keys and slices of arcs that start at these elements as values.
func getInputAndOutputArcs(arcs []*net.Arc) (map[string][]*net.Arc, map[string][]*net.Arc) {
	elementToOutputArcs := make(map[string][]*net.Arc)
	elementToInputArcs := make(map[string][]*net.Arc)

	for _, arc := range arcs {
		from := arc.Source
		to := arc.Target

		elementToOutputArcs[from] = append(elementToOutputArcs[from], arc)
		elementToInputArcs[to] = append(elementToInputArcs[to], arc)
	}
	return elementToInputArcs, elementToOutputArcs
}
