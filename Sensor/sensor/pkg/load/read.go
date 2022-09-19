package load

import (
	"fmt"
	"io/ioutil"
	"sensor/pkg/global"
	"sensor/pkg/logger"
	"time"

	"gopkg.in/yaml.v2"
)

type Devices struct {
	Device []Device `yaml:"devices"`
}
type measurement struct {
	Time     time.Time `yaml:"measuredAt"`
	Value    float64   `yaml:"value"`
	SensorId string    `yaml:"sensorId"`
	DeviceId string    `yaml:"deviceId"`
}

type sensor struct {
	Id           string        `yaml:"id"`
	Name         string        `yaml:"name"`
	Description  string        `yaml:"description"`
	Unit         string        `yaml:"unit"`
	SensorGroups []string      `yaml:"sensorGroups"`
	Measurements []measurement `yaml:"measurements"`
}
type Device struct {
	Id          string   `yaml:"id"`
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Sensors     []sensor `yaml:"sensors"`
}

func readFile() (Devices, error) {
	devices := Devices{}
	data, err := ioutil.ReadFile("model.yaml")
	if err != nil {
		errString := fmt.Errorf("еrror reading file %w", err)
		logger.Errorloggers.ErrorLogger.Println(errString)
		return devices, err
	}
	err = yaml.Unmarshal(data, &devices)
	if err != nil {
		errString := fmt.Errorf("еrror unmrasheling struct %w", err)
		logger.Errorloggers.ErrorLogger.Println(errString)
		return devices, err
	}
	return devices, err
}
func GetSensorID(sensorGroupName string) (string, string) {
	var devices Devices
	var sensorId string
	var unit string
	devices, err := readFile()
	if err != nil {
		errString := fmt.Errorf("еrror reading file %w", err)
		logger.Errorloggers.ErrorLogger.Println(errString)
	}
	_, found := global.SensorGroups[sensorGroupName]
	if found == false {
		errString := "there is no " + sensorGroupName + " in the yaml file"
		logger.Errorloggers.ErrorLogger.Println(errString)
		return errString, ""
	}

	for i := 0; i < len(devices.Device); i++ {
		for j := 0; j < len(devices.Device[i].Sensors); j++ {
			if devices.Device[i].Sensors[j].Name == sensorGroupName {
				sensorId = devices.Device[i].Sensors[j].Id
				unit = devices.Device[i].Sensors[j].Unit
			}
		}
	}
	return sensorId, unit
}
func GetDeviceID(deviceName string) string {
	var deviceId string
	var devices Devices
	devices, err := readFile()
	if err != nil {
		errString := fmt.Errorf("еrror reading file %w", err)
		logger.Errorloggers.ErrorLogger.Println(errString)
	}
	_, found := global.DeviceGroups[deviceName]
	if found == false {
		errString := "there is no " + deviceName + " in the yaml file"
		logger.Errorloggers.ErrorLogger.Println(errString)
		return errString
	}
	for i := 0; i < len(devices.Device); i++ {
		if devices.Device[i].Name == deviceName {
			deviceId = devices.Device[i].Id
		}
	}
	return deviceId
}
