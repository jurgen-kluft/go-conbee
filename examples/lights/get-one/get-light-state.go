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
	lightID    = 0
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: get-light-state -host=[string] -key=[string]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	on := new(bool)
	*on = true
	fmt.Println(os.Getenv("DECONZ_CONBEE_HOST"))
	flag.StringVar(&conbeeHost, "host", os.Getenv("DECONZ_CONBEE_HOST"), "conbee host addr")
	flag.StringVar(&conbeeKey, "key", os.Getenv("DECONZ_CONBEE_APIKEY"), "conbee api key")
	flag.IntVar(&lightID, "id", 31, "light ID")
	flag.Parse()
	flag.Usage = usage

	if conbeeKey != "" {
		ll := lights.New(conbeeHost, conbeeKey)
		allLights, err := ll.GetAllLights()
		if err != nil {
			fmt.Println("lights.GetAllLights() ERROR: ", err)
			os.Exit(1)
		}

		fmt.Println()
		if len(allLights) > 0 {
			fmt.Println("------")
			light, err := ll.GetLightState(lightID)
			if err != nil {
				fmt.Println("lights.GetLightState() ERROR: ", err)
				os.Exit(1)
			}
			fmt.Printf("Light:\n%s\n", light.String())
		}
	} else {
		usage()
	}
}
