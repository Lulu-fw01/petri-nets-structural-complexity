package pipe

import "encoding/xml"

type Place struct {
	XmlName xml.Name `xml:"place"`
	Id      string   `xml:"id,attr"`
}

type Transition struct {
	XMLName xml.Name `xml:"transition"`
	Id      string   `xml:"id,attr"`
}

type Arc struct {
	XMLName xml.Name `xml:"arc"`
	Source  string   `xml:"source,attr"`
	Target  string   `xml:"target,attr"`
}

type PetriNet struct {
	XMLName     xml.Name      `xml:"net"`
	Places      []*Place      `xml:"place"`
	Transitions []*Transition `xml:"transition"`
	Arcs        []*Arc        `xml:"arc"`
}

type Pnml struct {
	XMLName xml.Name  `xml:"pnml"`
	Net     *PetriNet `xml:"net"`
}
