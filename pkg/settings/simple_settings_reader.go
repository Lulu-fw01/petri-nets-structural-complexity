package settings

import (
	"encoding/json"
	"log"
	"os"
)

func ReadSimpleSettings(path string) (*SimpleSettings, error) {
	settingJson, err := os.Open(path)
	if err != nil {
		log.Fatalf("Could not open file with settings. Error: %s", err)
		return nil, err
	}
	defer settingJson.Close()

	jsonDecoder := json.NewDecoder(settingJson)

	var settings SimpleSettings
	err = jsonDecoder.Decode(&settings)
	if err != nil {
		log.Fatalf("Error decoding json file: %s", err)
		return nil, err
	}
	return &settings, nil
}
