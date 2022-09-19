package service

import (
	"errors"
	"server/logger"
	"server/models"
	"server/repository"
	"strconv"
)

type MeasurementService interface {
	GetAll() (interface{}, error)
	Create(measurement models.Measurement) (interface{}, error)
	SensorAvarageValue(sensorID string, deviceID string, startTime string, endTime string) (float64, error)
	SensorsCorrelationCoefficient(sensorID1 string, deviceID1 string, sensorID2 string, deviceID2 string, startTime string, endTime string) (float64, error)
	ConfigureRepoMeasurement(repository repository.MeasurementRepo)
}
type measurement struct {
	Repo repository.MeasurementRepo
}

var BaseExecutorMeasurement MeasurementService = &measurement{}

func NewServiceMeasurement() MeasurementService {
	return &measurement{repository.NewRepoMeasurement()}
}
func (m *measurement) ConfigureRepoMeasurement(repository repository.MeasurementRepo) {
	m.Repo = repository
}
func (m *measurement) GetAll() (interface{}, error) {
	logger.GetInstance().InfoLogger.Println("Getting all measurements-service...")
	result, err := m.Repo.GetAll()
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	logger.GetInstance().InfoLogger.Println("Got all measurements-service")
	return result, nil
}
func (m *measurement) Create(measurement models.Measurement) (interface{}, error) {
	logger.GetInstance().InfoLogger.Println("Creating new measurements-service...")
	if measurement.SensorID == "" {
		err := errors.New("sensorID cannot be empty")
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	if measurement.DeviceID == "" {
		err := errors.New("deviceid cannot be empty")
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	stringSensorID := measurement.SensorID
	sensorID, err := strconv.ParseInt(stringSensorID, 10, 64)
	if err != nil {
		err := errors.New("problem converting to int")
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	if sensorID < 1 {
		err := errors.New("id cannot be lower than one")
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	stringDeviceID := measurement.SensorID
	deviceID, err := strconv.ParseInt(stringDeviceID, 10, 64)
	if err != nil {
		err := errors.New("problem converting to int")
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	if deviceID < 1 {
		err := errors.New("id cannot be lower than one")
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	result := m.Repo.Create(measurement)
	logger.GetInstance().InfoLogger.Println("Created new measurements-service")
	return result, nil
}

func (m *measurement) SensorAvarageValue(sensorID string, deviceID string, startTime string, endTime string) (float64, error) {
	logger.GetInstance().InfoLogger.Println("Average value measurement-service about to be getted...")
	result, err := m.Repo.SensorAvarageValue(sensorID, deviceID, startTime, endTime)
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return 0, err
	}
	logger.GetInstance().InfoLogger.Println("Average value measurement-service getted")
	return result, nil
}

func (m *measurement) SensorsCorrelationCoefficient(sensorID1 string, deviceID1 string, sensorID2 string, deviceID2 string, startTime string, endTime string) (float64, error) {
	logger.GetInstance().InfoLogger.Println("Correlation coefficient measurement-service about to be getted...")
	result, err := m.Repo.SensorsCorrelationCoefficient(sensorID1, deviceID1, sensorID2, deviceID2, startTime, endTime)
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return 0, err
	}
	logger.GetInstance().InfoLogger.Println("Correlation coefficient measurement-service getted")
	return result, nil
}
