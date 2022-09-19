package sensor

import (
	"errors"
	"sensor/pkg/global"
	"sensor/pkg/load"
)

type memorySensor struct {
}

func NewSensorMemory() Metricer {
	return &memorySensor{}
}

func (m *memorySensor) GetMetrics() ([]*Measurement, error) {
	var еxecutor Executor
	еxecutor = NewЕxecutor()
	deviceID := load.GetDeviceID(global.DeviceName)
	sensorIdTotal, _ := load.GetSensorID(global.MemoryTotal)
	sensorIdAvailable, _ := load.GetSensorID(global.MemoryAvailableBytes)
	sensorIdUsed, _ := load.GetSensorID(global.MemoryUsedBytes)
	sensorIdUsedPercent, _ := load.GetSensorID(global.MemoryUsedPercent)
	var measurments []*Measurement
	total, available, used, usedPercent, err := еxecutor.GetMemory()
	if err != nil {
		return measurments, errors.New("problem getting memory")
	}
	return []*Measurement{NewMeasurement(float64(total), sensorIdTotal, deviceID), NewMeasurement(float64(available), sensorIdAvailable, deviceID), NewMeasurement(float64(used), sensorIdUsed, deviceID), NewMeasurement(float64(usedPercent), sensorIdUsedPercent, deviceID)}, err
}
