package algorithm

import (
	"complexity/internal/reader"
	"complexity/internal/reader/pipe"
	"complexity/pkg/settings"
	"complexity/utils/assertions"
	testUtils "complexity/utils/test"
	"testing"
)

func TestCountRatiosFor2AgentsHappyPath(t *testing.T) {
	netSettings, newNet := testUtils.ReadSettingsAndNet(t, "testdata/2-agents-settings.json", "testdata/2-agents.xml")

	result := CountMetricVersion1(newNet, netSettings)

	assertions.AssertMetric(t, 0.625, result)
}

// 2 agents, no connections.
func TestCountMetricForNetWithNoChannels(t *testing.T) {
	netSettings, newNet := testUtils.ReadSettingsAndNet(t, "testdata/common-settings.json", "testdata/no-channels-net.xml")

	result := CountMetricVersion1(newNet, netSettings)

	assertions.AssertMetric(t, 1.0, result)
}

// 2 agents^ 1 channel, 1 connection.
func TestCountMetricForNetWith1ConnectionBetweenAgents(t *testing.T) {
	netSettings, newNet := testUtils.ReadSettingsAndNet(t, "testdata/common-settings.json", "testdata/2-agents-v2.xml")

	result := CountMetricVersion1(newNet, netSettings)

	assertions.AssertMetric(t, 0.8, result)
}

// 2 agents, 1 channel, 2 connections.
func TestCountMetricForNetWith2ConnectionsBetweenAgents(t *testing.T) {
	netSettings, newNet := testUtils.ReadSettingsAndNet(t, "testdata/common-settings.json", "testdata/2-agents-v3.xml")

	result := CountMetricVersion1(newNet, netSettings)

	assertions.AssertMetric(t, 0.666667, result)
}

// 2 agents, 2 channels, 2 and 2 connections.
func TestCountMetricForNetWith4ConnectionsBetweenAgentsAnd2Channels(t *testing.T) {
	netSettings, newNet := testUtils.ReadSettingsAndNet(t, "testdata/common-settings.json", "testdata/2-agents-v4.xml")

	result := CountMetricVersion1(newNet, netSettings)

	assertions.AssertMetric(t, 0.5, result)
}

// 2 agents, 2 channels, 2 and 4 connections.
func TestCountMetricForNetWith5ConnectionsBetweenAgentsAnd2Channels(t *testing.T) {
	netSettings, newNet := testUtils.ReadSettingsAndNet(t, "testdata/common-settings.json", "testdata/2-agents-v5.xml")

	result := CountMetricVersion1(newNet, netSettings)

	assertions.AssertMetric(t, 0.4, result)
}

func TestCountMetricV1ForNetWith14CausalConnections(t *testing.T) {
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
	result := CountMetricVersion1(newNet, netSettings)

	assertions.AssertMetric(t, 0.428571, result)
}
