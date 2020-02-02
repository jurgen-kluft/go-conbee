package configuration

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jurgen-kluft/go-conbee/conbee"
	"github.com/jurgen-kluft/go-conbee/groups"
	"github.com/jurgen-kluft/go-conbee/lights"
	"github.com/jurgen-kluft/go-conbee/rules"
	"github.com/jurgen-kluft/go-conbee/schedules"
	"github.com/jurgen-kluft/go-conbee/sensors"
)

var (
	aquireAPIkeyURL  = "http://%s/api"
	configurationURL = "http://%s/api/%s/config"
	deleteAPIkeyURL  = "http://%s/api/%s/config/whitelist/%s"
	fullStateURL     = "http://%s/api/%s"
)

type User struct {
	UserName   string `json:"username,omitempty"`
	DeviceType string `json:"devicetype"`
}

type SWUpdate struct {
	UpdateState int    `json:updatestate`
	URL         string `json:url`
	Text        string `json:text`
	Notify      bool   `json:notify`
}

type Whitelist struct {
	LastUseDate string `json:"last use date"`
	CreateDate  string `json:"create date"`
	Name        string `json:name`
}

type Configuration struct {
	Hostname            string
	APIVersion          string               `json:"apiversion,omitempty"`
	DHCP                bool                 `json:dhcp`
	Gateway             string               `json:gateway`
	IPAddress           string               `json:ipaddress`
	LinkButton          bool                 `json:linkbutton`
	LocalTime           string               `json:localtime`
	Mac                 string               `json:mac`
	Name                string               `json:"name,omitempty"`
	NetMask             string               `json:netmask`
	NetworkOpenDuration int                  `json:networkopenduration`
	PanID               int                  `json:panid`
	PortalServices      bool                 `json:portalservices`
	ProxyAddress        string               `json:proxyaddress`
	ProxyPort           int                  `json:proxyport`
	SWUpdate            SWUpdate             `json:swupdate`
	SWVersion           string               `json:swversion`
	TimeFormat          string               `json:timeformat`
	TimeZone            string               `json:timezone`
	UTC                 string               `json:utc`
	UUID                string               `json:uuid`
	Whitelist           map[string]Whitelist `json:whitelist`
	ZigbeeChannel       int                  `json:zigbeechannel`
}

// FullState is populated when calling GetFullState()
type FullState struct {
	Config    Configuration        `json:config`
	Groups    []groups.Group       `json:groups`
	Lights    []lights.Light       `json:lights`
	Sensors   []sensors.Sensor     `json:sensors`
	Schedules []schedules.Schedule `json:schedules`
	Rules     []rules.Rule         `json:rules`
}

func New(hostname string) *Configuration {
	return &Configuration{
		Hostname: hostname,
	}
}

func (c *Configuration) AquireAPIKey(userName string, deviceType string) ([]conbee.ApiResponse, error) {
	user := User{
		UserName:   userName,
		DeviceType: deviceType,
	}
	jsonData, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf(aquireAPIkeyURL, c.Hostname)
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
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
	var apiResponse []conbee.ApiResponse
	err = json.Unmarshal(contents, &apiResponse)
	if err != nil {
		return nil, err
	}
	if apiResponse[0].Error != nil {
		return nil, errors.New(apiResponse[0].Error.Description)
	}
	return apiResponse, err
}

func (c *Configuration) DeleteAPIKey(apikey string, apikeyToBeDeleted string) ([]conbee.ApiResponse, error) {
	url := fmt.Sprintf(deleteAPIkeyURL, c.Hostname, apikey, apikeyToBeDeleted)
	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
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
	var apiResponse []conbee.ApiResponse
	err = json.Unmarshal(contents, &apiResponse)
	if err != nil {
		return nil, err
	}
	return apiResponse, err
}

func (c *Configuration) GetFullState(username string) (FullState, error) {
	var fullstate FullState
	url := fmt.Sprintf(fullStateURL, c.Hostname, username)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fullstate, err
	}
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return fullstate, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fullstate, err
	}
	err = json.Unmarshal(contents, &fullstate)
	if err != nil {
		return fullstate, err
	}
	return fullstate, err
}

func (c *Configuration) GetConfiguration(username string) (Configuration, error) {
	var configuration Configuration
	url := fmt.Sprintf(configurationURL, c.Hostname, username)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return configuration, err
	}
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return configuration, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return configuration, err
	}
	if response.StatusCode != 200 {
		fmt.Println(string(contents))
	}
	err = json.Unmarshal(contents, &configuration)
	if err != nil {
		return configuration, err
	}
	return configuration, err
}

func (c *Configuration) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Name:         %s\n", c.Name))
	buffer.WriteString(fmt.Sprintf("UTC:          %s\n", c.UTC))
	buffer.WriteString(fmt.Sprintf("SWVersion:    %s\n", c.SWVersion))
	buffer.WriteString(fmt.Sprintf("ProxyAddress: %s\n", c.ProxyAddress))
	buffer.WriteString(fmt.Sprintf("ProxyPort:    %d\n", c.ProxyPort))
	buffer.WriteString(fmt.Sprintf("Mac:          %s\n", c.Mac))
	buffer.WriteString(fmt.Sprintf("LinkButton:   %t\n", c.LinkButton))
	buffer.WriteString(fmt.Sprintf("IPAddress:    %s\n", c.IPAddress))
	buffer.WriteString(fmt.Sprintf("NetMask:      %s\n", c.NetMask))
	buffer.WriteString(fmt.Sprintf("Gateway:      %s\n", c.Gateway))
	buffer.WriteString(fmt.Sprintf("DHCP:         %t\n", c.DHCP))
	buffer.WriteString(fmt.Sprintf("SWUpdate:\n"))
	buffer.WriteString(fmt.Sprintf("\tUpdateState: %d\n", c.SWUpdate.UpdateState))
	buffer.WriteString(fmt.Sprintf("\tURL:         %s\n", c.SWUpdate.URL))
	buffer.WriteString(fmt.Sprintf("\tText:        %s\n", c.SWUpdate.Text))
	buffer.WriteString(fmt.Sprintf("\tNotify:      %t\n", c.SWUpdate.Notify))
	buffer.WriteString(fmt.Sprintf("Whitelist:\n"))
	for key := range c.Whitelist {
		buffer.WriteString(fmt.Sprintf("\tKey: %s\n", key))
		buffer.WriteString(fmt.Sprintf("\t\tLastUseDate: %s\n", c.Whitelist[key].LastUseDate))
		buffer.WriteString(fmt.Sprintf("\t\tCreateDate:  %s\n", c.Whitelist[key].CreateDate))
		buffer.WriteString(fmt.Sprintf("\t\tName:        %s\n", c.Whitelist[key].Name))
	}
	return buffer.String()
}
