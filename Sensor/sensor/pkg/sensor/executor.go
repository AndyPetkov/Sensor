package sensor

import (
	"errors"
	"fmt"
	"sensor/pkg/logger"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type executor struct {
}

type Executor interface {
	GetCPU() (float64, float64, float64, error)
	GetMemory() (uint64, uint64, uint64, float64, error)
	GetTemp(unit string) (float64, error)
}

var BaseExecutor Executor = &executor{}

func NewЕxecutor() Executor {
	return BaseExecutor
}
func (m *executor) GetMemory() (uint64, uint64, uint64, float64, error) {
	memory, err := mem.SwapMemory()
	if err != nil {
		return 0, 0, 0, 0, errors.New("problem getting memory")
	}
	total := memory.Total
	available := memory.Free
	used := memory.Used
	usedPercent := memory.UsedPercent
	return total, available, used, usedPercent, nil
}
func (t *executor) GetTemp(unit string) (float64, error) {
	temperature, err := host.SensorsTemperatures()
	if err != nil {
		errString := fmt.Errorf("error getting temperature %w", err)
		logger.Errorloggers.ErrorLogger.Print(errString)
		return 0.0, err
	}
	temp := temperature[0].Temperature
	if unit == "F" {
		logger.Infologgers.InfoLogger.Println("Converting into fahrenheit")
		temp = (temp * 1.8) + 32
	}
	return temp, nil
}
func (c *executor) GetCPU() (float64, float64, float64, error) {
	percent, err := cpu.Percent(5, false)
	if err != nil {
		return 0, 0, 0, errors.New("problem getting percentage")
	}
	numOfCores, err := cpu.Info()
	if err != nil {
		return 0, 0, 0, errors.New("problem getting number of cores")
	}
	mhz, err := cpu.Info()
	if err != nil {
		return 0, 0, 0, errors.New("problem getting frequency")
	}
	return percent[0], float64(numOfCores[0].Cores), mhz[0].Mhz, err
}
