package service

import (
	"errors"
	"server/logger"
	"server/models"
	"server/repository"
	"strconv"
)

type sensor struct {
	Repo repository.Repository
}

var BaseExecutorSensor Service = &sensor{}

func (d *sensor) ConfigureRepo(repository repository.Repository) {
	d.Repo = repository
}

func (s *sensor) GetAll() (interface{}, error) {
	logger.GetInstance().InfoLogger.Println("Getting all sensors - service...")
	result, err := s.Repo.GetAll()
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	logger.GetInstance().InfoLogger.Println("get all sensors-service")
	return result, nil
}
func (s *sensor) GetById(id string) (interface{}, error) {
	logger.GetInstance().InfoLogger.Println("Getting sensor by id - service...")
	result, err := s.Repo.GetById(id)
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	logger.GetInstance().InfoLogger.Println("Got sensor by id - service")
	return result, nil
}
func (s *sensor) Create(device models.Device, sensor models.Sensor) (interface{}, error) {
	logger.GetInstance().InfoLogger.Println("Creating sensor - service...")
	if sensor.Id == "" {
		err := errors.New("id cannot be empty")
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	if sensor.DeviceId == "" {
		err := errors.New("deviceid cannot be empty")
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	stringID := sensor.Id
	ID, err := strconv.ParseInt(stringID, 10, 64)
	if err != nil {
		err := errors.New("problem converting to int")
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	if ID < 1 {
		err := errors.New("sensor id cannot be lower than one")
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	stringDeviceID := sensor.DeviceId
	deviceID, err := strconv.ParseInt(stringDeviceID, 10, 64)
	if err != nil {
		err := errors.New("problem converting to int")
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	if deviceID < 1 {
		err := errors.New("device id cannot be lower than one")
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	result, err := s.Repo.Create(device, sensor)
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	logger.GetInstance().InfoLogger.Println("Created sensor-service")
	return result, nil
}
func (s *sensor) Delete(id string) error {
	logger.GetInstance().InfoLogger.Println("Deleting sensor-service...")
	err := s.Repo.Delete(id)
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return err
	}
	logger.GetInstance().InfoLogger.Println("Deleted sensor-service")
	return nil
}
func (s *sensor) Update(id string, device models.Device, sensor models.Sensor) (interface{}, error) {
	logger.GetInstance().InfoLogger.Println("Updating sensor-service")
	result, err := s.Repo.Update(id, device, sensor)
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	logger.GetInstance().InfoLogger.Println("Updated sensor-service")
	return result, nil
}
