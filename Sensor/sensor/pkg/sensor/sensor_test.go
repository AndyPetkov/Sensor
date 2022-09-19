package sensor

import (
	"math"
	"sensor/pkg/mock"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("sensor", func() {
	BaseExecutor = &mock.MockExecutor{}
	Describe("validate", func() {
		DescribeTable("should return error message when wrong parametars are passed", func(unit string, errormsg string) {
			err := validate(unit)
			Expect(err).To(MatchError(ContainSubstring(errormsg)))
		},
			Entry("Then it prints unit cannot be A", "A", "unit cannot be A"),
		)
		It("should return nil when parametars are valid", func() {
			err := validate("F")
			Expect(err).To(BeNil())
		})

	})
	Describe("GetTemp", func() {
		DescribeTable("should return positive temperature when parametars are valid", func(unit string, errorr string) {
			var executor Executor
			executor = NewЕxecutor()
			got, err := executor.GetTemp(unit)
			checkNegative := math.Signbit(got)
			Expect(checkNegative).To(BeFalse())
			Expect(err).To(BeNil())

		},
			Entry("Then error is equal to nil and temperature is positive", "F", nil),
			Entry("Then error is equal to nil and temperature is positive", "C", nil),
		)
	})

	Describe("GetMetrics", func() {
		It("should equal memory usage test function", func() {
			var metricer Metricer
			metricer, _ = NewSensor("MEMORY_USAGE")
			var measurments []*Measurement
			measurments, err := metricer.GetMetrics()
			total := NewMeasurement(float64(39257714688), "", "")
			available := NewMeasurement(float64(19993272320), "", "")
			used := NewMeasurement(float64(19264442368), "", "")
			usedPercent := NewMeasurement(float64(49.071736653811406), "", "")
			Expect(measurments[0]).To(Equal(total))
			Expect(measurments[1]).To(Equal(available))
			Expect(measurments[2]).To(Equal(used))
			Expect(measurments[3]).To(Equal(usedPercent))
			Expect(err).To(BeNil())

		})
		It("should equal cpu usage test function", func() {
			var metricer Metricer
			metricer, _ = NewSensor("CPU_USAGE")
			var measurments []*Measurement
			measurments, err := metricer.GetMetrics()
			percent := NewMeasurement(float64(57.00), "", "")
			numOfCores := NewMeasurement(float64(8.0), "", "")
			mhz := NewMeasurement(float64(5456.4556), "", "")
			Expect(measurments[0]).To(Equal(percent))
			Expect(measurments[1]).To(Equal(numOfCores))
			Expect(measurments[2]).To(Equal(mhz))
			Expect(err).To(BeNil())
		})
		It("should equal temp test function", func() {
			var metricer Metricer
			metricer, _ = NewSensor("CPU_TEMP")
			var measurments []*Measurement
			measurments, err := metricer.GetMetrics()
			temp := NewMeasurement(60.00, "", "")
			Expect(measurments[0]).To(Equal(temp))
			Expect(err).To(BeNil())
		})
	})

})
