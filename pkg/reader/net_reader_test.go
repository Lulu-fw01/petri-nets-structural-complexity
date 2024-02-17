package reader

import (
	"testing"
)

func TestReadNetHappyPath(t *testing.T) {
	net, err := ReadNet("testdata/happy-path.xml")
	if err != nil {
		t.Fatalf("Error reading xml: %s", err)
	}

	if len(net.Places) != 5 {
		t.Fatalf("Expecting 5 places, got: %d", len(net.Places))
	}

	if len(net.Transitions) != 3 {
		t.Fatalf("Expecting 3 transitions, got: %d", len(net.Transitions))
	}

	if len(net.Arcs) != 8 {
		t.Fatalf("Expecting 8 arcs, got: %d", len(net.Arcs))
	}

}
