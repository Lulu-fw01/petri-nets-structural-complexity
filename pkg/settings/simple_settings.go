package settings

import (
	"complexity/pkg/net"
	"fmt"
)

type SimpleSettings struct {
	AgentsToTransitions map[string][]string `json:"agentsToTransitions"`
	SilentTransitions   []string            `json:"silentTransitions"`
}

func (s SimpleSettings) GetTransitionAgent(transition *net.Transition) (*string, error) {
	transitionToAgent := s.GetTransitionToAgentMap([]*net.Transition{transition})
	agent, contains := transitionToAgent[transition.Id]
	if !contains {
		return nil, fmt.Errorf("there are no transition with id %s", transition.Id)
	}
	return &agent, nil
}

func (s SimpleSettings) GetTransitionToAgentMap(transitions []*net.Transition) map[string]string {
	transitionToAgent := make(map[string]string)
	for agent, transitions := range s.AgentsToTransitions {
		for _, t := range transitions {
			transitionToAgent[t] = agent
		}
	}
	resultTransitionToAgentMap := make(map[string]string)
	for _, transition := range transitions {
		resultTransitionToAgentMap[transition.Id] = transitionToAgent[transition.Id]
	}

	return resultTransitionToAgentMap
}

func (s SimpleSettings) IsSilentTransition(transitionId string) bool {
	for _, t := range s.SilentTransitions {
		if t == transitionId {
			return true
		}
	}
	return false
}
