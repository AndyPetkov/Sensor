package util

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"sensor/pkg/global"
	"sensor/pkg/logger"
	"sensor/pkg/sensor"

	"gopkg.in/yaml.v2"
)

func WriteToCSV(output_file string, data []string) error {
	var csvFile *os.File
	var err error
	if global.FileExist == false {
		csvFile, err = os.Create(output_file + ".csv")
		if err != nil {
			errString := fmt.Errorf("error creating csv file")
			logger.Errorloggers.ErrorLogger.Print(errString)

			return err
		}
		csvwriter := csv.NewWriter(csvFile)
		csvwriter.Comma = '|'
		_ = csvwriter.Write(data)
		csvwriter.Flush()
		global.FileExist = true
	} else {
		csvFile, err = os.OpenFile(output_file+".csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			errString := fmt.Errorf("error opening csv file")
			logger.Errorloggers.ErrorLogger.Print(errString)
			return err
		}
		csvwriter := csv.NewWriter(csvFile)
		csvwriter.Comma = '|'
		_ = csvwriter.Write(data)
		csvwriter.Flush()
	}
	csvFile.Close()
	return err
}

func Print(format string, structs []*sensor.Measurement, message chan []string) {
	var data []string
	var output string
	if format == "JSON" {
		for i := 0; i < len(structs); i++ {
			output, _ = JSONparser(structs[i])
			fmt.Println(output)
			data = append(data, output)
		}
		message <- data
	}
	if format == "YAML" {
		for i := 0; i < len(structs); i++ {
			output, _ = YAMLparser(structs[i])
			fmt.Println(output)
			data = append(data, output)
		}
		message <- data
	}
	close(message)
}
func JSONparser(m interface{}) (string, error) {
	data, err := json.Marshal(m)
	logger.Infologgers.InfoLogger.Println("Parsing in JSON")

	if err != nil {
		errString := fmt.Errorf("error converting output to JOSN %w", err)
		logger.Errorloggers.ErrorLogger.Print(errString)
		return "", err
	}
	return string(data), nil
}
func YAMLparser(m interface{}) (string, error) {

	logger.Infologgers.InfoLogger.Println("Parsing in YAML")
	data, err := yaml.Marshal(m)
	if err != nil {
		errString := fmt.Errorf("error converting output to YAML %w", err)
		logger.Errorloggers.ErrorLogger.Print(errString)
		return "", err
	}
	return string(data), nil
}
