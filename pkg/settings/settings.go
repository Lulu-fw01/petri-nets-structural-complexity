package settings

import "complexity/pkg/net"

type Settings interface {
	GetTransitionAgent(*net.Transition) (*string, error)
	GetTransitionToAgentMap(transitions []*net.Transition) map[string]string
	IsSilentTransition(string) bool
	GetAgents() []string
	GetWeightType() *string
	GetAgentsWeights() *map[string]float64
}
