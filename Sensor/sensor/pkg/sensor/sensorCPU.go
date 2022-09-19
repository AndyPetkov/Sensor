package sensor

import (
	"sensor/pkg/global"
	"sensor/pkg/load"
)

type cpuSensor struct {
}

func NewSensorCPU() Metricer {
	return &cpuSensor{}
}

func (c *cpuSensor) GetMetrics() ([]*Measurement, error) {
	var еxecutor Executor
	еxecutor = NewЕxecutor()
	deviceID := load.GetDeviceID(global.DeviceName)
	sensorIdPercent, _ := load.GetSensorID(global.CpuUsagePercent)
	sensorIdNumOfCores, _ := load.GetSensorID(global.CpuCoresCount)
	sensorIdMhz, _ := load.GetSensorID(global.CpuFrequency)
	percent, numOfCores, mhz, err := еxecutor.GetCPU()
	return []*Measurement{NewMeasurement(percent, sensorIdPercent, deviceID), NewMeasurement(float64(numOfCores), sensorIdNumOfCores, deviceID), NewMeasurement(mhz, sensorIdMhz, deviceID)}, err
}
