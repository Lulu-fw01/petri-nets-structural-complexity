package algorithm

import (
	"complexity/internal/reader"
	"complexity/internal/reader/pipe"
	"complexity/pkg/settings"
	"slices"
	"testing"
)

func TestFindCausalConnectionsHappyPath(t *testing.T) {
	netSettings, err := settings.ReadSettings[settings.SimpleSettings]("testdata/2-agents-settings.json")
	if err != nil {
		t.Fatalf("Error reading settings from testdata/2-agents-settings.json. err: %s", err)
	}
	newNet, err := reader.ReadNet[pipe.Pnml]("testdata/2-agents.xml", netSettings)
	if err != nil {
		t.Fatalf("Error reading net from testdata/2-agents.xml. err: %s", err)
	}

	connections := FindCausalConnections(newNet)
	actualConnections := dereferenceConnections(connections)

	assertConnectionsCount(t, actualConnections, 8)

	expectedConnections := []CausalConnection{
		{"T1", "T3"},
		{"T1", "T4"},
		{"T1", "Q1"},
		{"T4", "Q2"},
		{"Q1", "Q2"},
		{"Q2", "Q4"},
		{"Q1", "Q5"},
		{"Q5", "T3"},
	}

	assertConnectionsExist(t, actualConnections, expectedConnections)
}

func dereferenceConnections(connections []*CausalConnection) []CausalConnection {
	var actualConnections []CausalConnection
	for _, c := range connections {
		actualConnections = append(actualConnections, *c)
	}
	return actualConnections
}

func assertConnectionsCount(t *testing.T, actualConnections []CausalConnection, expectedCount int) {
	if len(actualConnections) != expectedCount {
		t.Fatalf("Incorrect number of causal connections, expected: %d, actual: %d", expectedCount, len(actualConnections))
	}
}

func assertConnectionsExist(t *testing.T, actualConnections []CausalConnection, expectedConnections []CausalConnection) {
	for _, el := range expectedConnections {
		if !slices.Contains(actualConnections, el) {
			t.Fatalf("Missing connection: %s", el)
		}
	}
}

func TestFindCausalConnectionsHappyPath3Agents(t *testing.T) {
	// todo implement.
}

func TestCausalConnection2IputArcsForChannel(t *testing.T) {
	settingsPath := "testdata/metric-1/common-settings.json"
	netPath := "testdata/metric-1/8.xml"
	netSettings, err := settings.ReadSettings[settings.RegexpSettings](settingsPath)
	if err != nil {
		t.Fatalf("Error reading settings from %s. err: %s", settingsPath, err)
	}
	newNet, err := reader.ReadNet[pipe.Pnml](netPath, netSettings)
	if err != nil {
		t.Fatalf("Error reading net from %s. err: %s", netPath, err)
	}

	connections := FindCausalConnections(newNet)
	if len(connections) != 14 {
		t.Fatalf("Expected %d causal connections got %d", 14, len(connections))
	}
}
