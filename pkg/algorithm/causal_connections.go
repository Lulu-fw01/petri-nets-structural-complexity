package algorithm

import (
	"complexity/pkg/net"
	"complexity/pkg/settings"
)

func FindCausalConnections(net *net.PetriNet, settings *settings.Settings) {
	transitionToAgent := getTransitionToAgentMap(settings)

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

// TODO rename.
func CountCausalConnectionsMetric(net *net.PetriNet) {

}

// result:
// Ai
// Ab
// metric result
