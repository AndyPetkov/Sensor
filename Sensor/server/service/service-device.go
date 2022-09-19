package service

import (
	"errors"
	"server/global"
	"server/logger"
	"server/models"
	"server/repository"
	"strconv"
)

type Service interface {
	GetAll() (interface{}, error)
	GetById(id string) (interface{}, error)
	Create(device models.Device, sensor models.Sensor) (interface{}, error)
	Delete(id string) error
	Update(id string, device models.Device, sensor models.Sensor) (interface{}, error)
	ConfigureRepo(repository repository.Repository)
}
type Device struct {
	Repo repository.Repository
}

func NewService(sensorType string) Service {
	if sensorType == global.Device {
		return &Device{repository.NewRepo(global.Device)}

	}
	return &sensor{repository.NewRepo(global.Sensor)}
}

func (d *Device) ConfigureRepo(repository repository.Repository) {
	d.Repo = repository
}

func (d *Device) GetAll() (interface{}, error) {
	logger.GetInstance().InfoLogger.Println("Getting all devices-service...")
	result, err := d.Repo.GetAll()
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	logger.GetInstance().InfoLogger.Println("Got all devices-service")
	return result, nil
}
func (d *Device) GetById(id string) (interface{}, error) {
	logger.GetInstance().InfoLogger.Println("Geting device by id - service...")
	result, err := d.Repo.GetById(id)
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	logger.GetInstance().InfoLogger.Println("Got device by id - service")
	return result, nil
}
func (d *Device) Create(device models.Device, sensor models.Sensor) (interface{}, error) {
	logger.GetInstance().InfoLogger.Println("Creating device - service...")
	if device.Id == "" {
		err := errors.New("id cannot be empty")
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	stringID := device.Id
	ID, err := strconv.ParseInt(stringID, 10, 64)
	if err != nil {
		err := errors.New("problem converting to int")
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	if ID < 1 {
		err := errors.New("id cannot be lower than one")
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	result, err := d.Repo.Create(device, sensor)
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	logger.GetInstance().InfoLogger.Println("Created new device-service")
	return result, nil
}
func (d *Device) Delete(id string) error {
	logger.GetInstance().InfoLogger.Println("Deleting device by id - service...")
	err := d.Repo.Delete(id)
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return err
	}
	logger.GetInstance().InfoLogger.Println("Deleted device-service")
	return nil
}
func (d *Device) Update(id string, device models.Device, sensor models.Sensor) (interface{}, error) {
	logger.GetInstance().InfoLogger.Println("Updating device - service...")
	result, err := d.Repo.Update(id, device, sensor)
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	logger.GetInstance().InfoLogger.Println("Updated device-service")
	return result, nil
}
