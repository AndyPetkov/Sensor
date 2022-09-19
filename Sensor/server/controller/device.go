package controller

import (
	"errors"
	"net/http"
	"server/global"
	"server/logger"
	"server/service"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type deviceHandler struct {
	Service service.Service
}
type Controller interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	ConfigureService(service service.Service)
}

func NewController(sensorType string) Controller {
	if sensorType == global.Device {
		return &deviceHandler{service.NewService(global.Device)}
	}
	return &sensorHandler{service.NewService(global.Sensor)}
}
func (d *deviceHandler) ConfigureService(service service.Service) {
	d.Service = service
}

func (d *deviceHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	logger.GetInstance().InfoLogger.Println("Getting all devices...")
	result, err := d.Service.GetAll()
	if err != nil {
		writeResponse(w, err, result, http.StatusBadRequest)
		return
	}
	writeResponse(w, err, result, http.StatusOK)
	logger.GetInstance().InfoLogger.Println("All devices are got")
}
func (d *deviceHandler) GetById(w http.ResponseWriter, r *http.Request) {
	logger.GetInstance().InfoLogger.Println("Getting device by id...")
	params := mux.Vars(r)
	id := params["id"]
	if id == " " {
		err := errors.New("you are missing id parameter.")
		writeResponse(w, err, nil, http.StatusBadRequest)
		return
	}
	result, err := d.Service.GetById(id)
	if err != nil {
		writeResponse(w, err, result, http.StatusBadRequest)
		return
	}
	writeResponse(w, err, result, http.StatusOK)
	logger.GetInstance().InfoLogger.Println("Device by id is got")
}
func (d *deviceHandler) Create(w http.ResponseWriter, r *http.Request) {
	logger.GetInstance().InfoLogger.Println("New device being created...")
	device, sensor, err := decodeDevice(r)
	if err != nil {
		writeResponse(w, err, nil, http.StatusBadRequest)
		return
	}
	result, err := d.Service.Create(device, sensor)
	if err != nil {
		writeResponse(w, err, result, http.StatusBadRequest)
		return
	}
	writeResponse(w, err, result, http.StatusOK)
	logger.GetInstance().InfoLogger.Println("Device is created")
}
func (d *deviceHandler) Delete(w http.ResponseWriter, r *http.Request) {
	logger.GetInstance().InfoLogger.Println("Device is being deleted...")
	params := mux.Vars(r)
	id := params["id"]
	if id == " " {
		err := errors.New("you are missing id parameter.")
		writeResponse(w, err, nil, http.StatusBadRequest)
		return
	}
	err := d.Service.Delete(id)
	if err != nil {
		writeResponse(w, err, nil, http.StatusBadRequest)
		return
	}
	writeResponse(w, err, nil, http.StatusNoContent)
	logger.GetInstance().InfoLogger.Println("Device has been deleted")
}
func (d *deviceHandler) Update(w http.ResponseWriter, r *http.Request) {
	logger.GetInstance().InfoLogger.Println("Device is about to be updated...")
	params := mux.Vars(r)
	id := params["id"]
	if id == " " {
		err := errors.New("you are missing id parameter.")
		writeResponse(w, err, nil, http.StatusBadRequest)
		return
	}
	device, sensor, err := decodeDevice(r)
	if err != nil {
		err := errors.New("error decoding body")
		writeResponse(w, err, nil, http.StatusBadRequest)
		return
	}
	result, err := d.Service.Update(id, device, sensor)
	if err != nil {
		writeResponse(w, err, result, http.StatusBadRequest)
		return
	}
	writeResponse(w, err, result, http.StatusOK)
	logger.GetInstance().InfoLogger.Println("Device has been updated")
}
