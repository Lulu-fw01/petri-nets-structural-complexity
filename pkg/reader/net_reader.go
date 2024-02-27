package reader

import (
	"complexity/pkg/net"
	"encoding/xml"
	"golang.org/x/net/html/charset"
	"log"
	"os"
)

func ReadNet(path string) (*net.PetriNet, error) {
	netFile, err := os.Open(path)
	if err != nil {
		log.Fatalf("Could not open file with net. Error: %s", err)
		return nil, err
	}
	defer netFile.Close()

	xmlDecoder := xml.NewDecoder(netFile)
	xmlDecoder.CharsetReader = charset.NewReaderLabel

	var newNet net.Pnml
	err = xmlDecoder.Decode(&newNet)
	if err != nil {
		log.Fatalf("Error decoding xml file: %s", err)
		return nil, err
	}

	return newNet.Net, nil
}
