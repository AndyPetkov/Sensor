package load

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("sensor", func() {

	Describe("GetSensorID", func() {
		DescribeTable("should return sensorId and unit", func(sensorGroupName string, sensorIdExpected string, unitExpected string) {
			sensorId, unit := GetSensorID(sensorGroupName)
			Expect(sensorId).To(Equal(sensorIdExpected))
			Expect(unit).To(Equal(unitExpected))
		},
			Entry("Then it returns sensorId=1 and unit=C", "cpuTempCelsius", "1", "C"),
			Entry("Then it returns sensorId=2 and unit=C", "cpuUsagePercent", "2", "%"),
			Entry("Then it returns sensorId=2 and unit=C", "cpuCoresCount", "3", "count"),
			Entry("Then it returns sensorId=2 and unit=C", "memoryAvailableBytes", "6", "Bytes"),
		)
	})
	Describe("GetDeviceID", func() {
		DescribeTable("should return sensorId and unit", func(deviceName string, deviceIdExpected string) {
			deviceId := GetDeviceID(deviceName)
			Expect(deviceId).To(Equal(deviceIdExpected))
		},
			Entry("Then it returns deviceId=1", "device_name", "1"),
		)
	})

})
