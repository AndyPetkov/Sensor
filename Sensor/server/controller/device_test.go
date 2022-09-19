package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"server/global"
	"server/mocks"
	"server/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func GetResponse(method string, url string, body io.Reader, handler http.HandlerFunc) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, url, body)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	return response
}

var _ = Describe("controller-device", func() {

	var (
		actualDeviceTest models.Device
		deviceTestEmpty  models.Device
		errTest          error
		testDevice       Controller
		mockTest         *mocks.Service
		sensorTest       models.Sensor
		expectedDevice   models.Device
	)
	BeforeEach(func() {
		actualDeviceTest = models.Device{Id: "17", Name: "my_device", Description: "home_laptop", Sensors: nil}
		deviceTestEmpty = models.Device{Id: "", Name: "", Description: "", Sensors: nil}
		errTest = errors.New("")
		mockTest = &mocks.Service{}
		testDevice = NewController(global.Device)
		testDevice.ConfigureService(mockTest)
		expectedDevice = models.Device{}
	})

	Describe("GetAll", func() {

		It("should return all devices and nil when all passed parameters are valid", func() {
			mockTest.On("GetAll").Return(actualDeviceTest, nil)
			response := GetResponse("GET", "/devices", nil, http.HandlerFunc(testDevice.GetAll))
			json.NewDecoder(io.Reader(response.Body)).Decode(&expectedDevice)
			Expect(response.Code).To(Equal(http.StatusOK))
			Expect(expectedDevice).To(Equal(actualDeviceTest))
		},
		)
		It("should return nil and error when wrong parameters are passed", func() {
			mockTest.On("GetAll").Return(nil, errTest)
			response := GetResponse("GET", "/devices", nil, http.HandlerFunc(testDevice.GetAll))
			json.NewDecoder(io.Reader(response.Body)).Decode(&expectedDevice)
			Expect(response.Code).To(Equal(http.StatusBadRequest))
			Expect(expectedDevice).To(Equal(deviceTestEmpty))
		},
		)
	})
	Describe("GetById", func() {

		It("should return device by ID and nil when all passed parameters are valid", func() {
			mockTest.On("GetById", "").Return(actualDeviceTest, nil)
			response := GetResponse("GET", "/device/1", nil, http.HandlerFunc(testDevice.GetById))
			json.NewDecoder(io.Reader(response.Body)).Decode(&expectedDevice)
			Expect(response.Code).To(Equal(http.StatusOK))
			Expect(expectedDevice).To(Equal(actualDeviceTest))
		},
		)
		It("should return nil and error when wrong parameters are passed", func() {
			mockTest.On("GetById", "").Return(nil, errTest)
			response := GetResponse("GET", "/device/2", nil, http.HandlerFunc(testDevice.GetById))
			json.NewDecoder(io.Reader(response.Body)).Decode(&expectedDevice)
			Expect(response.Code).To(Equal(http.StatusBadRequest))
			Expect(expectedDevice).To(Equal(deviceTestEmpty))
		},
		)
	})
	Describe("Create", func() {

		It("should create device and return nil when all passed parameters are valid", func() {
			mockTest.On("Create", actualDeviceTest, sensorTest).Return(actualDeviceTest, nil)
			payloadBuf := new(bytes.Buffer)
			json.NewEncoder(payloadBuf).Encode(actualDeviceTest)
			response := GetResponse("POST", "/device", payloadBuf, http.HandlerFunc(testDevice.Create))
			json.NewDecoder(io.Reader(response.Body)).Decode(&expectedDevice)
			Expect(response.Code).To(Equal(http.StatusOK))
			Expect(expectedDevice).To(Equal(actualDeviceTest))
		},
		)
		It("should return nil and error when wrong parameters are passed", func() {
			mockTest.On("Create", actualDeviceTest, sensorTest).Return(nil, errTest)
			payloadBuf := new(bytes.Buffer)
			json.NewEncoder(payloadBuf).Encode(actualDeviceTest)
			response := GetResponse("POST", "/device", payloadBuf, http.HandlerFunc(testDevice.Create))
			json.NewDecoder(io.Reader(response.Body)).Decode(&expectedDevice)
			Expect(response.Code).To(Equal(http.StatusBadRequest))
			Expect(expectedDevice).To(Equal(deviceTestEmpty))
		},
		)
	})
	Describe("Delete", func() {

		It("should delete device by ID when all passed parameters are valid", func() {
			mockTest.On("Delete", "").Return(nil)
			response := GetResponse("DELETE", "/device/1", nil, http.HandlerFunc(testDevice.Delete))
			Expect(response.Code).To(Equal(http.StatusNoContent))
		},
		)
		It("should return error when wrong parameters are passed", func() {
			mockTest.On("Delete", "").Return(errTest)
			response := GetResponse("DELETE", "/device/2", nil, http.HandlerFunc(testDevice.Delete))
			Expect(response.Code).To(Equal(http.StatusBadRequest))
		},
		)
	})
	Describe("Update", func() {

		It("should return updated device and nil when all passed parameters are valid", func() {
			mockTest.On("Update", "", actualDeviceTest, sensorTest).Return(actualDeviceTest, nil)
			payloadBuf := new(bytes.Buffer)
			json.NewEncoder(payloadBuf).Encode(actualDeviceTest)
			response := GetResponse("PUT", "/device/1", payloadBuf, http.HandlerFunc(testDevice.Update))
			json.NewDecoder(io.Reader(response.Body)).Decode(&expectedDevice)
			Expect(response.Code).To(Equal(http.StatusOK))
			Expect(expectedDevice).To(Equal(actualDeviceTest))
		},
		)
		It("should return nil and error when wrong parameters are passed", func() {
			mockTest.On("Update", "", actualDeviceTest, sensorTest).Return(nil, errTest)
			payloadBuf := new(bytes.Buffer)
			json.NewEncoder(payloadBuf).Encode(actualDeviceTest)
			response := GetResponse("PUT", "/device/2", payloadBuf, http.HandlerFunc(testDevice.Update))
			json.NewDecoder(io.Reader(response.Body)).Decode(&expectedDevice)
			Expect(response.Code).To(Equal(http.StatusBadRequest))
			Expect(expectedDevice).To(Equal(deviceTestEmpty))
		},
		)
	})
})
