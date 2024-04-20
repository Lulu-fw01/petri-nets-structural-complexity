package reader

import (
	"complexity/pkg/net"
	"complexity/pkg/settings"
)

// NetConverter interface for converting external nets to lib net.
type NetConverter interface {
	Convert(settings settings.Settings) *net.PetriNet
}
