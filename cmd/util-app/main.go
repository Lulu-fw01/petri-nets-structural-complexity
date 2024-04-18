package main

import (
	"complexity/internal/reader/pipe"
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
	SimpleSettingsType = "simple"
	RegexpSettingsType = "regexp"
	AllMetricType      = "all"
	V1MetricType       = "v1"
	V2MetricType       = "v2"
)

func main() {
	metric := flag.String(MetricTypeFlag, AllMetricType, "metric version")
	isBatchProcess := flag.Bool(BatchProcessFlag, false, "Process package of nets")
	settingsType := flag.String(SettingsTypeFlag, SimpleSettingsType, "settings type (simple or regexp)")
	settingsPath := flag.String(SettingsPathFlag, "", "net settings")
	netPath := flag.String(NetPathFlag, "", "net description")
	flag.Parse()

	validateSettingsPath(*settingsPath)
	validateNetPath(*netPath)

	netSettings, err := getSettings(*settingsPath, *settingsType)
	if err != nil {
		fmt.Printf("Erorr: %s", err)
		return
	}

	if *isBatchProcess {
		packageFlow(*netPath, *metric, netSettings)
	} else {
		standardFlow(*netPath, *metric, netSettings)
	}
}

func standardFlow(netPath, metric string, netSettings settings.Settings) {
	netToProcess, err := pipe.ReadNet(netPath, netSettings)
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

func packageFlow(dirPath, metric string, netSettings settings.Settings) {
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
