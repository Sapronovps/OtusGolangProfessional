package main

import (
	"flag"
	"fmt"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", "/etc/calendar-sender/config.yaml", "Path to configuration file")
	flag.Parse()

	config := NewConfig(configFile)

	fmt.Println(config)
}
