package main

import (
	"complexity/internal/reader"
	"complexity/internal/reader/pipe"
	"complexity/internal/writer"
	"complexity/pkg/algorithm"
	"complexity/pkg/settings"
	"flag"
	"fmt"
	"log"
	"os"
)

const (
	MetricTypeFlag     = "metrics"
	BatchProcessFlag   = "batch"
	SettingsTypeFlag   = "settings-type"
	SettingsPathFlag   = "settings"
	NetPathFlag        = "net"
	FileOutputFlag     = "file"
	SimpleSettingsType = "simple"
	RegexpSettingsType = "regexp"
	AllMetricType      = "all"
	V1MetricType       = "v1"
	V2MetricType       = "v2"
)

func main() {
	metric := flag.String(MetricTypeFlag, AllMetricType, "metric version")
	isBatchProcess := flag.Bool(BatchProcessFlag, false, "Process batch of nets")
	settingsType := flag.String(SettingsTypeFlag, SimpleSettingsType, "settings type (simple or regexp)")
	settingsPath := flag.String(SettingsPathFlag, "", "net settings")
	netPath := flag.String(NetPathFlag, "", "net description")
	filePath := flag.String(FileOutputFlag, "", "path to output file")
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

	if *isBatchProcess {
		batchFlow(*netPath, *metric, netSettings, output)
	} else {
		standardFlow(*netPath, *metric, netSettings, output)
	}
}

func standardFlow(netPath, metric string, netSettings settings.Settings, fn writer.OutputFunc) {
	netToProcess, err := reader.ReadNet[pipe.Pnml](netPath, netSettings)
	if err != nil {
		fmt.Printf("Erorr: %s", err)
		return
	}
	var message string
	switch metric {
	case AllMetricType:
		c1 := algorithm.CountMetricVersion1(netToProcess, netSettings)
		c2 := algorithm.CountCharacteristicV2(netToProcess, netSettings)
		message = getCharacteristicV1Message(c1) + getCharacteristicV2Message(c2)
	case V1MetricType:
		c := algorithm.CountMetricVersion1(netToProcess, netSettings)
		message = getCharacteristicV1Message(c)
	case V2MetricType:
		c := algorithm.CountCharacteristicV2(netToProcess, netSettings)
		message = getCharacteristicV2Message(c)
	default:
		println("Incorrect metric type.")
		return
	}
	fn(message)
}

func batchFlow(dirPath, metric string, netSettings settings.Settings, fn writer.OutputFunc) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("Error reading directory: %s", err)
		return
	}

	// Iterate over the files.
	for _, file := range files {
		// Check if the file is a directory.
		if !file.IsDir() {
			// Output the file name.
			fn(file.Name())
			standardFlow(dirPath+"/"+file.Name(), metric, netSettings, fn)
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

func getCharacteristicV1Message(value float64) string {
	return fmt.Sprintf("Characteristic 1 equals %f\n", value)
}

func getCharacteristicV2Message(value float64) string {
	return fmt.Sprintf("Characteristic 2 equals %f\n", value)
}

func getOutputFunction(filePath string) (writer.OutputFunc, *os.File, error) {
	if filePath == "" {
		// return nil file if console output.
		return consoleOutput, nil, nil
	}

	f, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Can't create file %s", err)
		return consoleOutput, nil, err
	}

	return func(text string) {
		data := []byte(text)

		_, err = f.Write(data)

		if err != nil {
			log.Fatalf("Error writing to file %s", err)
		}
	}, f, nil
}

func consoleOutput(text string) {
	fmt.Println(text)
}
