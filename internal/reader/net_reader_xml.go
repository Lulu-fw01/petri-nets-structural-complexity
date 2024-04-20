package reader

import (
	"complexity/pkg/net"
	"complexity/pkg/settings"
	"encoding/xml"
	"golang.org/x/net/html/charset"
	"log"
	"os"
)

func ReadNet[N NetConverter](path string, netSettings settings.Settings) (*net.PetriNet, error) {
	pnml, err := readNet[N](path)
	if err != nil {
		return nil, err
	}
	return (*pnml).Convert(netSettings), nil
}

func readNet[N NetConverter](path string) (*N, error) {
	netFile, err := os.Open(path)
	if err != nil {
		log.Fatalf("Could not open file with net. Error: %s", err)
		return nil, err
	}
	defer netFile.Close()

	xmlDecoder := xml.NewDecoder(netFile)
	xmlDecoder.CharsetReader = charset.NewReaderLabel

	var newNet N
	err = xmlDecoder.Decode(&newNet)
	if err != nil {
		log.Fatalf("Error decoding xml file: %s", err)
		return nil, err
	}

	return &newNet, nil
}
