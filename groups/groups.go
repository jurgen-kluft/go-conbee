package groups

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/jurgen-kluft/go-conbee/conbee"
	"github.com/jurgen-kluft/go-conbee/scenes"
)

var (
	createGroupURL   = "http://%s/api/%s/groups"
	getAllGroupsURL  = "http://%s/api/%s/groups"
	getGroupAttrsURL = "http://%s/api/%s/groups/%d"
	setGroupAttrsURL = "http://%s/api/%s/groups/%d"
	setGroupStateURL = "http://%s/api/%s/groups/%d/action"
	deleteGroupURL   = "http://%s/api/%s/groups/%d"
)

type Groups struct {
	Hostname string
	APIkey   string
}

type State struct {
	On     *bool     `json:"on"`
	Hue    uint16    `json:"hue,omitempty"`
	Effect string    `json:"effect,omitempty"`
	Bri    *uint8    `json:"bri,omitempty"`
	Sat    uint8     `json:"sat,omitempty"`
	CT     *uint16   `json:"ct,omitempty"`
	XY     []float32 `json:"xy,omitempty"`
}

type Group struct {
	ID               int            `json:"id,omitempty"`
	ETag             string         `json:"etag,omitempty"`
	Name             string         `json:"name"`
	Hidden           bool           `json:"hidden"`
	Action           State          `json:"action,omitempty"`
	Lights           []string       `json:"lights,omitempty"`
	LightSequence    []string       `json:"lightsequence,omitempty"`
	MultiDeviceIDs   []string       `json:"multideviceids,omitempty"`
	DeviceMembership []string       `json:"devicemembership,omitempty"`
	Scenes           []scenes.Scene `json:"scenes,omitempty"`
}

func New(hostname string, apikey string) *Groups {
	return &Groups{
		Hostname: hostname,
		APIkey:   apikey,
	}
}

func (g *Groups) GetAllGroups() ([]Group, error) {
	url := fmt.Sprintf(getAllGroupsURL, g.Hostname, g.APIkey)
	request, err := http.NewRequest("GET", url, nil)
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	groupsMap := map[string]Group{}
	json.Unmarshal(contents, &groupsMap)
	groups := make([]Group, 0, len(groupsMap))
	for groupID, group := range groupsMap {
		group.ID, err = strconv.Atoi(groupID)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	return groups, err
}

func (g *Groups) CreateGroup(group Group) ([]conbee.ApiResponse, error) {
	var apiResponse []conbee.ApiResponse
	url := fmt.Sprintf(createGroupURL, g.Hostname, g.APIkey)
	jsonData, err := json.Marshal(&group)
	if err != nil {
		return apiResponse, err
	}
	postBody := strings.NewReader(string(jsonData))
	request, err := http.NewRequest("POST", url, postBody)
	if err != nil {
		return apiResponse, err
	}
	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return apiResponse, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return apiResponse, err
	}
	err = json.Unmarshal(contents, &apiResponse)
	if err != nil {
		return apiResponse, err
	}
	return apiResponse, err
}

func (g *Groups) GetGroupAttrs(groupID int) (Group, error) {
	var gg Group
	url := fmt.Sprintf(getGroupAttrsURL, g.Hostname, g.APIkey, groupID)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return gg, err
	}
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return gg, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return gg, err
	}
	gg.ID = groupID
	err = json.Unmarshal(contents, &gg)
	if err != nil {
		return gg, err
	}
	return gg, err
}

func (g *Groups) SetGroupAttrs(groupID int, group Group) ([]conbee.ApiResponse, error) {
	var apiResponse []conbee.ApiResponse
	url := fmt.Sprintf(setGroupAttrsURL, g.Hostname, g.APIkey, groupID)
	jsonData, err := json.Marshal(&group)
	if err != nil {
		return apiResponse, err
	}
	body := strings.NewReader(string(jsonData))
	request, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return apiResponse, err
	}
	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return apiResponse, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return apiResponse, err
	}
	json.Unmarshal(contents, &apiResponse)
	return apiResponse, err
}

func (g *Groups) SetGroupState(groupID int, state State) ([]conbee.ApiResponse, error) {
	var apiResponse []conbee.ApiResponse
	url := fmt.Sprintf(setGroupStateURL, g.Hostname, g.APIkey, groupID)
	jsonData, err := json.Marshal(&state)
	if err != nil {
		return apiResponse, err
	}
	body := strings.NewReader(string(jsonData))
	request, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return apiResponse, err
	}
	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return apiResponse, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return apiResponse, err
	}
	err = json.Unmarshal(contents, &apiResponse)
	if err != nil {
		return apiResponse, err
	}
	return apiResponse, err
}

// DeleteGroup deletes a group from the deconz hub
func (g *Groups) DeleteGroup(groupID int) ([]conbee.ApiResponse, error) {
	var apiResponse []conbee.ApiResponse
	url := fmt.Sprintf(deleteGroupURL, g.Hostname, g.APIkey)
	request, err := http.NewRequest("DELETE", url, nil)
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(contents, &apiResponse)
	if err != nil {
		return apiResponse, err
	}
	return apiResponse, err
}

func (s *State) String() string {
	var buffer bytes.Buffer
	if s.On != nil {
		buffer.WriteString(fmt.Sprintf("On:              %t\n", s.On))
	}
	buffer.WriteString(fmt.Sprintf("Hue:             %d\n", s.Hue))
	buffer.WriteString(fmt.Sprintf("Effect:          %s\n", s.Effect))
	if s.Bri != nil {
		buffer.WriteString(fmt.Sprintf("Bri:             %d\n", s.Bri))
	}
	buffer.WriteString(fmt.Sprintf("Sat:             %d\n", s.Sat))
	if s.CT != nil {
		buffer.WriteString(fmt.Sprintf("CT:              %d\n", s.CT))
	}
	if len(s.XY) > 0 {
		buffer.WriteString(fmt.Sprintf("XY:              %g, %g\n", s.XY[0], s.XY[1]))
	}
	return buffer.String()
}

func (g *Group) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("ID:              %d\n", g.ID))
	buffer.WriteString(fmt.Sprintf("Name:            %s\n", g.Name))
	buffer.WriteString("Action:\n")
	buffer.WriteString(g.Action.String())
	buffer.WriteString("Lights:\n")
	for _, lightID := range g.Lights {
		buffer.WriteString(fmt.Sprintf("\t%s\n", lightID))
	}
	return buffer.String()
}
