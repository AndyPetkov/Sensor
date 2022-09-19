package service

import (
	"errors"
	"server/global"
	"server/mocks"
	"server/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("service-sensor", func() {
	var (
		actualDeviceTest models.Device
		errTest          error
		testSensor       Service
		mockTest         *mocks.Repository
		actualSensorTest models.Sensor
	)
	BeforeEach(func() {
		actualSensorTest = models.Sensor{
			Id:          "1",
			DeviceId:    "1",
			Name:        "cpuTempCelsius",
			Description: "Measures CPU temp Celsius",
			Unit:        "C",
			Sensorgroup: "CPU_TEMP",
		}
		errTest = errors.New("")
		mockTest = &mocks.Repository{}
		testSensor = NewService(global.Sensor)
		testSensor.ConfigureRepo(mockTest)
	})

	Describe("GetAll", func() {

		It("should return interface and nil", func() {
			mockTest.On("GetAll").Return(actualDeviceTest, nil)
			expectedResult, expectedErr := testSensor.GetAll()
			Expect(expectedErr).To(BeNil())
			Expect(expectedResult).To(Equal(actualDeviceTest))
		},
		)
		It("should return nil and error", func() {
			mockTest.On("GetAll").Return(nil, errTest)
			expectedResult, expectedErr := testSensor.GetAll()
			Expect(expectedErr).ToNot(BeNil())
			Expect(expectedResult).To(BeNil())
		},
		)
	})
	Describe("GetById", func() {

		It("should return interface and nil", func() {
			mockTest.On("GetById", "1").Return(actualDeviceTest, nil)
			expectedResult, expectedErr := testSensor.GetById("1")
			Expect(expectedErr).To(BeNil())
			Expect(expectedResult).To(Equal(actualDeviceTest))
		},
		)
		It("should return nil and error", func() {
			mockTest.On("GetById", "2").Return(nil, errTest)
			expectedResult, expectedErr := testSensor.GetById("2")
			Expect(expectedErr).ToNot(BeNil())
			Expect(expectedResult).To(BeNil())
		},
		)
	})
	Describe("Create", func() {

		It("should return interface and nil", func() {
			mockTest.On("Create", actualDeviceTest, actualSensorTest).Return(actualDeviceTest, nil)
			expectedResult, expectedErr := testSensor.Create(actualDeviceTest, actualSensorTest)
			Expect(expectedErr).To(BeNil())
			Expect(expectedResult).To(Equal(actualDeviceTest))
		},
		)
		It("should return nil and error", func() {
			mockTest.On("Create", actualDeviceTest, actualSensorTest).Return(nil, errTest)
			expectedResult, expectedErr := testSensor.Create(actualDeviceTest, actualSensorTest)
			Expect(expectedErr).ToNot(BeNil())
			Expect(expectedResult).To(BeNil())
		},
		)
	})
	Describe("Delete", func() {

		It("should return interface and nil", func() {
			mockTest.On("Delete", "1").Return(nil)
			expectedErr := testSensor.Delete("1")
			Expect(expectedErr).To(BeNil())
		},
		)
		It("should return nil and error", func() {
			mockTest.On("Delete", "2").Return(errTest)
			expectedErr := testSensor.Delete("2")
			Expect(expectedErr).ToNot(BeNil())
		},
		)
	})
	Describe("Update", func() {

		It("should return interface and nil", func() {
			mockTest.On("Update", "1", actualDeviceTest, actualSensorTest).Return(actualDeviceTest, nil)
			expectedResult, expectedErr := testSensor.Update("1", actualDeviceTest, actualSensorTest)
			Expect(expectedErr).To(BeNil())
			Expect(expectedResult).To(Equal(actualDeviceTest))
		},
		)
		It("should return nil and error", func() {
			mockTest.On("Update", "2", actualDeviceTest, actualSensorTest).Return(nil, errTest)
			expectedResult, expectedErr := testSensor.Update("2", actualDeviceTest, actualSensorTest)
			Expect(expectedErr).ToNot(BeNil())
			Expect(expectedResult).To(BeNil())
		},
		)
	})
})
