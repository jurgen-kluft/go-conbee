package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jurgen-kluft/go-conbee/groups"
)

var (
	conbeeHost = "10.0.0.18"
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
	blinkState = groups.State{On: on, Alert: "lselect"}
	flag.StringVar(&conbeeHost, "host", os.Getenv("DECONZ_CONBEE_HOST"), "conbee host addr")
	flag.StringVar(&conbeeKey, "key", os.Getenv("DECONZ_CONBEE_APIKEY"), "conbee api key")
	flag.Parse()
	flag.Usage = usage
}

func main() {
	if conbeeKey != "" {
		fmt.Println("Groups")

		gg := groups.New(conbeeHost, conbeeKey)
		allGroups, err := gg.GetAllGroups()
		if err != nil {
			fmt.Println("groups.GetAllGroups() ERROR: ", err)
			os.Exit(1)
		}
		fmt.Println()
		for _, g := range allGroups {
			fmt.Println("------")
			fmt.Printf("Group:\n%s\n", g.String())
		}
	} else {
		usage()
	}
}
