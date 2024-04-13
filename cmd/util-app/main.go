package main

import (
	"complexity/internal/reader/pipe"
	"complexity/pkg/algorithm"
	"complexity/pkg/net"
	"complexity/pkg/settings"
	"flag"
	"fmt"
)

func main() {
	metric := flag.String("metrics", "all", "metric version")
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

	netSettings, err := settings.ReadSettings(*settingsPath)
	if err != nil {
		fmt.Printf("Erorr: %s", err)
		return
	}
	netToProcess, err := pipe.ReadNet(*netPath, netSettings.SilentTransitions)
	if err != nil {
		fmt.Printf("Erorr: %s", err)
		return
	}
	switch *metric {
	case "all":
		printMetricV1(netToProcess, netSettings)
		printMetricV2(netToProcess, netSettings)
		return
	case "v1":
		printMetricV1(netToProcess, netSettings)
	case "v2":
		printMetricV2(netToProcess, netSettings)
	default:
		println("Incorrect metric type.")
		return
	}
}

func printMetricV1(net *net.PetriNet, settings *settings.Settings) {
	metric := algorithm.CountMetricVersion1(net, settings)
	fmt.Printf("Mettric 1 equals %f\n", metric)
}

func printMetricV2(net *net.PetriNet, settings *settings.Settings) {
	metric := algorithm.CountMetric(net, settings)
	fmt.Printf("Mettric 2 equals %f\n", metric)
}
