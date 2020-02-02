package configuration

import (
	"os"
	"testing"
)

var (
	testUsername       string
	testHostname       string
	testDeleteUsername string
	testConfiguration  *Configuration
)

func init() {
	testUsername = os.Getenv("DECONZ_CONBEE_APIKEY")
	testHostname = os.Getenv("DECONZ_CONBEE_HOST")
	testDeleteUsername = ""
	testConfiguration = New(testHostname)
}

func TestAquireKey(t *testing.T) {
	testUsername := "deadc0ffee"
	testDeviceType := "unittest"
	apiResponse, err := testConfiguration.AquireAPIKey(testUsername, testDeviceType)
	if err != nil {
		t.Log("TestAquireKey Error: ", err)
		t.Fail()
	} else {
		t.Log(apiResponse[0].Success["username"].(string))
	}
}

func TestGetFullState(t *testing.T) {
	_, err := testConfiguration.GetFullState(testUsername)
	if err != nil {
		t.Log("TestGetFullState Error: ", err)
		t.Fail()
	}
}

func TestGetConfiguration(t *testing.T) {
	_, err := testConfiguration.GetConfiguration(testUsername)
	if err != nil {
		t.Log("TestGetConfiguration Error: ", err)
		t.Fail()
	}
}

func TestDeleteKey(t *testing.T) {
	_, err := testConfiguration.DeleteAPIKey(testUsername, testDeleteUsername)
	if err != nil {
		t.Log("TestDeleteKey Error: ", err)
		t.Fail()
	}
}
