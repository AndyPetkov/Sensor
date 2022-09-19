package controller

import (
	"errors"
	"net/http"
	"server/logger"
	"server/service"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type sensorHandler struct {
	Service service.Service
}

func (s *sensorHandler) ConfigureService(service service.Service) {
	s.Service = service
}
func (s *sensorHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	logger.GetInstance().InfoLogger.Println("Getting all sensors...")
	result, err := s.Service.GetAll()
	if err != nil {
		writeResponse(w, err, result, http.StatusBadRequest)
		return
	}
	writeResponse(w, err, result, http.StatusOK)
	logger.GetInstance().InfoLogger.Println("All sensors got")
}
func (s *sensorHandler) GetById(w http.ResponseWriter, r *http.Request) {
	logger.GetInstance().InfoLogger.Println("Getting sensor by id...")
	params := mux.Vars(r)
	id := params["id"]
	if id == " " {
		err := errors.New("you are missing id parameter.")
		writeResponse(w, err, nil, http.StatusBadRequest)
		return
	}
	result, err := s.Service.GetById(id)
	if err != nil {
		writeResponse(w, err, result, http.StatusBadRequest)
		return
	}
	writeResponse(w, err, result, http.StatusOK)
	logger.GetInstance().InfoLogger.Println("Sensor by id got")
}
func (s *sensorHandler) Create(w http.ResponseWriter, r *http.Request) {
	logger.GetInstance().InfoLogger.Println("New sensor being created...")
	device, sensor, err := decodeSensor(r)
	if err != nil {
		writeResponse(w, err, nil, http.StatusBadRequest)
		return
	}

	result, err := s.Service.Create(device, sensor)
	if err != nil {
		writeResponse(w, err, result, http.StatusBadRequest)
		return
	}
	writeResponse(w, err, result, http.StatusOK)
	logger.GetInstance().InfoLogger.Println("Sensor is created")
}
func (s *sensorHandler) Delete(w http.ResponseWriter, r *http.Request) {
	logger.GetInstance().InfoLogger.Println("Sensor is being deleted...")
	params := mux.Vars(r)
	id := params["id"]
	if id == " " {
		err := errors.New("you are missing id parameter.")
		writeResponse(w, err, nil, http.StatusBadRequest)
		return
	}
	err := s.Service.Delete(id)
	if err != nil {
		writeResponse(w, err, nil, http.StatusBadRequest)
		return
	}
	writeResponse(w, err, nil, http.StatusNoContent)
	logger.GetInstance().InfoLogger.Println("Sensor has been deleted")
}
func (s *sensorHandler) Update(w http.ResponseWriter, r *http.Request) {
	logger.GetInstance().InfoLogger.Println("Sensor is about to be updated...")
	params := mux.Vars(r)
	id := params["id"]
	if id == " " {
		err := errors.New("you are missing id parameter.")
		writeResponse(w, err, nil, http.StatusBadRequest)
		return
	}
	device, sensor, err := decodeSensor(r)
	if err != nil {
		err := errors.New("error decoding body")
		writeResponse(w, err, nil, http.StatusBadRequest)
		return
	}
	result, err := s.Service.Update(id, device, sensor)
	if err != nil {
		writeResponse(w, err, result, http.StatusBadRequest)
		return
	}
	writeResponse(w, err, result, http.StatusOK)
	logger.GetInstance().InfoLogger.Println("Sensor has been updated")
}
