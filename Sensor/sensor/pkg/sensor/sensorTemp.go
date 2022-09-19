package sensor

import (
	"errors"
	"fmt"
	"sensor/pkg/global"
	"sensor/pkg/load"
	"sensor/pkg/logger"
	"time"
)

type Measurement struct {
	Time     time.Time `json:"measuredAt"  yaml:"measuredAt"`
	Value    float64   `json:"value:" yaml:"value"`
	SensorID string    `json:"sensorId:" yaml:"sensorId"`
	DeviceID string    `json:"deviceId:" yaml:"deviceId"`
}
type tempSensor struct {
}

func NewSensorTemp() Metricer {
	return &tempSensor{}
}

type Metricer interface {
	GetMetrics() ([]*Measurement, error)
}

func NewMeasurement(temperatureOfCPU float64, sensorId string, deviceId string) *Measurement {
	return &Measurement{
		Time:     time.Now(),
		Value:    temperatureOfCPU,
		SensorID: sensorId,
		DeviceID: deviceId,
	}
}
func (t *tempSensor) GetMetrics() ([]*Measurement, error) {
	var еxecutor Executor
	еxecutor = NewЕxecutor()
	var measurmants []*Measurement
	deviceID := load.GetDeviceID(global.DeviceName)
	sensorId, unit := load.GetSensorID(global.CpuTempCelsius)
	err := validate("C")
	if err != nil {
		errString := fmt.Errorf("error validating %w", err)
		logger.Errorloggers.ErrorLogger.Print(errString)
		return measurmants, err
	}
	temperatureOfCPU, err := еxecutor.GetTemp(unit)
	logger.Infologgers.InfoLogger.Println("Temp being getted")
	if err != nil {
		errString := fmt.Errorf("error geting temp %w", err)
		logger.Errorloggers.ErrorLogger.Print(errString)
		return measurmants, err
	}

	return []*Measurement{NewMeasurement(temperatureOfCPU, sensorId, deviceID)}, err

}

func validate(unit string) error {
	if unit != global.Celsius && unit != global.Fahrenheit {
		return errors.New("unit cannot be " + unit)
	}
	return nil
}
