package cmd

import (
	"errors"
	"fmt"
	"os"

	"sensor/pkg/client"
	"sensor/pkg/global"
	"sensor/pkg/logger"
	"sensor/pkg/sensor"
	"sensor/pkg/util"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	unit           string
	total_duration int
	delta_duration int
	format         string
	sensor_group   []string
	output_file    string
	web_hook_url   string
)

var Verbose bool
var Source string
var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "sensor",
	Short: "Measuring CPU temperature",
	Long:  ` Golang CLI application called sensor, which measures the CPU temperature of your local laptop`,

	RunE: func(cmd *cobra.Command, args []string) error {

		logger.Infologgers.InfoLogger.Println("Starting the application...")
		var err error
		err = validate(delta_duration, total_duration, format)
		if err != nil {
			logger.Errorloggers.ErrorLogger.Println(err)
			return err
		}
		err = executeThroughTime(delta_duration, total_duration, format, sensor_group, output_file, web_hook_url)
		if err != nil {
			logger.Errorloggers.ErrorLogger.Println(err)
			return err
		}
		logger.Infologgers.InfoLogger.Println("End of the application")
		return err

	},
}

func executeThroughTime(timeBetweenMeasurmants int, duration int, format string, sensor_group []string, output_file string, web_hook_url string) error {
	ticker := time.NewTicker(time.Duration(timeBetweenMeasurmants) * time.Second)
	defer ticker.Stop()
	timeout := time.After(time.Duration(duration) * time.Second)
	var err error
	for {
		select {
		case <-timeout:
			logger.Infologgers.InfoLogger.Println("Finish printing")
			return err
		case <-ticker.C:
			if err != nil {
				errString := fmt.Errorf("error getting metrics %w", err)
				logger.Errorloggers.ErrorLogger.Println(errString)
				return err
			}
			err := print(format, sensor_group, output_file, web_hook_url)
			if err != nil {
				return err
			}
			logger.Infologgers.InfoLogger.Println("Printing measurmants...")
		}
	}
}

func print(format string, sensor_group []string, output_file string, web_hook_url string) error {
	for i := 0; i < len(sensor_group); i++ {
		_, found := global.SensorGroupName[sensor_group[i]]
		if found == false {
			err := errors.New("there is no " + sensor_group[i] + " in the yaml file")
			logger.Errorloggers.ErrorLogger.Println(err)
			return err
		}
	}
	for i := 0; i < len(sensor_group); i++ {
		var measurements []*sensor.Measurement
		measurements, err := getMeasurements(sensor_group[i])
		if err != nil {
			errString := fmt.Errorf("error getting metrics %w", err)
			logger.Errorloggers.ErrorLogger.Println(errString)
			return err
		}
		if web_hook_url != "" {
			for i := 0; i < len(measurements); i++ {
				err := client.PostData(web_hook_url, measurements[i])
				if err != nil {
					errString := fmt.Errorf("error posting data to server %w", err)
					logger.Errorloggers.ErrorLogger.Println(errString)
					return err
				}
			}
		}
		messages := make(chan []string)
		go util.Print(format, measurements, messages)
		if output_file != "" {
			data := <-messages
			err := util.WriteToCSV(output_file, data)
			if err != nil {
				errString := fmt.Errorf("error writing to file %w", err)
				logger.Errorloggers.ErrorLogger.Println(errString)
				return err
			}
		}
	}
	return nil
}
func getMeasurements(sensor_group string) ([]*sensor.Measurement, error) {
	var measurements []*sensor.Measurement
	var err error
	var metricer sensor.Metricer
	metricer, err = sensor.NewSensor(sensor_group)
	if err != nil {
		errString := fmt.Errorf("error getting instance %w", err)
		logger.Errorloggers.ErrorLogger.Println(errString)
		return nil, err
	}
	measurements, err = metricer.GetMetrics()
	if err != nil {
		errString := fmt.Errorf("error getting metrics %w", err)
		logger.Errorloggers.ErrorLogger.Println(errString)
		return nil, err
	}
	return measurements, err
}
func validate(delta_duration int, total_duration int, format string) error {
	var err error
	if delta_duration > total_duration {
		err = errors.New("error delta_duration cannot be higher than total_duration")
		logger.Errorloggers.ErrorLogger.Println(err)
		return err
	}
	if format != global.YAML && format != global.JSON {
		return errors.New("format cannot be " + format)
	}
	return nil
}
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {

	cobra.OnInitialize(initConfig)

	rootCmd.Flags().IntVar(&total_duration, "total_duration", 15, "Enter duration of the program")
	rootCmd.Flags().IntVar(&delta_duration, "delta_duration", 3, "Enter duration between the measurments")
	rootCmd.Flags().StringVar(&format, "format", "JSON", "Choose between JSON or YAML")
	rootCmd.Flags().StringSliceVar(&sensor_group, "sensor_group", []string{"CPU_USAGE", "MEMORY_USAGE"}, "Choose between CPU_USAGE, MEMORY_USAGE and CPU_TEMP")
	rootCmd.Flags().StringVar(&output_file, "output_file", "", "Enter file name")
	rootCmd.Flags().StringVar(&web_hook_url, "web_hook_url", "", "Enter url")
}

func initConfig() {
	if cfgFile != "" {

		viper.SetConfigFile(cfgFile)
	} else {

		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".sensor")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
