package settings

import (
	"complexity/pkg/net"
	"fmt"
)

type Settings struct {
	AgentsToTransitions map[string][]string `json:"agentsToTransitions"`
	SilentTransitions   []string            `json:"silentTransitions"`
}

func (s *Settings) GetTransitionAgent(transition *net.Transition) (*string, error) {
	transitionToAgent := s.GetTransitionToAgentMap([]*net.Transition{transition})
	agent, contains := transitionToAgent[transition.Id]
	if !contains {
		return nil, fmt.Errorf("there are no transition with id %s", transition.Id)
	}
	return &agent, nil
}

func (s *Settings) GetTransitionToAgentMap(transitions []*net.Transition) map[string]string {
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
