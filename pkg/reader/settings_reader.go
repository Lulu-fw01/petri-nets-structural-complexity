package reader

import (
	"encoding/json"
	"log"
	"os"
)

type Settings struct {
	AgentsToTransitions map[string][]string `json:"agentsToTransitions"`
}

func ReadSettings(path string) (*Settings, error) {
	settingJson, err := os.Open(path)
	if err != nil {
		log.Fatalf("Could not open file with net. Error: %s", err)
		return nil, err
	}
	defer settingJson.Close()

	jsonDecoder := json.NewDecoder(settingJson)

	var settings Settings
	err = jsonDecoder.Decode(&settings)
	if err != nil {
		log.Fatalf("Error decoding json file: %s", err)
		return nil, err
	}
	return &settings, nil
}
