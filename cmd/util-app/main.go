package main

import (
	"complexity/internal/reader/pipe"
	"complexity/pkg/algorithm"
	"complexity/pkg/net"
	"complexity/pkg/settings"
	"flag"
	"fmt"
)

const SimpleSettingsType = "simple"
const RegexpSettingsType = "regexp"
const AllMetricType = "all"
const V1MetricType = "v1"
const V2MetricType = "v2"

func main() {
	metric := flag.String("metrics", AllMetricType, "metric version")
	settingsType := flag.String("settings-type", SimpleSettingsType, "settings type (simple or regexp)")
	settingsPath := flag.String("settings", "", "net settings")
	netPath := flag.String("net", "", "net description")
	flag.Parse()

	if *settingsPath == "" {
		fmt.Println("Please provide path to net settings.")
		flag.Usage()
		return
	}

	if *netPath == "" {
		fmt.Println("Please provide path to net description.")
		flag.Usage()
		return
	}

	netSettings, err := getSettings(*settingsPath, *settingsType)
	if err != nil {
		fmt.Printf("Erorr: %s", err)
		return
	}
	netToProcess, err := pipe.ReadNet(*netPath, netSettings)
	if err != nil {
		fmt.Printf("Erorr: %s", err)
		return
	}
	switch *metric {
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
	fmt.Printf("Mettric 1 equals %f\n", metric)
}

func printMetricV2(net *net.PetriNet, settings settings.Settings) {
	metric := algorithm.CountMetric(net, settings)
	fmt.Printf("Mettric 2 equals %f\n", metric)
}
