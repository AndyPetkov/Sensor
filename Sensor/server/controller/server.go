package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")

}

func HandleRequests() error {

	device := NewController("device")
	sensor := NewController("sensor")
	measurement := NewControllerMeasurement()
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Use(commonMiddleware)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/measurements", measurement.GetAll).Methods("GET")
	myRouter.HandleFunc("/measurements", measurement.Create).Methods("POST")
	myRouter.HandleFunc("/sensorAvarageValue", measurement.SensorAvarageValue).Methods("GET")
	myRouter.HandleFunc("/sensorsCorrelationCoefficient", measurement.SensorsCorrelationCoefficient).Methods("GET")
	myRouter.HandleFunc("/devices", device.GetAll).Methods("GET")
	myRouter.HandleFunc("/device/{id}", device.GetById).Methods("GET")
	myRouter.HandleFunc("/device", device.Create).Methods("POST")
	myRouter.HandleFunc("/device/{id}", device.Update).Methods("PUT")
	myRouter.HandleFunc("/device/{id}", device.Delete).Methods("DELETE")
	myRouter.HandleFunc("/sensors", sensor.GetAll).Methods("GET")
	myRouter.HandleFunc("/sensor/{id}", sensor.GetById).Methods("GET")
	myRouter.HandleFunc("/sensor", sensor.Create).Methods("POST")
	myRouter.HandleFunc("/sensor/{id}", sensor.Update).Methods("PUT")
	myRouter.HandleFunc("/sensor/{id}", sensor.Delete).Methods("DELETE")

	err := http.ListenAndServe(":4000", myRouter)
	return err
}
func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
