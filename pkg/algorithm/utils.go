package algorithm

import "complexity/pkg/net"

// return map where key is place id and Place as value.
func getPlacesMap(places []*net.Place) map[string]*net.Place {
	placesById := make(map[string]*net.Place)
	for _, p := range places {
		placesById[p.Id] = p
	}
	return placesById
}
