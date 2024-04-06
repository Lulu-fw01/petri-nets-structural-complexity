package algorithm

import (
	"complexity/pkg/net"
	"complexity/pkg/settings"
)

func FindChannels(pNet *net.PetriNet, settings *settings.Settings) map[string]*Channel {
	idToPlace := make(map[string]*Channel)
	transitionToAgent := getTransitionToAgentMap(settings)
	idToInputArcs, idToOutputArcs := getInputAndOutputArcs(pNet.Arcs)

	// Move through all places.
	for _, p := range pNet.Places {
		// Get input and output arcs for this place.
		inputArcs, hasInput := idToInputArcs[p.Id]
		outputArcs, hasOutput := idToOutputArcs[p.Id]
		if !hasInput || !hasOutput {
			continue
		}
		// Get agents of input arcs.
		// Get agents of output arcs.
		fromAgents := make(map[string]int)
		toAgents := make(map[string]int)
		for _, inA := range inputArcs {
			fromAgents[transitionToAgent[inA.Source]]++
		}
		for _, outA := range outputArcs {
			toAgents[transitionToAgent[outA.Target]]++
		}
		isChannel := true
		if len(fromAgents) == 1 && len(toAgents) == 1 {
			fromKeys := getKeys(fromAgents)
			toKeys := getKeys(toAgents)
			if fromKeys[0] == toKeys[0] {
				isChannel = false
			}
		}
		if isChannel {
			idToPlace[p.Id] = &Channel{
				PlaceId:    p.Id,
				InputArcs:  inputArcs,
				OutputArcs: outputArcs,
			}
		}
	}
	return idToPlace
}

func getKeys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

type Channel struct {
	PlaceId    string
	InputArcs  []*net.Arc
	OutputArcs []*net.Arc
}
