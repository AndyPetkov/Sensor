package repository

import (
	"errors"
	"server/database"
	"server/logger"
	"server/models"
)

type Devices []models.Device
type SensorsWithoutDeviceId []models.SensorWithoutDeviceID
type Repository interface {
	GetAll() (interface{}, error)
	GetById(id string) (interface{}, error)
	Create(device models.Device, sensor models.Sensor) (interface{}, error)
	Delete(id string) error
	Update(id string, device models.Device, sensor models.Sensor) (interface{}, error)
}
type Device struct {
}

var BaseExecutorDevice Repository = &Device{}

func NewRepo(sensorType string) Repository {

	if sensorType == "device" {
		return BaseExecutorDevice
	}
	return BaseExecutorSensor
}

func (d *Device) GetAll() (interface{}, error) {
	var devices Devices
	data, err := database.Database.Query("SELECT * FROM device")
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	for data.Next() {
		perDevice := models.Device{}
		err = data.Scan(&perDevice.Id, &perDevice.Name, &perDevice.Description)
		if err != nil {
			logger.GetInstance().ErrorLogger.Println(err)
			return nil, err
		}
		perDevice.Sensors, err = GetSensorSliceByDeviceId(perDevice.Id)
		if err != nil {
			logger.GetInstance().ErrorLogger.Println(err)
			return nil, err
		}
		devices = append(devices, perDevice)
	}
	logger.GetInstance().InfoLogger.Println("get all devices-repo")
	return devices, err
}
func (d *Device) GetById(id string) (interface{}, error) {
	err := checkRowExistbyId("SELECT id FROM device WHERE id=$1", id)
	if err != nil {
		err := errors.New("device with id " + id + " does not exist")
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	data, err := database.Database.Query("SELECT * FROM device WHERE id=$1", id)
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	device := models.Device{}
	for data.Next() {
		err = data.Scan(&device.Id, &device.Name, &device.Description)
		if err != nil {
			logger.GetInstance().ErrorLogger.Println(err)
			return nil, err
		}
	}
	device.Sensors, err = GetSensorSliceByDeviceId(id)
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	return device, nil
}
func (d *Device) Create(device models.Device, sensor models.Sensor) (interface{}, error) {
	logger.GetInstance().InfoLogger.Println("Creating new device-repo...")
	err := checkRowExistbyId("SELECT id FROM device WHERE id=$1", device.Id)
	if err == nil {
		err := errors.New("device with id " + device.Id + " already exist")
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	sqlStatement := `INSERT INTO device (id,name,description) VALUES ($1, $2,$3)`
	_, err = database.Database.Exec(sqlStatement, device.Id, device.Name, device.Description)
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	logger.GetInstance().InfoLogger.Println("Created new device-repo")
	return device, err
}
func (d *Device) Delete(id string) error {
	logger.GetInstance().InfoLogger.Println("Deleting device-repo...")
	err := checkRowExistbyId("SELECT id FROM device WHERE id=$1", id)
	if err != nil {
		err := errors.New("device with id " + id + " does not exist")
		logger.GetInstance().ErrorLogger.Println(err)
		return err
	}
	err = checkRowExistbyId("SELECT deviceid FROM sensor WHERE deviceid=$1", id)
	if err == nil {
		err := errors.New("there are sensors in this device, delete them first then delete the device")
		logger.GetInstance().ErrorLogger.Println(err)
		return err
	}
	_, err = database.Database.Exec("DELETE FROM device where id = $1", id)
	if err != nil {
		err := errors.New("device with id " + id + " does not exist")
		logger.GetInstance().ErrorLogger.Println(err)
		return err
	}
	logger.GetInstance().InfoLogger.Println("Deleted device-repo")
	return nil
}
func (d *Device) Update(id string, device models.Device, sensor models.Sensor) (interface{}, error) {
	logger.GetInstance().InfoLogger.Println("Updating device-repo...")
	err := checkRowExistbyId("SELECT id FROM device WHERE id=$1", id)
	if err != nil {
		err := errors.New("device with id " + id + " does not exist")
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	sqlStatement := `
                             UPDATE device
                              SET name = $2, description = $3
                             WHERE id = $1;`
	_, err = database.Database.Exec(sqlStatement, id, device.Name, device.Description)
	device.Id = id
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	logger.GetInstance().InfoLogger.Println("Updated device-repo")
	return device, nil
}
