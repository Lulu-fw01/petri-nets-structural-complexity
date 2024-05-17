package woped

import (
	"complexity/internal/reader"
	"complexity/pkg/settings"
	"testing"
)

func TestReadWopedNetHappyPath(t *testing.T) {
	netSettings := settings.SimpleSettings{AgentsToTransitions: make(map[string][]string), SilentTransitions: []string{}}
	newNet, err := reader.ReadNet[Pnml]("testdata/IP-1_ref_model.pnml", &netSettings)
	if err != nil {
		t.Fatalf("Error reading xml: %s", err)
	}
	if len(newNet.Places) != 35 {
		t.Fatalf("Expecting 35 places, got: %d", len(newNet.Places))
	}
	if len(newNet.Transitions) != 31 {
		t.Fatalf("Expecting 31 transitions, got: %d", len(newNet.Transitions))
	}
	if len(newNet.Arcs) != 73 {
		t.Fatalf("Expecting 73 arcs, got: %d", len(newNet.Arcs))
	}

	//places := newNet.Places
	//expectedPlaceIds := []string{"P0", "P1", "P2", "P3", "P4"}
	//actualPlaceIds := make([]string, 5)
	//for i, p := range places {
	//	actualPlaceIds[i] = p.Id
	//}
	//list.CheckStringList(t, expectedPlaceIds, actualPlaceIds)
	//
	//transitions := newNet.Transitions
	//expectedTransitionIds := []string{"T0", "T1", "T2"}
	//actualTransitionIds := make([]string, 5)
	//for i, tr := range transitions {
	//	actualTransitionIds[i] = tr.Id
	//}
	//list.CheckStringList(t, expectedTransitionIds, actualTransitionIds)
	//
	//actualArcs := make([]net.Arc, 8)
	//for i, a := range newNet.Arcs {
	//	actualArcs[i] = *a
	//}
	//expectedArcs := []net.Arc{
	//	{Source: "P0", Target: "T0"},
	//	{Source: "P1", Target: "T1"},
	//	{Source: "P2", Target: "T2"},
	//	{Source: "P3", Target: "T2"},
	//	{Source: "T0", Target: "P1"},
	//	{Source: "T1", Target: "P2"},
	//	{Source: "T1", Target: "P3"},
	//	{Source: "T2", Target: "P4"},
	//}
	//for _, el := range expectedArcs {
	//	if !slices.Contains(actualArcs, el) {
	//		t.Fatalf("There are not al arcs in list. Can't find %s", el)
	//	}
	//}
}
