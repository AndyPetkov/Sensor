package repository

import (
	"context"
	"errors"
	"fmt"
	"math"
	"server/global"
	"server/logger"
	"server/models"
	"strconv"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	influxapi "github.com/influxdata/influxdb-client-go/v2/api"
)

type Measurements []models.Measurement
type MeasurementRepo interface {
	GetAll() (interface{}, error)
	Create(measurement models.Measurement) interface{}
	SensorAvarageValue(sensorID string, deviceID string, startTime string, endTime string) (float64, error)
	SensorsCorrelationCoefficient(sensorID1 string, deviceID1 string, sensorID2 string, deviceID2 string, startTime string, endTime string) (float64, error)
}
type Measurement struct {
}

var BaseExecutorMeasurement MeasurementRepo = &Measurement{}

func NewRepoMeasurement() MeasurementRepo {
	return BaseExecutorMeasurement
}

func (m *Measurement) GetAll() (interface{}, error) {
	logger.GetInstance().InfoLogger.Println("Geting all measurements-repo...")
	var measurements Measurements
	query := `from(bucket: "Bucket")
	|> range(start:2021-11-30)
	|> filter(fn: (r) => r["_measurement"] == "stat")`

	result, err := getQuery(query)
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return measurements, err
	}
	for result.Next() {
		perMeasurment := models.Measurement{}
		perMeasurment.DeviceID = fmt.Sprintf("%v", result.Record().ValueByKey("deviceID"))
		perMeasurment.SensorID = fmt.Sprintf("%v", result.Record().ValueByKey("sensorID"))
		value := fmt.Sprintf("%v", result.Record().ValueByKey("_value"))
		perMeasurment.Value, err = strconv.ParseFloat(value, 64)
		if err != nil {
			logger.GetInstance().ErrorLogger.Println(err)
			return measurements, err
		}
		perMeasurment.Time = result.Record().Time()
		measurements = append(measurements, perMeasurment)
	}
	if result.Err() != nil {
		logger.GetInstance().ErrorLogger.Println(result.Err())
		return measurements, result.Err()
	}

	logger.GetInstance().InfoLogger.Println("Got all measurements-repo")
	return measurements, err
}
func (m *Measurement) Create(measurement models.Measurement) interface{} {
	logger.GetInstance().InfoLogger.Println("Creating new measurements-repo")
	client := influxdb2.NewClient(global.URL, global.Token)
	writeAPI := client.WriteAPI(global.Org, global.Bucket)

	p := influxdb2.NewPointWithMeasurement("stat").
		AddField("measure", measurement.Value).
		AddTag("deviceID", measurement.DeviceID).
		AddTag("sensorID", measurement.SensorID).
		SetTime(time.Now())
	writeAPI.WritePoint(p)
	writeAPI.Flush()
	logger.GetInstance().InfoLogger.Println("Created new measurements-repo")
	return measurement
}
func (m *Measurement) SensorAvarageValue(sensorID string, deviceID string, startTime string, endTime string) (float64, error) {
	logger.GetInstance().InfoLogger.Println("Getting average value measurement-repo...")
	query := `from(bucket: "Bucket")
	|> range(start:` + startTime + `, stop:` + endTime + `)
	|> filter(fn: (r) => r["_measurement"] == "stat")`
	result, err := getQuery(query)
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return 0, err
	}
	var sumValue float64 = 0.0
	var counter int = 0
	var existSensorIdDeviceId bool = false
	for result.Next() {
		DeviceID := fmt.Sprintf("%v", result.Record().ValueByKey("deviceID"))
		SensorID := fmt.Sprintf("%v", result.Record().ValueByKey("sensorID"))
		if deviceID == DeviceID && sensorID == SensorID {
			valueString := fmt.Sprintf("%v", result.Record().ValueByKey("_value"))
			valueFloat, err := strconv.ParseFloat(valueString, 64)
			if err != nil {
				logger.GetInstance().ErrorLogger.Println(err)
				return 0, err
			}
			sumValue = valueFloat + sumValue
			counter++
			existSensorIdDeviceId = true
		}
	}
	if existSensorIdDeviceId == false {
		err = errors.New("there is no such device with " + deviceID + " id and sensor " + sensorID + " id in the database")
		logger.GetInstance().ErrorLogger.Println(err)
		return 0, err
	}
	value := sumValue / float64(counter)
	if result.Err() != nil {
		logger.GetInstance().ErrorLogger.Println(result.Err())
		return 0, result.Err()
	}
	logger.GetInstance().InfoLogger.Println("Got average value measurement-repo")
	return value, err
}

func (m *Measurement) SensorsCorrelationCoefficient(sensorID1 string, deviceID1 string, sensorID2 string, deviceID2 string, startTime string, endTime string) (float64, error) {
	logger.GetInstance().InfoLogger.Println("Getting correlation coefficient measurement-repo...")
	query := `from(bucket: "Bucket")
	|> range(start:` + startTime + `, stop:` + endTime + `)
	|> filter(fn: (r) => r["_measurement"] == "stat")`
	result, err := getQuery(query)
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return 0, err
	}
	var measurementValues []float64
	var measurementValuesSecond []float64
	var existSensorId1DeviceId1 bool = false
	var existSensorId2DeviceId2 bool = false
	for result.Next() {
		DeviceID := fmt.Sprintf("%v", result.Record().ValueByKey("deviceID"))
		SensorID := fmt.Sprintf("%v", result.Record().ValueByKey("sensorID"))
		if deviceID1 == DeviceID && sensorID1 == SensorID {
			valueString := fmt.Sprintf("%v", result.Record().ValueByKey("_value"))
			valueFloat, err := strconv.ParseFloat(valueString, 64)
			if err != nil {
				logger.GetInstance().ErrorLogger.Println(err)
				return 0, err
			}
			measurementValues = append(measurementValues, valueFloat)
			existSensorId1DeviceId1 = true
		}
		if deviceID2 == DeviceID && sensorID2 == SensorID {
			valueString := fmt.Sprintf("%v", result.Record().ValueByKey("_value"))
			valueFloat, err := strconv.ParseFloat(valueString, 64)
			if err != nil {
				logger.GetInstance().ErrorLogger.Println(err)
				return 0, err
			}
			measurementValuesSecond = append(measurementValuesSecond, valueFloat)
			existSensorId2DeviceId2 = true
		}
	}
	if existSensorId1DeviceId1 == false {
		err = errors.New("there is no such device with " + deviceID1 + " id and sensor " + sensorID1 + " id in the database")
		logger.GetInstance().ErrorLogger.Println(err)
		return 0, err
	}
	if existSensorId2DeviceId2 == false {
		err = errors.New("there is no such device with " + deviceID2 + " id and sensor " + sensorID2 + " id in the database")
		logger.GetInstance().ErrorLogger.Println(err)
		return 0, err
	}
	if len(measurementValues) != len(measurementValuesSecond) {
		err = errors.New("lenth of slices cant be different")
		logger.GetInstance().ErrorLogger.Println(err)
		return 0, errors.New("lenth of slices cant be different")
	}
	value := calculatePearson(measurementValues, measurementValuesSecond)
	if value == 0 {
		err = errors.New("pearson value is 0, because all values are the same")
		logger.GetInstance().ErrorLogger.Println(err)
		return 0, err
	}

	logger.GetInstance().InfoLogger.Println("Got correlation coefficient measurement-repo")
	return value, err
}
func calculatePearson(valueSliceX []float64, valueSliceY []float64) float64 {
	logger.GetInstance().InfoLogger.Println("Calculating pearsons value measurement-repo...")
	var sumX float64
	var sumY float64
	for i := 0; i < len(valueSliceX); i++ {
		sumX = valueSliceX[i] + sumX
		sumY = valueSliceY[i] + sumY
	}
	meanX := sumX / float64(len(valueSliceX))
	meanY := sumY / float64(len(valueSliceY))

	var xSliceMinus []float64
	var ySliceMinus []float64
	for i := 0; i < len(valueSliceX); i++ {
		valueX := valueSliceX[i] - meanX
		valueY := valueSliceY[i] - meanY
		xSliceMinus = append(xSliceMinus, valueX)
		ySliceMinus = append(ySliceMinus, valueY)
	}
	var multiplySlice []float64
	for i := 0; i < len(valueSliceX); i++ {
		value := xSliceMinus[i] * ySliceMinus[i]
		multiplySlice = append(multiplySlice, value)
	}
	var sumMultiplied float64
	for i := 0; i < len(multiplySlice); i++ {
		sumMultiplied = multiplySlice[i] + sumMultiplied
	}
	for i := 0; i < len(valueSliceX); i++ {
		xSliceMinus[i] = xSliceMinus[i] * xSliceMinus[i]
		ySliceMinus[i] = ySliceMinus[i] * ySliceMinus[i]
	}
	var sumXPOw float64
	var sumYPOw float64
	for i := 0; i < len(valueSliceX); i++ {
		sumXPOw = xSliceMinus[i] + sumXPOw
		sumYPOw = ySliceMinus[i] + sumYPOw
	}
	if sumMultiplied == 0 {
		return 0
	}
	finalValue := sumMultiplied / math.Sqrt((sumXPOw * sumYPOw))
	logger.GetInstance().InfoLogger.Println("Calculated pearsons value measurement-repo")
	return finalValue
}

func getQuery(query string) (*influxapi.QueryTableResult, error) {
	client := influxdb2.NewClient(global.URL, global.Token)
	queryAPI := client.QueryAPI(global.Org)
	client.Close()
	return queryAPI.Query(context.Background(), query)
}
