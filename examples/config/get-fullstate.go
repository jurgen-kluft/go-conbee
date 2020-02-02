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
	fmt.Fprintf(os.Stderr, "usage: get-fullstate -host=[string] -key=[string]\n")
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
		fmt.Println("TEST: configuration.GetFullState()")

		state, err := ll.GetFullState(conbeeKey)
		if err != nil {
			fmt.Println("ERROR: ", err)
			os.Exit(1)
		}
		fmt.Println("OK: ", state.Config.String())
		fmt.Println()

	} else {
		usage()
	}
}
