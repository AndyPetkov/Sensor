package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"server/mocks"
	"server/models"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("controller-measurement", func() {
	var (
		actualMeasurementTest      models.Measurement
		measurementTestEmpty       models.Measurement
		errTest                    error
		testMeasurement            MeasurementController
		mockTest                   *mocks.MeasurementService
		actualAverageValueTest     map[string]float64
		actualAverageValueTestZero map[string]float64
		averageValueTestFloat      float64
		actualPearsonValueTest     map[string]float64
		actualPearsonValueTestZero map[string]float64
		pearsonValueTestFloat      float64
		expectedMeasurements       models.Measurement
		expectedValueTest          map[string]float64
	)
	BeforeEach(func() {
		actualMeasurementTest = models.Measurement{Time: time.Time{}, Value: 55.4, SensorID: "1", DeviceID: "1"}
		measurementTestEmpty = models.Measurement{Time: time.Time{}, Value: 0, SensorID: "", DeviceID: ""}
		errTest = errors.New("")
		mockTest = &mocks.MeasurementService{}
		testMeasurement = NewControllerMeasurement()
		testMeasurement.ConfigureServiceMeasurement(mockTest)
		actualAverageValueTest = map[string]float64{
			"average": 48.9,
		}
		averageValueTestFloat = 48.9
		actualPearsonValueTest = map[string]float64{
			"pearson value": 0.785,
		}
		pearsonValueTestFloat = 0.785
		expectedMeasurements = models.Measurement{}
		expectedValueTest = nil
	})

	Describe("GetAll", func() {

		It("should return all measurements and nil when all passed parameters are valid", func() {
			mockTest.On("GetAll").Return(actualMeasurementTest, nil)
			response := GetResponse("GET", "/measurements", nil, http.HandlerFunc(testMeasurement.GetAll))
			json.NewDecoder(io.Reader(response.Body)).Decode(&expectedMeasurements)
			Expect(response.Code).To(Equal(http.StatusOK))
			Expect(expectedMeasurements).To(Equal(actualMeasurementTest))
		},
		)
		It("should return nil and error when wrong parameters are passed", func() {
			mockTest.On("GetAll").Return(nil, errTest)
			response := GetResponse("GET", "/measurements", nil, http.HandlerFunc(testMeasurement.GetAll))
			json.NewDecoder(io.Reader(response.Body)).Decode(&expectedMeasurements)
			Expect(response.Code).To(Equal(http.StatusBadRequest))
			Expect(expectedMeasurements).To(Equal(measurementTestEmpty))
		},
		)
	})

	Describe("Create", func() {

		It("should create measurements and return nil when all passed parameters are valid", func() {
			mockTest.On("Create", actualMeasurementTest).Return(actualMeasurementTest, nil)
			payloadBuf := new(bytes.Buffer)
			json.NewEncoder(payloadBuf).Encode(actualMeasurementTest)
			response := GetResponse("POST", "/measurements", payloadBuf, http.HandlerFunc(testMeasurement.Create))
			json.NewDecoder(io.Reader(response.Body)).Decode(&expectedMeasurements)
			Expect(response.Code).To(Equal(http.StatusOK))
			Expect(expectedMeasurements).To(Equal(actualMeasurementTest))
		},
		)
		It("should return nil and error when wrong parameters are passed", func() {
			mockTest.On("Create", actualMeasurementTest).Return(nil, errTest)
			payloadBuf := new(bytes.Buffer)
			json.NewEncoder(payloadBuf).Encode(actualMeasurementTest)
			response := GetResponse("POST", "/measurements", payloadBuf, http.HandlerFunc(testMeasurement.Create))
			json.NewDecoder(io.Reader(response.Body)).Decode(&expectedMeasurements)
			Expect(response.Code).To(Equal(http.StatusBadRequest))
			Expect(expectedMeasurements).To(Equal(measurementTestEmpty))
		},
		)
	})
	Describe("SensorAvarageValue", func() {

		It("should return average value of a sensor by ID and device ID for a given period  when all passed parameters are valid", func() {
			mockTest.On("SensorAvarageValue", "1", "1", "15:24", "15:25").Return(averageValueTestFloat, nil)
			response := GetResponse("GET", "/sensorAvarageValue?sensorID=1&deviceID=1&startTime=15:24&endTime=15:25", nil, http.HandlerFunc(testMeasurement.SensorAvarageValue))
			json.NewDecoder(io.Reader(response.Body)).Decode(&expectedValueTest)
			Expect(response.Code).To(Equal(http.StatusOK))
			Expect(expectedValueTest).To(Equal(actualAverageValueTest))
		},
		)
		It("should return 0 and error when wrong parameters are passed", func() {
			mockTest.On("SensorAvarageValue", "1", "1", "15:24", "15:25").Return(0.0, errTest)
			response := GetResponse("GET", "/sensorAvarageValue?sensorID=1&deviceID=1&startTime=15:24&endTime=15:25", nil, http.HandlerFunc(testMeasurement.SensorAvarageValue))
			json.NewDecoder(io.Reader(response.Body)).Decode(&expectedValueTest)
			Expect(response.Code).To(Equal(http.StatusBadRequest))
			Expect(expectedValueTest).To(Equal(actualAverageValueTestZero))
		},
		)
	})
	Describe("SensorsCorrelationCoefficient", func() {

		It("should return Correlation Coefficient value between two sensors with ID and device ID for a given period when all passed parameters are valid", func() {
			mockTest.On("SensorsCorrelationCoefficient", "1", "1", "2", "1", "15:24", "15:25").Return(pearsonValueTestFloat, nil)
			response := GetResponse("GET", "/sensorsCorrelationCoefficient?sensorID1=1&deviceID1=1&sensorID2=2&deviceID2=1&startTime=15:24&endTime=15:25", nil, http.HandlerFunc(testMeasurement.SensorsCorrelationCoefficient))
			json.NewDecoder(io.Reader(response.Body)).Decode(&expectedValueTest)
			Expect(response.Code).To(Equal(http.StatusOK))
			Expect(expectedValueTest).To(Equal(actualPearsonValueTest))
		},
		)
		It("should return 0 and error when wrong parameters are passed", func() {
			mockTest.On("SensorsCorrelationCoefficient", "1", "1", "2", "1", "15:24", "15:25").Return(0.0, errTest)
			response := GetResponse("GET", "/sensorsCorrelationCoefficient?sensorID1=1&deviceID1=1&sensorID2=2&deviceID2=1&startTime=15:24&endTime=15:25", nil, http.HandlerFunc(testMeasurement.SensorsCorrelationCoefficient))
			json.NewDecoder(io.Reader(response.Body)).Decode(&expectedValueTest)
			Expect(response.Code).To(Equal(http.StatusBadRequest))
			Expect(expectedValueTest).To(Equal(actualPearsonValueTestZero))
		},
		)
	})
})
