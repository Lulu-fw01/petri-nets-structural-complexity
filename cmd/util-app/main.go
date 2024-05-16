package main

import (
	"complexity/internal/reader"
	"complexity/internal/reader/pipe"
	"complexity/internal/reader/woped"
	w "complexity/internal/writer"
	"complexity/pkg/algorithm"
	"complexity/pkg/net"
	"complexity/pkg/settings"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

const (
	MetricTypeFlag       = "metrics"
	BatchProcessFlag     = "batch"
	SettingsTypeFlag     = "settings-type"
	SettingsPathFlag     = "settings"
	NetPathFlag          = "net"
	FileOutputFlag       = "file"
	SourceTypeFlag       = "source"
	SimpleSettingsType   = "simple"
	RegexpSettingsType   = "regexp"
	AllMetricType        = "all"
	V1CharacteristicType = "v1"
	V2CharacteristicType = "v2"
	V3CharacteristicType = "v3"
	PipeSource           = "pipe"
	WopedSource          = "woped"
	PromSource           = "prom"
)

func main() {
	metric := flag.String(MetricTypeFlag, AllMetricType, "metric version")
	isBatchProcess := flag.Bool(BatchProcessFlag, false, "Process batch of nets")
	settingsType := flag.String(SettingsTypeFlag, SimpleSettingsType, "settings type (simple or regexp)")
	settingsPath := flag.String(SettingsPathFlag, "", "net settings")
	netPath := flag.String(NetPathFlag, "", "net description")
	filePath := flag.String(FileOutputFlag, "", "path to output file")
	sourceType := flag.String(SourceTypeFlag, PipeSource, "source of the net: pipe, woped, prom")
	flag.Parse()

	validateSettingsPath(*settingsPath)
	validateNetPath(*netPath)

	netSettings, err := getSettings(*settingsPath, *settingsType)
	if err != nil {
		log.Fatalf("Erorr: %s", err)
		return
	}

	output, fileOutput, err := getOutputFunction(*filePath)
	if err != nil {
		return
	}

	if fileOutput != nil {
		defer fileOutput.Close()
	}

	output([][]string{{"value", "type", "path-to-net"}})
	if *isBatchProcess {
		batchFlow(*netPath, *metric, netSettings, output, *sourceType)
	} else {
		standardFlow(*netPath, *metric, netSettings, output, *sourceType)
	}
}

func standardFlow(netPath, metric string, netSettings settings.Settings, fn w.OutputFunc, sourceType string) {
	netToProcess, err := getNet(netPath, netSettings, sourceType)
	if err != nil {
		fmt.Printf("Erorr: %s", err)
		return
	}
	var records [][]string
	switch metric {
	case AllMetricType:
		c1 := algorithm.CountCharacteristicV1(netToProcess, netSettings)
		c2 := algorithm.CountCharacteristicV2(netToProcess, netSettings)
		c3 := algorithm.CountCharacteristicV3(netToProcess, netSettings)
		records = append(records,
			getCharacteristicV1Record(c1, netPath),
			getCharacteristicV2Record(c2, netPath),
			getCharacteristicV3Record(c3, netPath))
	case V1CharacteristicType:
		c := algorithm.CountCharacteristicV1(netToProcess, netSettings)
		records = append(records, getCharacteristicV1Record(c, netPath))
	case V2CharacteristicType:
		c := algorithm.CountCharacteristicV2(netToProcess, netSettings)
		records = append(records, getCharacteristicV2Record(c, netPath))
	case V3CharacteristicType:
		c := algorithm.CountCharacteristicV3(netToProcess, netSettings)
		records = append(records, getCharacteristicV3Record(c, netPath))
	default:
		println("Incorrect metric type.")
		return
	}
	fn(records)
}

func batchFlow(dirPath, metric string, netSettings settings.Settings, fn w.OutputFunc, sourceType string) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("Error reading directory: %s", err)
		return
	}

	// Iterate over the files.
	for _, file := range files {
		// Check if the file is a directory.
		if !file.IsDir() {
			standardFlow(dirPath+"/"+file.Name(), metric, netSettings, fn, sourceType)
		}
	}
}

func validateSettingsPath(path string) {
	if path == "" {
		fmt.Println("Please provide path to net settings.")
		flag.Usage()
		return
	}
}

func validateNetPath(path string) {
	if path == "" {
		fmt.Println("Please provide path to net description.")
		flag.Usage()
		return
	}
}

func getSettings(path string, settingsType string) (settings.Settings, error) {
	switch settingsType {
	case SimpleSettingsType:
		return settings.ReadSettings[settings.SimpleSettings](path)
	case RegexpSettingsType:
		return settings.ReadSettings[settings.RegexpSettings](path)
	default:
		return nil, fmt.Errorf("wrong settings type: %s", settingsType)
	}
}

func getCharacteristicV1Record(value float64, netPath string) []string {
	return []string{fmt.Sprintf("%f", value), "v1", netPath}
}

func getCharacteristicV2Record(value float64, netPath string) []string {
	return []string{fmt.Sprintf("%f", value), "v2", netPath}
}

func getCharacteristicV3Record(value float64, netPath string) []string {
	return []string{fmt.Sprintf("%f", value), "v3", netPath}
}

func getOutputFunction(filePath string) (w.OutputFunc, *os.File, error) {
	if filePath == "" {
		// return nil file if console output.
		return consoleOutput, nil, nil
	}

	f, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Can't create file %s", err)
		return consoleOutput, nil, err
	}
	writer := csv.NewWriter(f)

	return func(records [][]string) {
		err := writer.WriteAll(records)
		if err != nil {
			log.Fatalf("Error writing to file %s", err)
		}
	}, f, nil
}

func consoleOutput(records [][]string) {
	for _, record := range records {
		for _, elem := range record {
			fmt.Print(elem + " | ")
		}
		fmt.Println()
	}
}

func getNet(netPath string, netSettings settings.Settings, sourceType string) (*net.PetriNet, error) {
	switch sourceType {
	case PromSource:
		//return reader.ReadNet[pipe.](netPath, netSettings)
		return nil, fmt.Errorf("Not implimented")
	case WopedSource:
		return reader.ReadNet[woped.Pnml](netPath, netSettings)
	default:
		return reader.ReadNet[pipe.Pnml](netPath, netSettings)
	}
}
