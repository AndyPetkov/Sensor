package service

import (
	"errors"
	"server/mocks"
	"server/models"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("service-measurement", func() {
	var (
		measurementTest models.Measurement
		errTest         error
		testMeasurement MeasurementService
		mockTest        *mocks.MeasurementRepo
	)
	BeforeEach(func() {
		measurementTest = models.Measurement{
			Time:     time.Now(),
			Value:    35,
			SensorID: "1",
			DeviceID: "1",
		}
		errTest = errors.New("")
		mockTest = &mocks.MeasurementRepo{}
		testMeasurement = NewServiceMeasurement()
		testMeasurement.ConfigureRepoMeasurement(mockTest)
	})

	Describe("GetAll", func() {

		It("should return interface and nil", func() {
			mockTest.On("GetAll").Return(measurementTest, nil)
			expectedResult, expectedErr := testMeasurement.GetAll()
			Expect(expectedErr).To(BeNil())
			Expect(expectedResult).To(Equal(measurementTest))
		},
		)
		It("should return nil and error", func() {
			mockTest.On("GetAll").Return(nil, errTest)
			expectedResult, expectedErr := testMeasurement.GetAll()
			Expect(expectedErr).ToNot(BeNil())
			Expect(expectedResult).To(BeNil())
		},
		)
	})
	Describe("Create", func() {

		It("should return interface and nil", func() {
			mockTest.On("Create", measurementTest).Return(measurementTest, nil)
			expectedResult, expectedErr := testMeasurement.Create(measurementTest)
			Expect(expectedErr).To(BeNil())
			Expect(expectedResult).To(Equal(measurementTest))
		},
		)
		It("should return nil and error", func() {
			measurementTestNotValid := models.Measurement{
				Time:     time.Now(),
				Value:    35,
				SensorID: "",
				DeviceID: "1",
			}
			expectedResult, expectedErr := testMeasurement.Create(measurementTestNotValid)
			errSimulate := errors.New("sensorID cannot be empty")
			Expect(expectedErr).To(Equal(errSimulate))
			Expect(expectedResult).To(BeNil())
		},
		)
	})
	Describe("SensorAvarageValue", func() {

		It("should return float64 and nil", func() {
			actualValue := 55.0
			mockTest.On("SensorAvarageValue", "1", "1", "2021-12-02", "2021-12-03").Return(actualValue, nil)
			expectedResult, expectedErr := testMeasurement.SensorAvarageValue("1", "1", "2021-12-02", "2021-12-03")
			Expect(expectedErr).To(BeNil())
			Expect(expectedResult).To(Equal(actualValue))
		},
		)
		It("should return 0 and error", func() {
			mockTest.On("SensorAvarageValue", "1", "1", "2021-12-02", "2021-12-03").Return(0.0, errTest)
			expectedResult, expectedErr := testMeasurement.SensorAvarageValue("1", "1", "2021-12-02", "2021-12-03")
			Expect(expectedErr).ToNot(BeNil())
			Expect(expectedResult).To(Equal(0.0))
		},
		)
	})
	Describe("SensorsCorrelationCoefficient", func() {

		It("should return float64 and nil", func() {
			actualValue := 0.554
			mockTest.On("SensorsCorrelationCoefficient", "1", "1", "2", "1", "2021-12-02", "2021-12-03").Return(actualValue, nil)
			expectedResult, expectedErr := testMeasurement.SensorsCorrelationCoefficient("1", "1", "2", "1", "2021-12-02", "2021-12-03")
			Expect(expectedErr).To(BeNil())
			Expect(expectedResult).To(Equal(actualValue))
		},
		)
		It("should return 0.0 and error", func() {
			mockTest.On("SensorsCorrelationCoefficient", "1", "1", "2", "1", "2021-12-02", "2021-12-03").Return(0.0, errTest)
			expectedResult, expectedErr := testMeasurement.SensorsCorrelationCoefficient("1", "1", "2", "1", "2021-12-02", "2021-12-03")
			Expect(expectedErr).ToNot(BeNil())
			Expect(expectedResult).To(Equal(0.0))
		},
		)
	})

})
