package repository

import (
	"errors"
	"server/database"
	"server/logger"
	"server/models"
)

type Sensors []models.Sensor

type Sensor struct {
}

var BaseExecutorSensor Repository = &Sensor{}

func (s *Sensor) GetAll() (interface{}, error) {
	logger.GetInstance().InfoLogger.Println("Geting all sensors-repo...")
	sqlSensor := "SELECT * FROM sensor"
	rows, err := database.Database.Query(sqlSensor)
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	var sensors Sensors
	for rows.Next() {
		perSensor := models.Sensor{}
		err = rows.Scan(&perSensor.Id, &perSensor.DeviceId, &perSensor.Name, &perSensor.Description, &perSensor.Unit, &perSensor.Sensorgroup)
		if err != nil {
			logger.GetInstance().ErrorLogger.Println(err)
			return nil, err
		}
		sensors = append(sensors, perSensor)
	}
	logger.GetInstance().InfoLogger.Println("Got all sensors-repo")
	return sensors, err
}
func (s *Sensor) GetById(id string) (interface{}, error) {
	logger.GetInstance().InfoLogger.Println("Geting sensor by id - repo...")
	err := checkRowExistbyId("SELECT id FROM sensor WHERE id=$1", id)
	if err != nil {
		err := errors.New("sensor with id " + id + " does not exist")
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	dataSensor, err := database.Database.Query("SELECT * FROM sensor WHERE id=$1", id)
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	sensor := models.Sensor{}
	for dataSensor.Next() {
		err = dataSensor.Scan(&sensor.Id, &sensor.DeviceId, &sensor.Name, &sensor.Description, &sensor.Unit, &sensor.Sensorgroup)
		if err != nil {
			logger.GetInstance().ErrorLogger.Println(err)
			return nil, err
		}
	}
	logger.GetInstance().InfoLogger.Println("Got sensor by id-repo")
	return sensor, nil
}
func (s *Sensor) Create(device models.Device, sensor models.Sensor) (interface{}, error) {
	logger.GetInstance().InfoLogger.Println("Creating sensor-repo...")
	err := checkRowExistbyId("SELECT id FROM sensor WHERE id=$1", sensor.Id)
	if err == nil {
		err := errors.New("sensor with id " + sensor.Id + " already exist")
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	sqlStatement := `INSERT INTO sensor (id,deviceid,name,description,unit,sensorgroup) VALUES ($1, $2,$3, $4,$5,$6)`
	_, err = database.Database.Exec(sqlStatement, sensor.Id, sensor.DeviceId, sensor.Name, sensor.Description, sensor.Unit, sensor.Sensorgroup)
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	logger.GetInstance().InfoLogger.Println("Created sensor-repo")
	return sensor, err
}
func (s *Sensor) Delete(id string) error {
	logger.GetInstance().InfoLogger.Println("Deleting sensor-repo...")
	err := checkRowExistbyId("SELECT id FROM sensor WHERE id=$1", id)
	if err != nil {
		err := errors.New("sensor with id " + id + " does not exist")
		logger.GetInstance().ErrorLogger.Println(err)
		return err
	}

	_, err = database.Database.Exec("DELETE FROM sensor where id = $1", id)
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return err
	}
	logger.GetInstance().InfoLogger.Println("Deleted sensor-repo")
	return nil
}
func (s *Sensor) Update(id string, device models.Device, sensor models.Sensor) (interface{}, error) {
	logger.GetInstance().InfoLogger.Println("Updating sensor-repo")
	err := checkRowExistbyId("SELECT id FROM sensor WHERE id=$1", id)
	if err != nil {
		err := errors.New("sensor with id " + id + " does not exist")
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	sqlStatement := `
	                             UPDATE sensor
	                              SET deviceid=$2, name = $3, description = $4, unit  = $5, sensorgroup  = $6
	                             WHERE id = $1;`
	_, err = database.Database.Exec(sqlStatement, id, sensor.DeviceId, sensor.Name, sensor.Description, sensor.Unit, sensor.Sensorgroup)
	sensor.Id = id
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	logger.GetInstance().InfoLogger.Println("Updated sensor-repo")
	return sensor, nil
}

func GetSensorSliceByDeviceId(deviceId string) ([]models.SensorWithoutDeviceID, error) {
	logger.GetInstance().InfoLogger.Println("Geting sensor slice by device id sensor-repo...")
	dataSensor, err := database.Database.Query("SELECT * FROM sensor WHERE deviceid=$1", deviceId)
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	var sensors SensorsWithoutDeviceId
	var dummmy string
	for dataSensor.Next() {
		perSensor := models.SensorWithoutDeviceID{}
		err = dataSensor.Scan(&perSensor.Id, &dummmy, &perSensor.Name, &perSensor.Description, &perSensor.Unit, &perSensor.Sensorgroup)
		if err != nil {
			logger.GetInstance().ErrorLogger.Println(err)
			return nil, err
		}
		sensors = append(sensors, perSensor)
	}
	logger.GetInstance().InfoLogger.Println("Got sensor slice by device id sensor-repo")
	return sensors, err
}
