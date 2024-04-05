package algorithm

import (
	"complexity/pkg/net"
	"complexity/pkg/settings"
)

func CountMetric(net *net.PetriNet, settings *settings.Settings) float64 {
	causalConnections := FindCausalConnections(net)
	channels := FindChannels(net, settings)
	transitionToAgent := getTransitionToAgentMap(settings)
	agentToCausalConnections := getCausalConnectionsInsideEveryAgent(causalConnections, transitionToAgent)

	sum := 0.0
	for _, channel := range channels {
		allChannelArcsCount := float64(len(channel.InputArcs) + len(channel.OutputArcs))
		inputToAllRatio := float64(len(channel.InputArcs)) / allChannelArcsCount
		outputToAllRatio := float64(len(channel.OutputArcs)) / allChannelArcsCount

		w := 0.0
		agentsPairToCount := make(map[AgentsPair]int)

		for _, inputArc := range channel.InputArcs {
			fromTransition := inputArc.Source
			for _, outputArc := range channel.OutputArcs {
				toTransition := outputArc.Target
				if transitionToAgent[fromTransition] != transitionToAgent[toTransition] {
					pair := AgentsPair{
						fromAgent: transitionToAgent[fromTransition],
						toAgent:   transitionToAgent[toTransition],
					}
					agentsPairToCount[pair]++
				}
			}
		}

		for pair, count := range agentsPairToCount {
			fromAgent := pair.fromAgent
			toAgent := pair.toAgent
			w += (float64(count) / float64(len(agentToCausalConnections[fromAgent])+count)) * inputToAllRatio
			w += (float64(count) / float64(len(agentToCausalConnections[toAgent])+count)) * outputToAllRatio
		}
		// todo sum += w * (Qc / Q)
		sum += w
	}
	return 1 - sum
}

type AgentsPair struct {
	fromAgent string
	toAgent   string
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
