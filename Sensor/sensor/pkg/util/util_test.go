package util

import (
	. "github.com/onsi/ginkgo/extensions/table"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Util", func() {
	type testMeasurement struct {
		Time  string  `json:"measuredAt"  yaml:"measuredAt"`
		Value float64 `json:"value:" yaml:"value"`
		Type  string  `json:"unit:" yaml:"unit"`
	}
	Describe("JSONparser", func() {
		DescribeTable("Checks if JSONparser returns the same metrics", func(unit string, temperatureOfCPU float64) {
			smallMeasurment := testMeasurement{
				Time:  "16:24:58",
				Value: temperatureOfCPU,
				Type:  unit,
			}
			got, err := JSONparser(smallMeasurment)
			expected := "{\"measuredAt\":\"16:24:58\",\"value:\":64,\"unit:\":\"C\"}"
			Expect(got).To(Equal(expected))
			Expect(err).To(BeNil())
		},
			Entry("Then JSONparser has equal metrics", "C", 64.00),
		)

	})

	Describe("YAMLparser", func() {
		DescribeTable("Checks if YAMLparser returns the same metrics", func(unit string, temperatureOfCPU float64) {
			smallMeasurment := testMeasurement{
				Time:  "10:53:24",
				Value: temperatureOfCPU,
				Type:  unit,
			}
			got, err := YAMLparser(smallMeasurment)
			expected := "measuredAt: \"10:53:24\"\nvalue: 64\nunit: C\n"
			Expect(got).To(Equal(expected))
			Expect(err).To(BeNil())
		},
			Entry("Then YAMLparser has equal metrics", "C", 64.00),
		)
	})

	Describe("WriteToCSV", func() {
		DescribeTable("Checks if WriteToCSV returns error", func(output_file string, data []string) {
			err := WriteToCSV(output_file, data)
			Expect(err).To(BeNil())

		},
			Entry("Then WriteToCSV has nil error", "SAP", []string{"measuredAt 2021-11-03T15:30:40.4582944+02:00 value::57.067582957416995 sensorId:8 deviceId:",
				"measuredAt2021-11-03T15:30:40.4497012+02:00 value:8,sensorId:3 deviceId: 1"}),
		)
	})
})
