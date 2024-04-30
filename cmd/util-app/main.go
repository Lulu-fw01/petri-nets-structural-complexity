package main

import (
	"complexity/internal/reader"
	"complexity/internal/reader/pipe"
	"complexity/internal/writer"
	"complexity/pkg/algorithm"
	"complexity/pkg/net"
	"complexity/pkg/settings"
	"flag"
	"fmt"
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
		fmt.Printf("Erorr: %s", err)
		return
	}

	output := getOutputFunction(*filePath)

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
	switch metric {
	case AllMetricType:
		printMetricV1(netToProcess, netSettings)
		printMetricV2(netToProcess, netSettings)
		return
	case V1MetricType:
		printMetricV1(netToProcess, netSettings)
	case V2MetricType:
		printMetricV2(netToProcess, netSettings)
	default:
		println("Incorrect metric type.")
		return
	}
}

func countCharacteristicsAndGetMessage(net *net.PetriNet, metric string, netSettings settings.Settings) string {

}

func batchFlow(dirPath, metric string, netSettings settings.Settings, fn writer.OutputFunc) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	// Iterate over the files.
	for _, file := range files {
		// Check if the file is a directory.
		if !file.IsDir() {
			// Print the file name.
			fmt.Println(file.Name())
			standardFlow(dirPath+"/"+file.Name(), metric, netSettings)
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

func printMetricV1(net *net.PetriNet, settings settings.Settings) {
	metric := algorithm.CountMetricVersion1(net, settings)
	fmt.Printf("Metric 1 equals %f\n", metric)
}

func printMetricV2(net *net.PetriNet, settings settings.Settings) {
	metric := algorithm.CountMetric(net, settings)
	fmt.Printf("Metric 2 equals %f\n", metric)
}

func getCharacteristicV1Message(value float64) string {
	return fmt.Sprintf("Characteristic 1 equals %f\n", value)
}

func getCharacteristicV2Message(value float64) string {
	return fmt.Sprintf("Characteristic 2 equals %f\n", value)

}

func getOutputFunction(filePath string) writer.OutputFunc {
	if filePath == "" {
		return consoleOutput
	}
	return func(text string) {
		f, err := os.Create(filePath)
		if err != nil {
			// todo add error to message.
			panic("Can't write to file")
		}
		defer f.Close()

		data := []byte(text)

		_, err = f.Write(data)

		if err != nil {
			// todo add error message.
			panic("Error writing to file")
		}
	}
}

func consoleOutput(text string) {
	fmt.Printf(text)
}
