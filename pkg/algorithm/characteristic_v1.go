package algorithm

import (
	"complexity/pkg/algorithm/model"
	"complexity/pkg/net"
	"complexity/pkg/settings"
)

func CountCharacteristicV1(net *net.PetriNet, settings settings.Settings) float64 {
	ratios := countRatios(net, settings)
	result := 0.0
	for _, r := range ratios {
		result += r.ratio
	}
	return 1 - result
}

func countRatios(net *net.PetriNet, settings settings.Settings) []RatioResult {
	// todo refactor.
	transitionToAgent := settings.GetTransitionToAgentMap(net.Transitions)
	connections := FindCausalConnections(net)

	pairToConnections := make(map[model.AgentsPair][]*CausalConnection)
	connectionsCount := len(connections)
	for _, conn := range connections {
		key1 := model.AgentsPair{FromAgent: transitionToAgent[conn.FromTransitionId], ToAgent: transitionToAgent[conn.ToTransitionId]}
		key2 := model.AgentsPair{FromAgent: transitionToAgent[conn.ToTransitionId], ToAgent: transitionToAgent[conn.FromTransitionId]}

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

type RatioResult struct {
	agentOne string
	agentTwo string
	ratio    float64
}
