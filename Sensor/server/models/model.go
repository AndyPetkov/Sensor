package models

import "time"

type Device struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Sensors     []SensorWithoutDeviceID
}

type Measurement struct {
	Time     time.Time `json:"measuredAt"`
	Value    float64   `json:"value:"`
	SensorID string    `json:"sensorId:"`
	DeviceID string    `json:"deviceId:"`
}

type Sensor struct {
	Id          string `json:"id"`
	DeviceId    string `json:"deviceid"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Unit        string `json:"unit"`
	Sensorgroup string `json:"sensorgroup"`
}
type SensorWithoutDeviceID struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Unit        string `json:"unit"`
	Sensorgroup string `json:"sensorgroup"`
}
