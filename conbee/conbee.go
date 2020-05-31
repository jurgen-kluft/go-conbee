package conbee

// This package started off from the hue package implemented by 'heatxsink',
// so all of the credits should go to 'heatxsink'.
// URL: https://github.com/heatxsink/go-hue

// Note: deconz conbee REST API is VERY similar to the HUE REST API

type ApiResponse struct {
	Success map[string]interface{} `json:"success"`
	Error   *ApiResponseError      `json:"error"`
}

type ApiResponseError struct {
	Type        uint   `json:"type"`
	Address     string `json:"address"`
	Description string `json:"description"`
}
