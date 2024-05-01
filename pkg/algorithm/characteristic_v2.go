package algorithm

import (
	"complexity/pkg/net"
	"complexity/pkg/settings"
)

func CountCharacteristicV2(net *net.PetriNet, settings settings.Settings) float64 {
	// Getting all causal connection.
	causalConnections := FindCausalConnections(net)
	transitionToConnections := make(map[string][]*CausalConnection)
	// Sort causal connections by from id.
	for _, causalConnection := range causalConnections {
		transitionToConnections[causalConnection.FromTransitionId] = append(transitionToConnections[causalConnection.FromTransitionId], causalConnection)
	}

	transitionToAgent := settings.GetTransitionToAgentMap(net.Transitions)

	sum := 0.0
	for transition, connections := range transitionToConnections {
		tAgent := transitionToAgent[transition]
		differentAgentsConnections := 0.0
		for _, c := range connections {
			if transitionToAgent[c.ToTransitionId] != tAgent {
				differentAgentsConnections++
			}
		}

		sum += differentAgentsConnections / float64(len(connections))
	}

	// Looking for not Silent transitions.
	notSilentTransitionsCount := 0.0
	for _, t := range net.Transitions {
		if !t.IsSilent {
			notSilentTransitionsCount++
		}
	}

	return 1 - sum/notSilentTransitionsCount
}
