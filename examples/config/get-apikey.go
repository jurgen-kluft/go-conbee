package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jurgen-kluft/go-conbee/configuration"
)

var (
	conbeeHost = "10.0.0.18"
	conbeeKey  = "0A498B9909"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: get-apikey -host=[string] -key=[string]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func init() {
	on := new(bool)
	*on = true
	flag.StringVar(&conbeeHost, "host", os.Getenv("DECONZ_CONBEE_HOST"), "conbee host addr")
	flag.StringVar(&conbeeKey, "key", os.Getenv("DECONZ_CONBEE_APIKEY"), "conbee api key")
	flag.Parse()
	flag.Usage = usage
}

func main() {
	if conbeeKey != "" {
		ll := configuration.New(conbeeHost)

		fmt.Println("Configuration")
		fmt.Println("------")
		response, err := ll.AquireAPIKey("988112a4e198cc1211", "get-apikey")
		if err != nil {
			fmt.Println("configuration.AquireAPIKey() ERROR: ", err)
			os.Exit(1)
		}
		fmt.Println("configuration.AquireAPIKey() OK: ", response[0].Success["username"])
		fmt.Println()

	} else {
		usage()
	}
}
