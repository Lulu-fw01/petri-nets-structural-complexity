package reader

import (
	"complexity/utils/test/list"
	"testing"
)

func TestReadSettingsHappyPath(t *testing.T) {
	settings, err := ReadSettings("testdata/test-settings-happy-path.json")
	if err != nil {
		t.Fatalf("Error reading settings from testdata/test-settings-happy-path.json. err: %s", err)
	}

	agentsToTransitions := settings.AgentsToTransitions
	if len(agentsToTransitions) != 2 {
		t.Fatalf("Incorrect number of keys in map. Actual value: %d", len(agentsToTransitions))
	}

	a1 := agentsToTransitions["a1"]
	a2 := agentsToTransitions["a2"]
	if len(a1) != 3 {
		t.Fatalf("Incorrect number of transitions for a1.")
	}

	if len(a2) != 3 {
		t.Fatalf("Incorrect number of transitions for a2.")
	}

	a1Expected := []string{"t1", "t2", "t3"}
	a2Expected := []string{"q1", "q2", "q3"}
	list.CheckStringList(t, a1Expected, a1)
	list.CheckStringList(t, a2Expected, a2)
}
