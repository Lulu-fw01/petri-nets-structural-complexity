package settings

import (
	"complexity/pkg/net"
	"fmt"
	"regexp"
)

type RegexpSettings struct {
	AgentsToTransitionRegexp map[string]string   `json:"agentToTransitionRegexp"`
	SilentTransitionsRegexp  string              `json:"silentTransitionRegexp"`
	WeightType               *string             `json:"weightType"`
	AgentToWeight            *map[string]float64 `json:"agentToWeight"`
}

func (r RegexpSettings) GetTransitionAgent(transition *net.Transition) (*string, error) {
	for agent, regexpStr := range r.AgentsToTransitionRegexp {
		match, _ := regexp.MatchString(regexpStr, transition.Id)
		if match {
			return &agent, nil
		}
	}
	return nil, fmt.Errorf("there are no agent  for transition with id %s", transition.Id)
}

func (r RegexpSettings) GetTransitionToAgentMap(transitions []*net.Transition) map[string]string {
	result := make(map[string]string)
	for _, t := range transitions {
		for agent, regexpStr := range r.AgentsToTransitionRegexp {
			match, _ := regexp.MatchString(regexpStr, t.Id)
			if match {
				result[t.Id] = agent
			}
		}
	}
	return result
}

func (r RegexpSettings) IsSilentTransition(transitionId string) bool {
	// todo мб что-то другое вместо этого if.
	if r.SilentTransitionsRegexp == "" {
		return false
	}
	match, _ := regexp.MatchString(r.SilentTransitionsRegexp, transitionId)
	return match
}

func (r RegexpSettings) GetAgents() []string {
	var agents []string
	for a := range r.AgentsToTransitionRegexp {
		agents = append(agents, a)
	}
	return agents
}

func (r RegexpSettings) GetWeightType() *string {
	return r.WeightType
}

func (r RegexpSettings) GetAgentsWeights() *map[string]float64 {
	return r.AgentToWeight
}
