package controller

import (
	"encoding/json"
	"net/http"
	"server/logger"
	"server/models"
)

func decodeDevice(r *http.Request) (models.Device, models.Sensor, error) {
	var device models.Device
	var sensor models.Sensor
	err := json.NewDecoder(r.Body).Decode(&device)
	return device, sensor, err
}
func decodeSensor(r *http.Request) (models.Device, models.Sensor, error) {
	var device models.Device
	var sensor models.Sensor
	err := json.NewDecoder(r.Body).Decode(&sensor)
	return device, sensor, err
}

func writeResponse(w http.ResponseWriter, err error, result interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		errorMessage := err.Error()
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	json.NewEncoder(w).Encode(result)
	return
}
