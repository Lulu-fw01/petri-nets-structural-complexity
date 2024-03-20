package algorithm

import (
	"complexity/pkg/net"
	"complexity/pkg/settings"
)

func CountRatios(net *net.PetriNet, settings *settings.Settings) []RatioResult {
	transitionToAgent := getTransitionToAgentMap(settings)
	connections := FindCausalConnections(net)

	pairToConnections := make(map[string][]*CausalConnection)
	connectionsCount := len(connections)
	for _, conn := range connections {
		key1 := transitionToAgent[conn.FromTransitionId] + "-" + transitionToAgent[conn.ToTransitionId]
		key2 := transitionToAgent[conn.ToTransitionId] + "-" + transitionToAgent[conn.FromTransitionId]
		pairConnections, exists := pairToConnections[key1]
		if exists {
			pairToConnections[key1] = append(pairConnections, conn)
		} else {
			pairConnections, exists := pairToConnections[key2]
			if exists {
				pairToConnections[key2] = append(pairConnections, conn)
			} else {
				pairToConnections[key1] = []*CausalConnection{conn}
			}
		}
	}

	var result []RatioResult
	for _, agentsConnections := range pairToConnections {
		first := agentsConnections[0]
		fromAgent := transitionToAgent[first.FromTransitionId]
		toAgent := transitionToAgent[first.ToTransitionId]
		if fromAgent != toAgent {
			arcsCount := len(agentsConnections)
			result = append(result, RatioResult{
				agentOne: fromAgent,
				agentTwo: toAgent,
				ratio:    float64(arcsCount) / float64(connectionsCount),
			})
		}
	}
	return result
}

func FindCausalConnections(net *net.PetriNet) []*CausalConnection {
	placesById := getPlacesMap(net.Places)
	transitionById := getTransitionsMap(net.Transitions)
	startPlaces := getStartPlaces(placesById, net.Arcs)
	idToElement := getElements(net.Places, net.Transitions)
	graph := getGraph(net.Arcs)
	description := graphDescription{
		graph:          graph,
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

// Return map where key is element id and value
// is list of element ids which connected to key with input arcs.
func getGraph(
	arcs []*net.Arc) map[string][]string {
	elements := make(map[string][]string)

	for _, arc := range arcs {
		source := arc.Source
		target := arc.Target
		if _, exists := elements[source]; exists {
			elements[source] = append(elements[source], target)
		} else {
			elements[source] = []string{target}
		}
	}
	return elements
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

func getArcsByElements(
	placesById map[string]*net.Place,
	transitionsById map[string]*net.Transition,
	arcs []*net.Arc) (map[string][]*net.Arc, map[string][]*net.Arc) {

	arcsFromPlace := make(map[string][]*net.Arc)
	arcsFromTransition := make(map[string][]*net.Arc)
	for _, arc := range arcs {
		if _, tExists := transitionsById[arc.Source]; tExists {
			arcsList := arcsFromTransition[arc.Source]
			arcsFromTransition[arc.Source] = append(arcsList, arc)
		} else if _, pExists := placesById[arc.Source]; pExists {
			arcsList := arcsFromPlace[arc.Source]
			arcsFromPlace[arc.Source] = append(arcsList, arc)
		}
	}
	return arcsFromPlace, arcsFromTransition
}

type graphDescription struct {
	graph          map[string][]string
	startPlaces    []string
	idToElement    map[string]*element
	transitionById map[string]*net.Transition
}

type RatioResult struct {
	agentOne string
	agentTwo string
	ratio    float64
}

type element struct {
	id           string
	isTransition bool
}

type CausalConnection struct {
	FromTransitionId string
	ToTransitionId   string
}

// result:
// Ai
// Ab
// metric result
