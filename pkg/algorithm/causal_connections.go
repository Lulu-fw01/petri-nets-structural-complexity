package algorithm

import (
	"complexity/pkg/algorithm/graph"
	"complexity/pkg/net"
)

func FindCausalConnections(net *net.PetriNet) []*CausalConnection {
	placesById := getPlacesMap(net.Places)
	transitionById := getTransitionsMap(net.Transitions)
	startPlaces := getStartPlaces(placesById, net.Arcs)
	idToElement := getElements(net.Places, net.Transitions)
	graphOfNet := graph.GetGraph(net.Arcs)
	description := graphDescription{
		graph:          graphOfNet,
		startPlaces:    startPlaces,
		idToElement:    idToElement,
		transitionById: transitionById,
	}
	connections := findCausalConnections(&description)
	return connections
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

func findCausalConnections(
	description *graphDescription,
) []*CausalConnection {
	// Need for unnecessary causal connections search.
	transitionGoNextBlock := make(map[string]bool)
	// Need for correct working with cycles.
	elementToCheck := make(map[string]bool)
	for key := range description.idToElement {
		elementToCheck[key] = false
	}
	for key, value := range description.transitionById {
		if !value.IsSilent {
			transitionGoNextBlock[key] = false
		}
	}
	var connections []*CausalConnection
	for _, p := range description.startPlaces {
		connections = append(connections, findCausalConnectionsRec(description, &elementToCheck, &transitionGoNextBlock, nil, p)...)
	}
	return connections
}

func findCausalConnectionsRec(
	description *graphDescription,
	elementToCheck *map[string]bool,
	transitionGoNextBlock *map[string]bool,
	fromTransitionId *string,
	elementId string) []*CausalConnection {

	var connections []*CausalConnection

	elem := (*description).idToElement[elementId]
	var fromElement *string
	if elem.isTransition {
		transition := description.transitionById[elem.id]
		if transition.IsSilent {
			// Black transition can be inside causal connection.
			// Do not change from transition and go next.
			fromElement = fromTransitionId
		} else {
			// Current element is transition.
			// Add new causal connection (if it is not first
			// met transition in algorithm).
			if fromTransitionId != nil {
				connections = append(connections, &CausalConnection{
					FromTransitionId: *fromTransitionId,
					ToTransitionId:   elementId,
				})
			}
			// If we checked all causal connections after this element (started at thi element).
			if (*transitionGoNextBlock)[elementId] {
				return connections
			}
			(*transitionGoNextBlock)[elementId] = true
			// Change from transition to current element.
			fromElement = &elementId
		}
	} else {
		// Current element is place.
		// Do not change from transition and go next.
		fromElement = fromTransitionId
	}

	// Check that we have never been to next elements.
	if (*elementToCheck)[elementId] {
		return connections
	} else {
		(*elementToCheck)[elementId] = true
	}
	nextElements, exists := (*description).graph[elementId]
	if exists {
		for _, e := range nextElements {
			nextConnections := findCausalConnectionsRec(description, elementToCheck, transitionGoNextBlock, fromElement, e)
			connections = append(connections, nextConnections...)
		}

	}
	(*elementToCheck)[elementId] = false
	return connections
}

func getElements(places []*net.Place, transitions []*net.Transition) map[string]*element {
	elements := make(map[string]*element)
	for _, p := range places {
		elements[p.Id] = &element{
			id:           p.Id,
			isTransition: false,
		}
	}
	for _, t := range transitions {
		elements[t.Id] = &element{
			id:           t.Id,
			isTransition: true,
		}
	}
	return elements
}

type graphDescription struct {
	graph          map[string][]string
	startPlaces    []string
	idToElement    map[string]*element
	transitionById map[string]*net.Transition
}

type element struct {
	id           string
	isTransition bool
}

type CausalConnection struct {
	FromTransitionId string
	ToTransitionId   string
}

func SortCausalConnectionsByFromId(causalConnections []*CausalConnection) map[string][]*CausalConnection {
	transitionToConnections := make(map[string][]*CausalConnection)
	// Sort causal connections by from id.
	for _, causalConnection := range causalConnections {
		transitionToConnections[causalConnection.FromTransitionId] = append(transitionToConnections[causalConnection.FromTransitionId], causalConnection)
	}
	return transitionToConnections
}
