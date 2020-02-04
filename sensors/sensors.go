package sensors

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/jurgen-kluft/go-conbee/conbee"
)

var (
	getAllSensorsURL = "http://%s/api/%s/sensors"
	getSensorURL     = "http://%s/api/%s/sensors/%d"
)

type Sensors struct {
	Hostname string
	APIkey   string
}

type Sensor struct {
	ID               int    `json:"id,omitempty"`
	Name             string `json:"name"`
	State            State  `json:"state,omitempty"`
	Config           Config `json:"config,omitempty"`
	Type             string `json:"type,omitempty"`
	ModelID          string `json:"modelid,omitempty"`
	SWVersion        string `json:"swversion,omitempty"`
	ManufacturerName string `json:"manufacturername,omitempty"`
	UniqueID         string `json:"uniqueid,omitempty"`
}

type Config struct {
	On            bool   `json:"on"`
	Long          string `json:"long,omitempty"`
	Lat           string `json:"lat,omitempty"`
	SunriseOffset int16  `json:"sunriseoffset,omitempty"`
	SunsetOffset  int16  `json:"sunsetoffset,omitempty"`
}

type State struct {
	Presence    bool   `json:"presence,omitempty"`
	LastUpdated string `json:"lastupdated,omitempty"`
	ButtonEvent int16  `json:"buttonevent,omitempty"`
	Status      int16  `json:"status,omitempty"`
	Temperature int16  `json:"temperature,omitempty"`
	LightLevel  int16  `json:"lightlevel,omitempty"`
	Dark        bool   `json:"dark,omitempty"`
	Daylight    bool   `json:"daylight,omitempty"`
}

func New(hostname string, apikey string) *Sensors {
	return &Sensors{
		Hostname: hostname,
		APIkey:   apikey,
	}
}

func (l *Sensors) GetSensor(sensorID int) (Sensor, error) {
	var ll Sensor
	url := fmt.Sprintf(getSensorURL, l.Hostname, l.APIkey, sensorID)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ll, err
	}
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return ll, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return ll, err
	}
	err = json.Unmarshal(contents, &ll)
	if err != nil {
		return ll, err
	}
	ll.ID = sensorID
	return ll, err
}

func (l *Sensors) UpdateSensor(sensorID int, sensorName string) ([]conbee.ApiResponse, error) {
	url := fmt.Sprintf(getSensorURL, l.Hostname, l.APIkey, sensorID)
	data := fmt.Sprintf("{\"name\": \"%s\"}", sensorName)
	postbody := strings.NewReader(data)
	request, err := http.NewRequest("PUT", url, postbody)
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
	return apiResponse, err
}

func (l *Sensors) GetAllSensors() ([]Sensor, error) {
	url := fmt.Sprintf(getAllSensorsURL, l.Hostname, l.APIkey)
	request, err := http.NewRequest("GET", url, nil)
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
	//fmt.Println(string(contents))

	sensorsMap := map[string]Sensor{}
	err = json.Unmarshal(contents, &sensorsMap)
	if err != nil {
		return nil, err
	}
	sensors := make([]Sensor, 0, len(sensorsMap))
	for sensorID, sensor := range sensorsMap {
		sensor.ID, _ = strconv.Atoi(sensorID)
		sensors = append(sensors, sensor)
	}

	sort.Slice(sensors, func(i, j int) bool { return sensors[i].ID < sensors[j].ID })

	return sensors, err
}

func (l *Sensor) String() string {
	return l.StringWithIndentation("")
}

func (l *Sensor) StringWithIndentation(indentation string) string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%sID:              %d\n", indentation, l.ID))
	buffer.WriteString(fmt.Sprintf("%sUUID:            %s\n", indentation, l.UniqueID))
	buffer.WriteString(fmt.Sprintf("%sName:            %s\n", indentation, l.Name))
	buffer.WriteString(fmt.Sprintf("%sType:            %s\n", indentation, l.Type))
	buffer.WriteString(fmt.Sprintf("%sModelId:         %s\n", indentation, l.ModelID))
	buffer.WriteString(fmt.Sprintf("%sSwVersion:       %s\n", indentation, l.SWVersion))
	buffer.WriteString(fmt.Sprintf("%sState:\n", indentation))
	buffer.WriteString(l.State.StringWithIndentation(indentation + indentation))
	return buffer.String()
}

func (s *State) String() string {
	return s.StringWithIndentation("")
}

func (s *State) StringWithIndentation(indentation string) string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%sButtonEvent:     %d\n", indentation, s.ButtonEvent))
	buffer.WriteString(fmt.Sprintf("%sDaylight:        %t\n", indentation, s.Daylight))
	buffer.WriteString(fmt.Sprintf("%sDark:            %t\n", indentation, s.Dark))
	buffer.WriteString(fmt.Sprintf("%sLastUpdated:     %s\n", indentation, s.LastUpdated))
	buffer.WriteString(fmt.Sprintf("%sLightLevel:      %d\n", indentation, s.LightLevel))
	buffer.WriteString(fmt.Sprintf("%sPresence:        %t\n", indentation, s.Presence))
	buffer.WriteString(fmt.Sprintf("%sStatus:          %d\n", indentation, s.Status))
	buffer.WriteString(fmt.Sprintf("%sTemperature:     %d\n", indentation, s.Temperature))
	return buffer.String()
}
