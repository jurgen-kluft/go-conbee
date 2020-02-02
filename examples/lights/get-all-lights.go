package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jurgen-kluft/go-conbee/lights"
)

var (
	conbeeHost = "10.0.0.18"
	conbeeKey  = "0A498B9909"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: get-light-state -host=[string] -key=[string]\n")
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
		ll := lights.New(conbeeHost, conbeeKey)
		allLights, err := ll.GetAllLights()
		if err != nil {
			fmt.Println("lights.GetAllLights() ERROR: ", err)
			os.Exit(1)
		}
		fmt.Println()
		fmt.Println("Lights")
		fmt.Println("------")
		for _, l := range allLights {
			fmt.Printf("ID: %d Name: %s\n", l.ID, l.Name)
		}
	} else {
		usage()
	}
}
