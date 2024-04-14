package settings

import (
	"complexity/pkg/net"
	"testing"
)

func TestReadRegexpSettingsHappyPath(t *testing.T) {
	settings, err := ReadSettings[RegexpSettings]("testdata/test-regexp-settings.json")
	if err != nil {
		t.Fatalf("Error reading settings from testdata/test-settings-happy-path.json. err: %s", err)
	}

	transitions := []*net.Transition{
		{"a1-t1", false},
		{"a1-t2", false},
		{"a2-t1", false},
		{"t1", false},
		{"t2", false},
	}

	transitionToAgent := settings.GetTransitionToAgentMap(transitions)

	expected := map[string]string{
		"a1-t1": "a1",
		"a1-t2": "a1",
		"a2-t1": "a2",
		"t1":    "a2",
		"t2":    "a2",
	}

	assertMap(t, transitionToAgent, expected)

	isSilent := settings.IsSilentTransition("s1")
	if !isSilent {
		t.Errorf("transition %s should be silent, but it is not", "s1")
	}

	isSilent = settings.IsSilentTransition("s3")
	if !isSilent {
		t.Errorf("transition %s should be silent, but it is not", "s3")
	}

	isSilent = settings.IsSilentTransition("s4")
	if isSilent {
		t.Errorf("transition %s should not be silent, but it is", "s4")
	}
}

func assertMap(t *testing.T, actual map[string]string, expected map[string]string) {
	if len(actual) != len(expected) {
		t.Fatalf("Actual map has different size: %d, expected: %d", len(actual), len(expected))
	}
	for key, value := range expected {
		if actual[key] != value {
			t.Errorf("Expected key %s to have value %s, but got %s", key, value, actual[key])
		}
	}
}
