package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"server/global"
	"server/mocks"
	"server/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("controller-sensor", func() {
	var (
		deviceTest            models.Device
		actualSensorTestEmpty models.Sensor
		errTest               error
		testSensor            Controller
		mockTest              *mocks.Service
		actualSensorTest      models.Sensor
		expectedSensor        models.Sensor
	)
	BeforeEach(func() {
		actualSensorTest = models.Sensor{Id: "17", DeviceId: "1", Name: "cpuTempCelsius", Description: "Measures CPU temp Celsius", Unit: "C", Sensorgroup: "CPU_USAGE"}
		actualSensorTestEmpty = models.Sensor{Id: "", DeviceId: "", Name: "", Description: "", Unit: "", Sensorgroup: ""}
		errTest = errors.New("")
		mockTest = &mocks.Service{}
		testSensor = NewController(global.Sensor)
		testSensor.ConfigureService(mockTest)
		expectedSensor = models.Sensor{}
	})

	Describe("GetAll", func() {

		It("should return all sensors and nil when all passed parameters are valid", func() {
			mockTest.On("GetAll").Return(actualSensorTest, nil)
			response := GetResponse("GET", "/sensors", nil, http.HandlerFunc(testSensor.GetAll))
			json.NewDecoder(io.Reader(response.Body)).Decode(&expectedSensor)
			Expect(response.Code).To(Equal(http.StatusOK))
			Expect(expectedSensor).To(Equal(actualSensorTest))
		},
		)
		It("should return nil and error when wrong parameters are passed", func() {
			mockTest.On("GetAll").Return(nil, errTest)
			response := GetResponse("GET", "/sensors", nil, http.HandlerFunc(testSensor.GetAll))
			json.NewDecoder(io.Reader(response.Body)).Decode(&expectedSensor)
			Expect(response.Code).To(Equal(http.StatusBadRequest))
			Expect(expectedSensor).To(Equal(actualSensorTestEmpty))
		},
		)
	})
	Describe("GetById", func() {

		It("should return sensor by ID and nil when all passed parameters are valid", func() {

			mockTest.On("GetById", "").Return(actualSensorTest, nil)
			response := GetResponse("GET", "/sensor/1", nil, http.HandlerFunc(testSensor.GetById))
			json.NewDecoder(io.Reader(response.Body)).Decode(&expectedSensor)
			Expect(response.Code).To(Equal(http.StatusOK))
			Expect(expectedSensor).To(Equal(actualSensorTest))
		},
		)
		It("should return nil and error when wrong parameters are passed", func() {
			mockTest.On("GetById", "").Return(nil, errTest)
			response := GetResponse("GET", "/sensor/2", nil, http.HandlerFunc(testSensor.GetById))
			json.NewDecoder(io.Reader(response.Body)).Decode(&expectedSensor)
			Expect(response.Code).To(Equal(http.StatusBadRequest))
			Expect(expectedSensor).To(Equal(actualSensorTestEmpty))
		},
		)
	})
	Describe("Create", func() {

		It("should create sensor and return nil when all passed parameters are valid", func() {
			mockTest.On("Create", deviceTest, actualSensorTest).Return(actualSensorTest, nil)
			payloadBuf := new(bytes.Buffer)
			json.NewEncoder(payloadBuf).Encode(actualSensorTest)
			response := GetResponse("POST", "/sensor", payloadBuf, http.HandlerFunc(testSensor.Create))
			json.NewDecoder(io.Reader(response.Body)).Decode(&expectedSensor)
			Expect(response.Code).To(Equal(http.StatusOK))
			Expect(expectedSensor).To(Equal(actualSensorTest))
		},
		)
		It("should return nil and error when wrong parameters are passed", func() {
			mockTest.On("Create", deviceTest, actualSensorTest).Return(nil, errTest)
			payloadBuf := new(bytes.Buffer)
			json.NewEncoder(payloadBuf).Encode(actualSensorTest)
			response := GetResponse("POST", "/sensor", payloadBuf, http.HandlerFunc(testSensor.Create))
			json.NewDecoder(io.Reader(response.Body)).Decode(&expectedSensor)
			Expect(response.Code).To(Equal(http.StatusBadRequest))
			Expect(expectedSensor).To(Equal(actualSensorTestEmpty))
		},
		)
	})
	Describe("Delete", func() {

		It("should delete sensor by ID when all passed parameters are valid", func() {
			mockTest.On("Delete", "").Return(nil)
			response := GetResponse("DELETE", "/sensor/1", nil, http.HandlerFunc(testSensor.Delete))
			Expect(response.Code).To(Equal(http.StatusNoContent))
		},
		)
		It("should return error when wrong parameters are passed", func() {
			mockTest.On("Delete", "").Return(errTest)
			response := GetResponse("DELETE", "/sensor/2", nil, http.HandlerFunc(testSensor.Delete))
			Expect(response.Code).To(Equal(http.StatusBadRequest))
		},
		)
	})
	Describe("Update", func() {

		It("should return updated sensor and nil when all passed parameters are valid", func() {
			mockTest.On("Update", "", deviceTest, actualSensorTest).Return(actualSensorTest, nil)
			payloadBuf := new(bytes.Buffer)
			json.NewEncoder(payloadBuf).Encode(actualSensorTest)
			response := GetResponse("PUT", "/sensor/1", payloadBuf, http.HandlerFunc(testSensor.Update))
			json.NewDecoder(io.Reader(response.Body)).Decode(&expectedSensor)
			Expect(response.Code).To(Equal(http.StatusOK))
			Expect(expectedSensor).To(Equal(actualSensorTest))
		},
		)
		It("should return nil and error when wrong parameters are passed", func() {
			mockTest.On("Update", "", deviceTest, actualSensorTest).Return(nil, errTest)
			payloadBuf := new(bytes.Buffer)
			json.NewEncoder(payloadBuf).Encode(actualSensorTest)
			response := GetResponse("PUT", "/sensor/2", payloadBuf, http.HandlerFunc(testSensor.Update))
			json.NewDecoder(io.Reader(response.Body)).Decode(&expectedSensor)
			Expect(response.Code).To(Equal(http.StatusBadRequest))
			Expect(expectedSensor).To(Equal(actualSensorTestEmpty))
		},
		)
	})
})
