package algorithm

import (
	"complexity/pkg/net"
	"complexity/pkg/settings"
)

func CountMetricVersion1(net *net.PetriNet, settings settings.Settings) float64 {
	ratios := CountRatios(net, settings)
	result := 0.0
	for _, r := range ratios {
		result += r.ratio
	}
	return result
}

func CountRatios(net *net.PetriNet, settings settings.Settings) []RatioResult {
	transitionToAgent := settings.GetTransitionToAgentMap(net.Transitions)
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
				ratio:    1 - float64(arcsCount)/float64(connectionsCount),
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
