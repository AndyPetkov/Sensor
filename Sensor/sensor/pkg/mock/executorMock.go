package mock

type MockExecutor struct{}

func (m *MockExecutor) GetMemory() (uint64, uint64, uint64, float64, error) {
	var total uint64 = 39257714688
	var available uint64 = 19993272320
	var used uint64 = 19264442368
	var usedPercent float64 = 49.071736653811406
	var err error = nil
	return total, available, used, usedPercent, err
}
func (t *MockExecutor) GetTemp(unit string) (float64, error) {
	temp := 60.00
	if unit == "F" {
		temp = (temp * 1.8) + 32
	}
	return temp, nil
}
func (c *MockExecutor) GetCPU() (float64, float64, float64, error) {
	percent := 57.00
	numOfCores := 8.0
	mhz := 5456.4556
	return percent, numOfCores, mhz, nil
}
