package client

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sensor/pkg/sensor"
)

func PostData(web_hook_url string, measurement *sensor.Measurement) error {
	buf, err := json.Marshal(measurement)
	if err != nil {
		return err
	}
	_, err = http.Post(web_hook_url, "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil
}
