package algorithm

import (
	"complexity/pkg/algorithm/graph"
	"complexity/pkg/net"
	"complexity/pkg/settings"
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

func getTransitionToAgentMap(settings *settings.Settings) map[string]string {
	transitionToAgent := make(map[string]string)
	for agent, transitions := range settings.AgentsToTransitions {
		for _, t := range transitions {
			transitionToAgent[t] = agent
		}
	}
	return transitionToAgent
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
	elementToCheck := make(map[string]bool)
	for key := range description.idToElement {
		elementToCheck[key] = false
	}
	var connections []*CausalConnection
	for _, p := range description.startPlaces {
		connections = append(connections, findCausalConnectionsRec(description, &elementToCheck, nil, p)...)
	}
	return connections
}

func findCausalConnectionsRec(
	description *graphDescription,
	elementToCheck *map[string]bool,
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
			nextConnections := findCausalConnectionsRec(description, elementToCheck, fromElement, e)
			connections = append(connections, nextConnections...)
		}
	}
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
