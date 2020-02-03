package scenes

import ()

type Scenes struct {
	Hostname string
	APIkey   string
}

type Scene struct {
	ID     string
	Name   string   `json:"name"`
	Lights []string `json:lights`
}

func New(hostname string, apikey string) *Scenes {
	return &Scenes{
		Hostname: hostname,
		APIkey:   apikey,
	}
}
