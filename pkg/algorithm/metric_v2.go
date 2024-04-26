package algorithm

import (
	"complexity/pkg/algorithm/model"
	"complexity/pkg/net"
	"complexity/pkg/settings"
)

func CountMetric(net *net.PetriNet, settings settings.Settings) float64 {
	causalConnections := FindCausalConnections(net)
	channels := FindChannels(net, settings)
	transitionToAgent := settings.GetTransitionToAgentMap(net.Transitions)
	agentToCausalConnections := getCausalConnectionsInsideEveryAgent(causalConnections, transitionToAgent)

	allChannelsArcs := 0.0
	for _, channel := range channels {
		allChannelsArcs += float64(len(channel.InputArcs) + len(channel.OutputArcs))
	}

	sum := 0.0
	for _, channel := range channels {
		currentChannelArcsCount := float64(len(channel.InputArcs) + len(channel.OutputArcs))
		inputToAllRatio := float64(len(channel.InputArcs)) / currentChannelArcsCount
		outputToAllRatio := float64(len(channel.OutputArcs)) / currentChannelArcsCount

		w := 0.0
		agentsPairToCount := make(map[model.AgentsPair]int)

		for _, inputArc := range channel.InputArcs {
			fromTransition := inputArc.Source
			for _, outputArc := range channel.OutputArcs {
				toTransition := outputArc.Target
				if transitionToAgent[fromTransition] != transitionToAgent[toTransition] {
					pair := model.AgentsPair{
						FromAgent: transitionToAgent[fromTransition],
						ToAgent:   transitionToAgent[toTransition],
					}
					agentsPairToCount[pair]++
				}
			}
		}

		for pair, count := range agentsPairToCount {
			fromAgent := pair.FromAgent
			toAgent := pair.ToAgent
			w += (float64(count) / float64(len(agentToCausalConnections[fromAgent])+count)) * inputToAllRatio
			w += (float64(count) / float64(len(agentToCausalConnections[toAgent])+count)) * outputToAllRatio
		}
		sum += w * (currentChannelArcsCount / allChannelsArcs)
	}
	return 1 - sum
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
