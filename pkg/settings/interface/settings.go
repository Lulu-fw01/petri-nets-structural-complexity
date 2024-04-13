package _interface

import "complexity/pkg/net"

type Settings interface {
	GetTransitionAgent(*net.Transition) (*string, error)
	GetTransitionToAgentMap([]*net.Transition) map[string]string
}
