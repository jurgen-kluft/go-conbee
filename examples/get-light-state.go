package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/jurgen-kluft/go-conbee/groups"
	"github.com/jurgen-kluft/go-conbee/lights"
)

var (
	conbeeHost = "http://10.0.0.18"
	conbeeKey  = "0A498B9909"
	blinkState groups.State
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: get-light-state -host=[string] -key=[string]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func init() {
	on := new(bool)
	*on = true
	blinkState = groups.State{On: on}
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
		gg := groups.New(conbeeHost, conbeeKey)
		allGroups, err := gg.GetAllGroups()
		if err != nil {
			fmt.Println("groups.GetAllGroups() ERROR: ", err)
			os.Exit(1)
		}
		fmt.Println()
		fmt.Println("Groups")
		fmt.Println("------")
		for _, g := range allGroups {
			fmt.Printf("ID: %d Name: %s\n", g.ID, g.Name)
			for _, lll := range g.Lights {
				fmt.Println("\t", lll)
			}
			previousState := g.Action
			_, err := gg.SetGroupState(g.ID, blinkState)
			if err != nil {
				fmt.Println("groups.SetGroupState() ERROR: ", err)
				os.Exit(1)
			}
			time.Sleep(time.Second * time.Duration(10))
			_, err = gg.SetGroupState(g.ID, previousState)
			if err != nil {
				fmt.Println("groups.SetGroupState() ERROR: ", err)
				os.Exit(1)
			}
		}
	} else {
		usage()
	}
}
