package pipe

import (
	"complexity/internal/reader"
	"complexity/pkg/net"
	"complexity/pkg/settings"
	"complexity/utils/test/list"
	"slices"
	"testing"
)

func TestReadPipeNetHappyPath(t *testing.T) {
	netSettings := settings.SimpleSettings{AgentsToTransitions: make(map[string][]string), SilentTransitions: []string{}}
	newNet, err := reader.ReadNet[Pnml]("testdata/happy-path.xml", &netSettings)
	if err != nil {
		t.Fatalf("Error reading xml: %s", err)
	}
	if len(newNet.Places) != 5 {
		t.Fatalf("Expecting 5 places, got: %d", len(newNet.Places))
	}
	if len(newNet.Transitions) != 3 {
		t.Fatalf("Expecting 3 transitions, got: %d", len(newNet.Transitions))
	}
	if len(newNet.Arcs) != 8 {
		t.Fatalf("Expecting 8 arcs, got: %d", len(newNet.Arcs))
	}

	places := newNet.Places
	expectedPlaceIds := []string{"P0", "P1", "P2", "P3", "P4"}
	actualPlaceIds := make([]string, 5)
	for i, p := range places {
		actualPlaceIds[i] = p.Id
	}
	list.CheckStringList(t, expectedPlaceIds, actualPlaceIds)

	transitions := newNet.Transitions
	expectedTransitionIds := []string{"T0", "T1", "T2"}
	actualTransitionIds := make([]string, 5)
	for i, tr := range transitions {
		actualTransitionIds[i] = tr.Id
	}
	list.CheckStringList(t, expectedTransitionIds, actualTransitionIds)

	actualArcs := make([]net.Arc, 8)
	for i, a := range newNet.Arcs {
		actualArcs[i] = *a
	}
	expectedArcs := []net.Arc{
		getArc("P0", "T0"),
		getArc("P1", "T1"),
		getArc("P2", "T2"),
		getArc("P3", "T2"),
		getArc("T0", "P1"),
		getArc("T1", "P2"),
		getArc("T1", "P3"),
		getArc("T2", "P4"),
	}
	for _, el := range expectedArcs {
		if !slices.Contains(actualArcs, el) {
			t.Fatalf("There are not al arcs in list. Can't find %s", el)
		}
	}
}

func getArc(source string, target string) net.Arc {
	return net.Arc{
		Source: source,
		Target: target,
	}
}
