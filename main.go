package main

import (
	"dns-tracker/capture"
	"dns-tracker/config"
	"dns-tracker/writer"
	"fmt"
	"log"
)

func main() {
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalln("Config load error:", err)
	}

	w := writer.NewJSONWriter(cfg.LogDir)

	fmt.Println("listening dns packets:", cfg.InterfaceName)
	if err = capture.ListenDNS(cfg.InterfaceName, w); err != nil {
		log.Fatalln("listening error:", err)
	}
}
