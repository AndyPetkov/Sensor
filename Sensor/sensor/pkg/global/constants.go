package global

var Bytes string = "Bytes"
var TempMeasure string = "TempMeasurement"
var CPU_USAGE string = "CPU_USAGE"
var MEMORY_USAGE string = "MEMORY_USAGE"
var CPU_TEMP string = "CPU_TEMP"
var JSON string = "JSON"
var YAML string = "YAML"
var Celsius string = "C"
var Fahrenheit string = "F"
var DeviceName string = "device_name"
var CpuTempCelsius string = "cpuTempCelsius"
var CpuUsagePercent string = "cpuUsagePercent"
var CpuCoresCount string = "cpuCoresCount"
var CpuFrequency string = "cpuFrequency"
var MemoryTotal string = "memoryTotal"
var MemoryAvailableBytes string = "memoryAvailableBytes"
var MemoryUsedBytes string = "memoryUsedBytes"
var MemoryUsedPercent string = "memoryUsedPercent"
var SensorGroups map[string]bool = map[string]bool{
	CpuTempCelsius:       true,
	CpuUsagePercent:      true,
	CpuCoresCount:        true,
	CpuFrequency:         true,
	MemoryTotal:          true,
	MemoryAvailableBytes: true,
	MemoryUsedBytes:      true,
	MemoryUsedPercent:    true,
}
var SensorGroupName map[string]bool = map[string]bool{
	CPU_TEMP:     true,
	MEMORY_USAGE: true,
	CPU_USAGE:    true,
}
var DeviceGroups map[string]bool = map[string]bool{
	DeviceName: true,
}
var FileExist bool = false
