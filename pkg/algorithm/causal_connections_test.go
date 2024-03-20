package algorithm

import (
	"complexity/internal/reader/pipe"
	"complexity/pkg/settings"
	"testing"
)

func TestCountRatiosFor2AgentsHappyPath(t *testing.T) {
	netSettings, err := settings.ReadSettings("testdata/2-agents-settings.json")
	if err != nil {
		t.Fatalf("Error reading settings from testdata/2-agents-settings.json. err: %s", err)
	}
	newNet, err := pipe.ReadNet("testdata/2-agents.xml", netSettings.SilentTransitions)
	if err != nil {
		t.Fatalf("Error reading net from testdata/2-agents.xml. err: %s", err)

	}

	result := CountRatios(newNet, netSettings)

	if len(result) != 1 {
		t.Fatalf("Expecte size of result list 1, actual: %d", len(result))
	}

	firstMetric := result[0]

	if firstMetric.ratio != 0.375 {
		t.Fatalf("Wrong metric, expected 0.375, actual: %f", firstMetric.ratio)
	}
}

func TestCountRatiosFor3AgentsHappyPath(t *testing.T) {

}
