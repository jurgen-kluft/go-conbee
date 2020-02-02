package schedules

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

var (
	getAllSchedulesURL  = "http://%s/api/%s/schedules"
	getScheduleAttrsURL = "http://%s/api/%s/schedule/%d"
)

type Schedules struct {
	Hostname string
	APIkey   string
}

type Schedule struct {
	ID          int
	Name        string  `json:"name"`
	ETag        string  `json:"etag"`
	Description string  `json:"description,omitempty"`
	Cmd         Command `json:"command"`
	Status      string  `json:"status, omitempty"`
	AutoDelete  bool    `json:"autodelete,omitempty"`
	Time        string  `json:"time"`
}

type Command struct {
	Address string `json:"address,omitempty"`
	Method  string `json:"method,omitempty"`
	Body    string `json:"body,omitempty"`
}

func New(hostname string, apikey string) *Schedules {
	return &Schedules{
		Hostname: hostname,
		APIkey:   apikey,
	}
}

func (l *Schedules) GetScheduleAttrs(scheduleID int) (Schedule, error) {
	var ll Schedule
	url := fmt.Sprintf(getScheduleAttrsURL, l.Hostname, l.APIkey, scheduleID)
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
	ll.ID = scheduleID
	return ll, err
}

func (l *Schedules) GetAllSensors() ([]Schedule, error) {
	url := fmt.Sprintf(getAllSchedulesURL, l.Hostname, l.APIkey)
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
	objMap := map[string]Schedule{}
	err = json.Unmarshal(contents, &objMap)
	if err != nil {
		return nil, err
	}
	objs := make([]Schedule, 0, len(objMap))
	for objID, obj := range objMap {
		obj.ID, _ = strconv.Atoi(objID)
		objs = append(objs, obj)
	}
	return objs, err
}

func (l *Schedule) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("etag:        %d\n", l.ETag))
	buffer.WriteString(fmt.Sprintf("name:        %s\n", l.Name))
	buffer.WriteString(fmt.Sprintf("description: %s\n", l.Description))
	buffer.WriteString(fmt.Sprintf("status:      %s\n", l.Status))
	buffer.WriteString(fmt.Sprintf("time:        %s\n", l.Time))
	buffer.WriteString(fmt.Sprintf("autodelete:  %v\n", l.AutoDelete))
	buffer.WriteString(fmt.Sprint("command:\n"))
	buffer.WriteString(l.Cmd.String())
	return buffer.String()
}

func (c *Command) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("address: %t\n", c.Address))
	buffer.WriteString(fmt.Sprintf("method:  %d\n", c.Method))
	buffer.WriteString(fmt.Sprintf("body:    %d\n", c.Body))
	return buffer.String()
}
