package test

import (
	"complexity/internal/reader/pipe"
	"complexity/pkg/net"
	"complexity/pkg/settings"
	"testing"
)

func ReadSettingsAndNet(t *testing.T, settingsPath, netPath string) (*settings.Settings, *net.PetriNet) {
	netSettings, err := settings.ReadSettings(settingsPath)
	if err != nil {
		t.Fatalf("Error reading settings from %s. err: %s", settingsPath, err)
	}
	newNet, err := pipe.ReadNet(netPath, netSettings.SilentTransitions)
	if err != nil {
		t.Fatalf("Error reading net from %s. err: %s", netPath, err)
	}
	return netSettings, newNet
}
