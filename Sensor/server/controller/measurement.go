package controller

import (
	"encoding/json"
	"net/http"
	"server/global"
	"server/logger"
	"server/models"
	"server/service"
)

type Measurements []models.Measurement
type measurementHandler struct {
	Service service.MeasurementService
}
type MeasurementController interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	SensorAvarageValue(w http.ResponseWriter, r *http.Request)
	SensorsCorrelationCoefficient(w http.ResponseWriter, r *http.Request)
	ConfigureServiceMeasurement(service service.MeasurementService)
}

var BaseExecutorMeasurement MeasurementController = &measurementHandler{}

func NewControllerMeasurement() MeasurementController {
	return &measurementHandler{service.NewServiceMeasurement()}
}
func (m *measurementHandler) ConfigureServiceMeasurement(service service.MeasurementService) {
	m.Service = service
}
func (m *measurementHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	logger.GetInstance().InfoLogger.Println("Getting all measurements...")
	result, err := m.Service.GetAll()
	if err != nil {
		writeResponse(w, err, result, http.StatusBadRequest)
		return
	}
	writeResponse(w, err, result, http.StatusOK)
	logger.GetInstance().InfoLogger.Println("All measurements are getted")
}

func (m *measurementHandler) Create(w http.ResponseWriter, r *http.Request) {
	logger.GetInstance().InfoLogger.Println("Creating measurement...")
	var measurement models.Measurement
	err := json.NewDecoder(r.Body).Decode(&measurement)
	if err != nil {
		writeResponse(w, err, nil, http.StatusBadRequest)
		return
	}
	result, err := m.Service.Create(measurement)
	if err != nil {
		writeResponse(w, err, result, http.StatusBadRequest)
		return
	}
	writeResponse(w, err, result, http.StatusOK)
	logger.GetInstance().InfoLogger.Println("New measurement has been created")
}

func (m *measurementHandler) SensorAvarageValue(w http.ResponseWriter, r *http.Request) {
	logger.GetInstance().InfoLogger.Println("Getting sensor average value...")
	sensorID := r.URL.Query().Get(global.SensorID)
	deviceID := r.URL.Query().Get(global.DeviceID)
	startTime := r.URL.Query().Get(global.StartTime)
	endTime := r.URL.Query().Get(global.EndTime)
	result, err := m.Service.SensorAvarageValue(sensorID, deviceID, startTime, endTime)
	if err != nil {
		writeResponse(w, err, result, http.StatusBadRequest)
		return
	}
	var averageValue map[string]float64 = map[string]float64{
		"average": result,
	}
	writeResponse(w, err, averageValue, http.StatusOK)
	logger.GetInstance().InfoLogger.Println("Sensor Avarage Value getted")
}

func (m *measurementHandler) SensorsCorrelationCoefficient(w http.ResponseWriter, r *http.Request) {
	logger.GetInstance().InfoLogger.Println("Getting sensor correlation coefficient...")
	sensorID1 := r.URL.Query().Get(global.SensorID1)
	deviceID1 := r.URL.Query().Get(global.DeviceID1)
	sensorID2 := r.URL.Query().Get(global.SensorID2)
	deviceID2 := r.URL.Query().Get(global.DeviceID2)
	startTime := r.URL.Query().Get(global.StartTime)
	endTime := r.URL.Query().Get(global.EndTime)
	result, err := m.Service.SensorsCorrelationCoefficient(sensorID1, deviceID1, sensorID2, deviceID2, startTime, endTime)
	if err != nil {
		writeResponse(w, err, result, http.StatusBadRequest)
		return
	}
	var pearsonsValue map[string]float64 = map[string]float64{
		"pearson value": result,
	}
	writeResponse(w, err, pearsonsValue, http.StatusOK)
	logger.GetInstance().InfoLogger.Println("Sensors Correlation Coefficient getted")
}
