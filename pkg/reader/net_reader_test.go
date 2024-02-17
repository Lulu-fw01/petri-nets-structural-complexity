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
}
