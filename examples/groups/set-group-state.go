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
	groupID    = 0
	groupOn    = ""
	groupBri   = 0
	groupCT    = 0
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: get-group-state -host=[string] -key=[string]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func init() {
	on := new(bool)
	*on = true
	flag.StringVar(&conbeeHost, "host", os.Getenv("DECONZ_CONBEE_HOST"), "conbee host addr")
	flag.StringVar(&conbeeKey, "key", os.Getenv("DECONZ_CONBEE_APIKEY"), "conbee api key")
	flag.IntVar(&groupID, "id", 18, "group ID")
	flag.StringVar(&groupOn, "on", "", "group On/Off")
	flag.IntVar(&groupBri, "bri", 0, "group Brightness")
	flag.IntVar(&groupCT, "ct", 0, "group Color Temperature")
	flag.Parse()
	flag.Usage = usage
}

func main() {
	if conbeeKey != "" {
		fmt.Println("Set Group State")

		gg := groups.New(conbeeHost, conbeeKey)
		//g, err := gg.GetGroupAttrs(groupID)
		//if err != nil {
		//	fmt.Println("groups.GetGroupAttrs() ERROR: ", err)
		//	os.Exit(1)
		//}
		fmt.Println()
		{
			g := groups.State{}
			if groupOn == "on" {
				g.SetOn(true)
			} else if groupOn == "off" {
				g.SetOn(false)
			}
			if groupBri != 0 || groupCT != 0 {
				g.SetCT(groupBri, groupCT)
			}
			//g.Lights = []string{"31", "32", "33", "34", "35", "36"}

			fmt.Println("------")
			fmt.Println(g.String())

			apiresponses, err := gg.SetGroupState(groupID, g)
			if err != nil {
				fmt.Println("groups.SetGroupState() ERROR: ", err)
				os.Exit(1)
			}
			for _, ar := range apiresponses {
				fmt.Println(ar)
			}
		}
	} else {
		usage()
	}
}
