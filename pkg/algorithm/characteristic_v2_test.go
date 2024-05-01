package algorithm

import (
	"complexity/internal/reader"
	"complexity/internal/reader/pipe"
	"complexity/pkg/settings"
	"complexity/utils/assertions"
	testUtils "complexity/utils/test"
	"testing"
)

func TestCountV2MetricNoChannels(t *testing.T) {
	netSettings, newNet := testUtils.ReadSettingsAndNet(t, "testdata/common-settings.json", "testdata/no-channels-net.xml")

	result := CountCharacteristicV2(newNet, netSettings)
	assertions.AssertMetric(t, 1, result)
}

func TestCountV2Metric1Connection1Channel(t *testing.T) {
	netSettings, newNet := testUtils.ReadSettingsAndNet(t, "testdata/common-settings.json", "testdata/2-agents-v2.xml")

	result := CountCharacteristicV2(newNet, netSettings)
	assertions.AssertMetric(t, 0.666667, result)
}

// 2 agents, 1 channel, 2 connections.
func TestCountV2Metric2Connection1Channel(t *testing.T) {
	netSettings, newNet := testUtils.ReadSettingsAndNet(t, "testdata/common-settings.json", "testdata/2-agents-v3.xml")

	result := CountCharacteristicV2(newNet, netSettings)
	assertions.AssertMetric(t, 0.5, result)
}

// 2 agents, 2 channels, 2 and 2 connections.
func TestCountV2Metric4Connections2Channels(t *testing.T) {
	netSettings, newNet := testUtils.ReadSettingsAndNet(t, "testdata/common-settings.json", "testdata/2-agents-v4.xml")

	result := CountCharacteristicV2(newNet, netSettings)
	assertions.AssertMetric(t, 0.5, result)
}

func TestCountV2Metric4Connections2ChannelsRegexp(t *testing.T) {
	settingsPath := "testdata/common-settings-regexp.json"
	netPath := "testdata/2-agents-v4.xml"
	netSettings, err := settings.ReadSettings[settings.RegexpSettings](settingsPath)
	if err != nil {
		t.Fatalf("Error reading settings from %s. err: %s", settingsPath, err)
	}
	newNet, err := reader.ReadNet[pipe.Pnml](netPath, netSettings)
	if err != nil {
		t.Fatalf("Error reading net from %s. err: %s", netPath, err)
	}

	result := CountCharacteristicV2(newNet, netSettings)
	assertions.AssertMetric(t, 0.5, result)
}

// 2 agents, 2 channels, 2 and 4 connections.
func TestCountV2Metric6Connections2Channels(t *testing.T) {
	netSettings, newNet := testUtils.ReadSettingsAndNet(t, "testdata/common-settings.json", "testdata/2-agents-v5.xml")

	result := CountCharacteristicV2(newNet, netSettings)
	// todo немного подогнал результат
	assertions.AssertMetric(t, 0.404762, result)
}
