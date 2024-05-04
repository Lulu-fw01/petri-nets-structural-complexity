package algorithm

import (
	"complexity/pkg/net"
	"complexity/pkg/settings"
	"errors"
	"log"
)

func CountCharacteristicV3(sourceNet *net.PetriNet, settings settings.Settings) float64 {
	agentToTransition := getAgentToNotSilentTransitionsMap(sourceNet, settings)
	causalConnections := FindCausalConnections(sourceNet)
	transitionToConnections := SortCausalConnectionsByFromId(causalConnections)

	transitionToAgent := settings.GetTransitionToAgentMap(sourceNet.Transitions)

	result := 0.0

	agentToWeight := getAgentsWeights(causalConnections, transitionToAgent, settings)

	for agent, transitions := range agentToTransition {
		sum := 0.0
		for _, t := range transitions {
			connections := transitionToConnections[t]
			differentAgentsConnections := countDifferentAgentsConnections(agent, connections, transitionToAgent)
			sum += differentAgentsConnections / float64(len(connections))
		}
		w := agentToWeight[agent]
		result += w * sum / float64(len(transitions))
	}

	return 1 - result
}

func getAgentToNotSilentTransitionsMap(sourceNet *net.PetriNet, settings settings.Settings) map[string][]string {
	var transitions []*net.Transition
	for _, t := range sourceNet.Transitions {
		if !t.IsSilent {
			transitions = append(transitions, t)
		}
	}
	transitionToAgent := settings.GetTransitionToAgentMap(transitions)
	result := make(map[string][]string)

	for t, a := range transitionToAgent {
		result[a] = append(result[a], t)
	}

	return result
}

func getAgentsWeights(causalConnections []*CausalConnection, transitionToAgent map[string]string, settings settings.Settings) map[string]float64 {
	if settings.GetWeightType() == nil {
		return getStandardWeights(settings.GetAgents())
	}

	weightType := *settings.GetWeightType()
	switch weightType {
	case "custom":
		weights, err := getCustomWeights(settings)
		if err != nil {
			log.Printf("getAgentsWeights: Error getting custom agent weights. using standard weights. err %s", err)
			weights = getStandardWeights(settings.GetAgents())
		}
		return weights
	case "internal":
		return getInternalWeights(causalConnections, transitionToAgent)
	case "external":
		return getExternalWeights(causalConnections, transitionToAgent)
	default:
		return getStandardWeights(settings.GetAgents())
	}
}

func getCustomWeights(settings settings.Settings) (map[string]float64, error) {
	agentsCount := len(settings.GetAgents())
	if settings.GetAgentsWeights() == nil {
		return nil, errors.New("there are no agents weights in settings")
	}
	if len(*settings.GetAgentsWeights()) != agentsCount {
		return nil, errors.New("there are not all agents in agents weights")

	}
	return *settings.GetAgentsWeights(), nil
}

func getStandardWeights(agents []string) map[string]float64 {
	agentsCount := len(agents)
	result := make(map[string]float64)

	w := 1.0 / float64(agentsCount)
	for _, a := range agents {
		result[a] = w
	}
	return result
}

func getInternalWeights(causalConnections []*CausalConnection, transitionToAgent map[string]string) map[string]float64 {
	var internalConnections []*CausalConnection
	// Looking for connections with same agents (internal behavior).
	for _, c := range causalConnections {
		if transitionToAgent[c.ToTransitionId] == transitionToAgent[c.FromTransitionId] {
			internalConnections = append(internalConnections, c)
		}
	}

	// Count connections for each agent.
	agentToConnectionsCount := make(map[string]int)
	for _, c := range internalConnections {
		agent := transitionToAgent[c.FromTransitionId]
		agentToConnectionsCount[agent]++
	}
	result := make(map[string]float64)
	internalConnectionsCount := len(internalConnections)
	// Count weights.
	for agent, count := range agentToConnectionsCount {
		result[agent] = float64(count) / float64(internalConnectionsCount)
	}

	return result
}

func getExternalWeights(causalConnections []*CausalConnection, transitionToAgent map[string]string) map[string]float64 {
	var externalConnections []*CausalConnection
	// Looking for connections with different agents (external behavior)
	for _, c := range causalConnections {
		if transitionToAgent[c.ToTransitionId] != transitionToAgent[c.FromTransitionId] {
			externalConnections = append(externalConnections, c)
		}
	}

	externalConnectionsCount := len(externalConnections)

	// Count connections for each agent.
	agentToConnectionsCount := make(map[string]int)
	for _, c := range externalConnections {
		agent := transitionToAgent[c.FromTransitionId]
		agentToConnectionsCount[agent]++
	}
	// Count weights.
	result := make(map[string]float64)
	for agent, count := range agentToConnectionsCount {
		result[agent] = float64(count) / float64(externalConnectionsCount)
	}

	return result
}
