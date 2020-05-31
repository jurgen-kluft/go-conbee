package groups

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var (
	testUsername       string
	testHostname       string
	testGroups         *Groups
	transitionTime     uint16
	sleepSeconds       int
	sleepMilliSeconds  int
	testLightNumbers   []int
	redState           State
	blueState          State
	whiteState         State
	offState           State
	onState            State
	virginAmericaState State
)

func init() {
	testUsername = os.Getenv("DECONZ_CONBEE_APIKEY")
	testHostname = os.Getenv("DECONZ_CONBEE_HOST")
	testGroups = New(testHostname, testUsername)
	transitionTime = uint16(4)
	sleepSeconds = 4
	sleepMilliSeconds = 100
	testLightNumbers = []int{1, 2, 3, 4}

	virginAmericaState = State{}
	virginAmericaState.SetOn(true)
	virginAmericaState.SetCT(230, 223)

	redState = State{}
	redState.SetOn(true)
	redState.SetCT(13, 253)
	redState.SetXY(0.6736, 0.3221)

	blueState = State{}
	blueState.SetOn(true)
	blueState.SetCT(254, 500)
	blueState.SetXY(0.1754, 0.0556)

	whiteState = State{}
	whiteState.SetOn(true)
	whiteState.SetCT(203, 232)
	whiteState.SetXY(0.3151, 0.3252)

	offState = State{}
	offState.SetOn(false)

	onState = State{}
	offState.SetOn(true)
}

func TestCreateGroup(t *testing.T) {
	group := Group{Name: "Office", Lights: []string{"1", "2"}}
	_, err := testGroups.CreateGroup(group)
	if err != nil {
		t.Fail()
	}
}

func TestSetGroup(t *testing.T) {
	group := Group{Name: "Office", Lights: []string{"1", "2"}}
	_, err := testGroups.SetGroupAttrs(1, group)
	if err != nil {
		t.Fail()
	}
	group = Group{Name: "Bedroom", Lights: []string{"3", "4"}}
	_, err = testGroups.SetGroupAttrs(2, group)
	if err != nil {
		t.Fail()
	}
	group = Group{Name: "Living Room", Lights: []string{"5", "6"}}
	_, err = testGroups.SetGroupAttrs(3, group)
	if err != nil {
		t.Fail()
	}
	group = Group{Name: "Upstairs", Lights: []string{"1", "2", "3", "4", "5", "6", "8"}}
	_, err = testGroups.SetGroupAttrs(4, group)
	if err != nil {
		t.Fail()
	}
}

func TestGetAllGroups(t *testing.T) {
	groups, err := testGroups.GetAllGroups()
	if err != nil {
		t.Fail()
	}
	for _, g := range groups {
		t.Log(g.String())
	}
}

func TestGetGroup(t *testing.T) {
	for _, groupID := range testLightNumbers {
		g, err := testGroups.GetGroupAttrs(groupID)
		if err != nil {
			fmt.Println(err)
			t.Fail()
		}
		fmt.Println(g.String())
	}
}

func test_set_group_state(t *testing.T, state State) {
	groups, err := testGroups.GetAllGroups()
	if err != nil {
		t.Fail()
	}
	for _, group := range groups {
		// TODO: need to test response.
		_, err := testGroups.SetGroupState(group.ID, state)
		if err != nil {
			t.Fail()
		}
		time.Sleep(time.Millisecond * time.Duration(sleepMilliSeconds))
	}
}

func TestSetGroupState(t *testing.T) {
	groupsBackup, err := testGroups.GetAllGroups()
	if err != nil {
		t.Fail()
	}
	test_set_group_state(t, offState)
	test_set_group_state(t, virginAmericaState)
	test_set_group_state(t, offState)
	test_set_group_state(t, onState)
	test_set_group_state(t, offState)
	test_set_group_state(t, blueState)
	test_set_group_state(t, offState)
	for _, group := range groupsBackup {
		if group.ID == 0 {
			continue
		}
		// TODO: need to test response.
		_, err := testGroups.SetGroupState(group.ID, group.Action)
		if err != nil {
			t.Fail()
		}
		time.Sleep(time.Millisecond * time.Duration(sleepMilliSeconds))
	}
}
