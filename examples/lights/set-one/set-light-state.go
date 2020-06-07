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
	lightOn    = ""
	lightBri   = 150
	lightCT    = 500
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: set-light-state -host=[string] -key=[string] -id=ID -on=true/false \n")
	flag.PrintDefaults()
	os.Exit(2)
}

func init() {
	on := new(bool)
	*on = true
	flag.StringVar(&conbeeHost, "host", os.Getenv("DECONZ_CONBEE_HOST"), "conbee host addr")
	flag.StringVar(&conbeeKey, "key", os.Getenv("DECONZ_CONBEE_APIKEY"), "conbee api key")
	flag.IntVar(&lightID, "id", 31, "light ID")
	flag.StringVar(&lightOn, "on", "", "light On/Off")
	flag.IntVar(&lightBri, "bri", 150, "light Brightness")
	flag.IntVar(&lightCT, "ct", 500, "light Color Temperature")
	flag.Parse()
	flag.Usage = usage
}

func main() {
	if conbeeKey != "" {
		ll := lights.New(conbeeHost, conbeeKey)
		//allLights, err := ll.GetAllLights()
		//if err != nil {
		//	fmt.Println("lights.GetAllLights() ERROR: ", err)
		//	os.Exit(1)
		//}
		//fmt.Println()
		//if len(allLights) > 0 {
		{
			fmt.Println("------")
			state := &lights.State{}
			if lightOn == "on" {
				state.SetOn(true)
			} else if lightOn == "off" {
				state.SetOn(false)
			}
			state.SetCT(lightBri, lightCT)
			_, err := ll.SetLightState(lightID, state)
			if err != nil {
				fmt.Println("lights.SetLightState() ERROR: ", err)
				os.Exit(1)
			}
			fmt.Printf("Light:\n%s\n", state.String())
		}
	} else {
		usage()
	}
}
