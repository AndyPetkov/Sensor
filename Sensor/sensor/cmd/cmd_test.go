package cmd

import (
	"errors"
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("cmd", func() {

	Describe("executeThroughTime", func() {
		DescribeTable("should not print empty string when valid parametars are passed", func(format string, timeBetweenMeasurmants int, duration int, sensor_group []string, output_file string, web_hook_url string) {
			r, w, _ := os.Pipe()
			tmp := os.Stdout
			defer func() {
				os.Stdout = tmp
			}()
			os.Stdout = w
			var err error
			go func() {
				err = executeThroughTime(timeBetweenMeasurmants, duration, format, sensor_group, output_file, web_hook_url)
				w.Close()
			}()
			stdout, _ := ioutil.ReadAll(r)

			Expect(string(stdout)).NotTo(Equal(""))
			Expect(err).To(BeNil())

		},

			Entry("Then executeThroughTime prints", "JSON", 3, 9, []string{"CPU_USAGE", "CPU_TEMP"}, "ANDY", "http://localhost:8086"),
			Entry("Then executeThroughTime prints", "JSON", 4, 15, []string{"CPU_USAGE", "MEMORY_USAGE"}, "ANDY", "http://localhost:8086"),
			Entry("Then executeThroughTime prints", "YAML", 5, 18, []string{"CPU_USAGE"}, "ANDY", "http://localhost:8086"),
		)
	})
	Describe("validate", func() {
		It("should not return error when valid parametars are passed", func() {
			err := validate(5, 2, "YAML")
			errOk := validate(2, 10, "JSON")
			errSimulate := errors.New("error delta_duration cannot be higher than total_duration")
			Expect(err).To(Equal(errSimulate))
			Expect(errOk).To(BeNil())
		},
		)
	})
	Describe("print", func() {
		It("should return suitable error when parametars are passed", func() {
			err := print("JSON", []string{"CPU_USAGE", "CPU_TEMP"}, "ANDY", "http://localhost:8086")
			err1 := print("JSON", []string{"CPU_AGE", "CPU_TEMP"}, "ANDY", "http://localhost:8086")
			err1Simulate := errors.New("there is no CPU_AGE in the yaml file")

			Expect(err).To(BeNil())
			Expect(err1).To(Equal(err1Simulate))
		},
		)
	})

})
