package test

import (
	"complexity/internal/reader"
	"complexity/internal/reader/pipe"
	"complexity/pkg/net"
	"complexity/pkg/settings"
	"testing"
)

func ReadSettingsAndNet(t *testing.T, settingsPath, netPath string) (settings.Settings, *net.PetriNet) {
	netSettings, err := settings.ReadSettings[settings.SimpleSettings](settingsPath)
	if err != nil {
		t.Fatalf("Error reading settings from %s. err: %s", settingsPath, err)
	}
	newNet, err := reader.ReadNet[pipe.Pnml](netPath, netSettings)
	if err != nil {
		t.Fatalf("Error reading net from %s. err: %s", netPath, err)
	}
	return netSettings, newNet
}

func ReadSettings[S settings.Settings](t *testing.T, settingsPath string) settings.Settings {
	netSettings, err := settings.ReadSettings[S](settingsPath)
	if err != nil {
		t.Fatalf("Error reading settings from %s. err: %s", settingsPath, err)
	}
	return *netSettings
}

func ReadNet(t *testing.T, netSettings settings.Settings, netPath string) *net.PetriNet {
	newNet, err := reader.ReadNet[pipe.Pnml](netPath, netSettings)
	if err != nil {
		t.Fatalf("Error reading net from %s. err: %s", netPath, err)
	}
	return newNet
}
