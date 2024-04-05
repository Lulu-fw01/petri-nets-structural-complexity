package algorithm

import (
	testUtils "complexity/utils/test"
	"complexity/utils/test/list"
	"testing"
)

func TestFindChannelsHappyPath(t *testing.T) {
	netSettings, newNet := testUtils.ReadSettingsAndNet(t, "testdata/common-settings.json", "testdata/2-agents-v5.xml")

	channels := FindChannels(newNet, netSettings)

	if len(channels) != 2 {
		t.Fatalf("Incorrect number of causal connections, expected: %d, actual: %d", 2, len(channels))
	}

	var ids []string
	for _, c := range channels {
		ids = append(ids, c.PlaceId)
	}

	list.CheckStringList(t, []string{"a", "b"}, ids)
}
