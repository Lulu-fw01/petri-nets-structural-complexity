package algorithm

import (
	"complexity/pkg/net"
	"complexity/pkg/settings"
)

func CountMetric(net *net.PetriNet, settings *settings.Settings) float64 {
	causalConnections := FindCausalConnections(net)
	channels := FindChannels(net, settings)
	agents := getAgents(settings.AgentsToTransitions)
	transitionToAgent := getTransitionToAgentMap(settings)
	agentToCausalConnections := getCausalConnectionsInsideEveryAgent(causalConnections, transitionToAgent)

	sum := 0.0
	for _, channel := range channels {
		allChannelArcsCount := float64(len(channel.InputArcs) + len(channel.OutputArcs))
		inputToAllRatio := float64(len(channel.InputArcs)) / allChannelArcsCount
		outputToAllRatio := float64(len(channel.OutputArcs)) / allChannelArcsCount

		w := 0.0
		for _, a1 := range agents {
			for _, a2 := range agents {
				if a1 == a2 {
					continue
				}
				w += countForAgents(channel, a1, a2, transitionToAgent, agentToCausalConnections)
			}
		}
		sum += w

	}
	return 1 - sum
}

func countForAgents(
	channel *Channel,
	fromAgent string,
	toAgent string,
	transitionToAgent map[string]string,
	agentToCausalConnections map[string][]*CausalConnection) float64 {
	causalConnectionsCountForChannel := 0.0
	for _, inputArc := range channel.InputArcs {
		fromTransition := inputArc.Source
		for _, outputArc := range channel.OutputArcs {
			toTransition := outputArc.Target
			if transitionToAgent[fromTransition] == fromAgent && transitionToAgent[toTransition] == toAgent {
				causalConnectionsCountForChannel++
			}
		}
	}
	fromAgentCausalConnectionsCount := float64(len(agentToCausalConnections[fromAgent]))
	toAgentCausalConnectionsCount := float64(len(agentToCausalConnections[toAgent]))
	allArcsCount := float64(len(channel.InputArcs) + len(channel.OutputArcs))

	fromAgentResult := causalConnectionsCountForChannel * float64(len(channel.InputArcs)) /
		((fromAgentCausalConnectionsCount + causalConnectionsCountForChannel) * allArcsCount)
	toAgentResult := causalConnectionsCountForChannel * float64(len(channel.InputArcs)) /
		((toAgentCausalConnectionsCount + causalConnectionsCountForChannel) * allArcsCount)
	return fromAgentResult + toAgentResult
}

func getCausalConnectionsInsideEveryAgent(
	causalConnections []*CausalConnection,
	transitionToAgent map[string]string) map[string][]*CausalConnection {
	agentToConnections := make(map[string][]*CausalConnection)
	for _, connection := range causalConnections {
		fromAgent := transitionToAgent[connection.FromTransitionId]
		toAgent := transitionToAgent[connection.ToTransitionId]
		if fromAgent == toAgent {
			agentToConnections[fromAgent] = append(agentToConnections[fromAgent], connection)
		}
	}
	return agentToConnections
}

func getAgents(agentToTransitions map[string][]string) []string {
	var agents []string
	for agent := range agentToTransitions {
		agents = append(agents, agent)
	}
	return agents
}
