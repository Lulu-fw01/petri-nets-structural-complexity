package algorithm

import (
	"complexity/utils/assertions"
	testUtils "complexity/utils/test"
	"testing"
)

func TestCountRatiosFor2AgentsHappyPath(t *testing.T) {
	netSettings, newNet := testUtils.ReadSettingsAndNet(t, "testdata/2-agents-settings.json", "testdata/2-agents.xml")

	result := CountRatios(newNet, netSettings)

	if len(result) != 1 {
		t.Fatalf("Expecte size of result list 1, actual: %d", len(result))
	}

	firstMetric := result[0]
	assertions.AssertMetric(t, 0.625, firstMetric.ratio)
}

// 2 agents, no connections.
func TestCountMetricForNetWithNoChannels(t *testing.T) {
	netSettings, newNet := testUtils.ReadSettingsAndNet(t, "testdata/common-settings.json", "testdata/no-channels-net.xml")

	result := CountRatios(newNet, netSettings)

	// if there are no connections between 2 list will be empty.
	if len(result) != 0 {
		t.Fatalf("Expecte size of result list 0, actual: %d", len(result))
	}
}

// 2 agents^ 1 channel, 1 connection.
func TestCountMetricForNetWith1ConnectionBetweenAgents(t *testing.T) {
	netSettings, newNet := testUtils.ReadSettingsAndNet(t, "testdata/common-settings.json", "testdata/2-agents-v2.xml")

	result := CountRatios(newNet, netSettings)

	if len(result) != 1 {
		t.Fatalf("Expecte size of result list 1, actual: %d", len(result))
	}

	firstMetric := result[0]
	assertions.AssertMetric(t, 0.8, firstMetric.ratio)
}

// 2 agents, 1 channel, 2 connections.
func TestCountMetricForNetWith2ConnectionsBetweenAgents(t *testing.T) {
	netSettings, newNet := testUtils.ReadSettingsAndNet(t, "testdata/common-settings.json", "testdata/2-agents-v3.xml")

	result := CountRatios(newNet, netSettings)

	if len(result) != 1 {
		t.Fatalf("Expecte size of result list 1, actual: %d", len(result))
	}

	firstMetric := result[0]
	assertions.AssertMetric(t, 0.666667, firstMetric.ratio)
}

// 2 agents, 2 channels, 2 and 2 connections.
func TestCountMetricForNetWith4ConnectionsBetweenAgentsAnd2Channels(t *testing.T) {
	netSettings, newNet := testUtils.ReadSettingsAndNet(t, "testdata/common-settings.json", "testdata/2-agents-v4.xml")

	result := CountRatios(newNet, netSettings)

	if len(result) != 1 {
		t.Fatalf("Expecte size of result list 1, actual: %d", len(result))
	}

	firstMetric := result[0]
	assertions.AssertMetric(t, 0.571429, firstMetric.ratio)
}

// 2 agents, 2 channels, 2 and 4 connections.
func TestCountMetricForNetWith5ConnectionsBetweenAgentsAnd2Channels(t *testing.T) {
	netSettings, newNet := testUtils.ReadSettingsAndNet(t, "testdata/common-settings.json", "testdata/2-agents-v5.xml")

	result := CountRatios(newNet, netSettings)

	if len(result) != 1 {
		t.Fatalf("Expecte size of result list 1, actual: %d", len(result))
	}

	firstMetric := result[0]
	assertions.AssertMetric(t, 0.5, firstMetric.ratio)
}
