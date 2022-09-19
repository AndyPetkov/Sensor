package service

import (
	"errors"
	"server/global"
	"server/mocks"
	"server/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("service-device", func() {
	var (
		actualDeviceTest models.Device
		errTest          error
		testDevice       Service
		mockTest         *mocks.Repository
		sensorTest       models.Sensor
	)
	BeforeEach(func() {
		actualDeviceTest = models.Device{
			Id:          "1",
			Name:        "my_device",
			Description: "home_laptop",
			Sensors:     nil,
		}
		errTest = errors.New("")
		mockTest = &mocks.Repository{}
		testDevice = NewService(global.Device)
		testDevice.ConfigureRepo(mockTest)
	})

	Describe("GetAll", func() {

		It("should return interface and nil", func() {
			mockTest.On("GetAll").Return(actualDeviceTest, nil)
			expectedResult, expectedErr := testDevice.GetAll()
			Expect(expectedErr).To(BeNil())
			Expect(expectedResult).To(Equal(actualDeviceTest))
		},
		)
		It("should return nil and error", func() {
			mockTest.On("GetAll").Return(nil, errTest)
			expectedResult, expectedErr := testDevice.GetAll()
			Expect(expectedErr).ToNot(BeNil())
			Expect(expectedResult).To(BeNil())
		},
		)
	})
	Describe("GetById", func() {

		It("should return interface and nil", func() {
			mockTest.On("GetById", "1").Return(actualDeviceTest, nil)
			expectedResult, expectedErr := testDevice.GetById("1")
			Expect(expectedErr).To(BeNil())
			Expect(expectedResult).To(Equal(actualDeviceTest))
		},
		)
		It("should return nil and error", func() {
			mockTest.On("GetById", "2").Return(nil, errTest)
			result, expectedErr := testDevice.GetById("2")
			Expect(expectedErr).ToNot(BeNil())
			Expect(result).To(BeNil())
		},
		)
		FIt("test", func() {
			var match []bool
			match[0] = true
			match[1] = true
			match[2] = false
			match[3] = true
			Expect(match).To(Equal(true))
		},
		)
	})
	Describe("Create", func() {

		It("should return interface and nil", func() {
			mockTest.On("Create", actualDeviceTest, sensorTest).Return(actualDeviceTest, nil)
			expectedResult, expectedErr := testDevice.Create(actualDeviceTest, sensorTest)
			Expect(expectedErr).To(BeNil())
			Expect(expectedResult).To(Equal(actualDeviceTest))
		},
		)
		It("should return nil and error", func() {
			mockTest.On("Create", actualDeviceTest, sensorTest).Return(nil, errTest)
			result, expectedErr := testDevice.Create(actualDeviceTest, sensorTest)
			Expect(expectedErr).ToNot(BeNil())
			Expect(result).To(BeNil())
		},
		)
	})
	Describe("Delete", func() {

		It("should return interface and nil", func() {
			mockTest.On("Delete", "1").Return(nil)
			expectedErr := testDevice.Delete("1")
			Expect(expectedErr).To(BeNil())
		},
		)
		It("should return nil and error", func() {
			mockTest.On("Delete", "2").Return(errTest)
			expectedErr := testDevice.Delete("2")
			Expect(expectedErr).ToNot(BeNil())
		},
		)
	})
	Describe("Update", func() {

		It("should return interface and nil", func() {
			mockTest.On("Update", "1", actualDeviceTest, sensorTest).Return(actualDeviceTest, nil)
			expectedResult, expectedErr := testDevice.Update("1", actualDeviceTest, sensorTest)
			Expect(expectedErr).To(BeNil())
			Expect(expectedResult).To(Equal(actualDeviceTest))
		},
		)
		It("should return nil and error", func() {
			mockTest.On("Update", "2", actualDeviceTest, sensorTest).Return(nil, errTest)
			expectedResult, expectedErr := testDevice.Update("2", actualDeviceTest, sensorTest)
			Expect(expectedErr).ToNot(BeNil())
			Expect(expectedResult).To(BeNil())
		},
		)
	})
})
