package sensor

import (
	"errors"
	"sensor/pkg/global"
	"sensor/pkg/logger"
)

func NewSensor(sensorGroup string) (Metricer, error) {
	var err error
	var flag bool = false
	var metricer Metricer
	if sensorGroup == global.CPU_USAGE {
		metricer = NewSensorCPU()
		flag = true
		logger.Infologgers.InfoLogger.Println(global.CPU_USAGE)
	}
	if sensorGroup == global.MEMORY_USAGE {
		metricer = NewSensorMemory()
		flag = true
		logger.Infologgers.InfoLogger.Println(global.MEMORY_USAGE)
	}
	if sensorGroup == global.CPU_TEMP {
		metricer = NewSensorTemp()
		flag = true
		logger.Infologgers.InfoLogger.Println(global.CPU_TEMP)
	}
	if flag == false {
		err = errors.New("there is no sensorGroup with the name " + sensorGroup)
		logger.Errorloggers.ErrorLogger.Println(err)
	}

	return metricer, err
}
