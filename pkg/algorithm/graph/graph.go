package graph

import "complexity/pkg/net"

// GetGraph constructs a graph representation from a slice of arcs.
// Each key in the returned map is an element ID, and the value is a slice of element IDs
// that are connected to the key via input arcs.
func GetGraph(
	arcs []*net.Arc) map[string][]string {
	elements := make(map[string][]string)

	for _, arc := range arcs {
		source := arc.Source
		target := arc.Target
		elements[source] = append(elements[source], target)
	}
	return elements
}
