package main

import (
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

	switch *metric {
	case "all":
		return
	case "v1":
	case "v2":
	default:
		println("Incorrect metric type.")
		return
	}

	fmt.Printf("metric %s \nnet: %s \nsettings: %s\n", *metric, *netPath, *settingsPath)
}
