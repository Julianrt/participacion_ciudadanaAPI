package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Julianrt/participacion_ciudadana/models"
	"github.com/gorilla/mux"
)

//GetAllStatus function
func GetAllStatus(w http.ResponseWriter, r *http.Request) {
	models.SendData(w, models.GetAllStatus())
}

//GetStatus function
func GetStatus(w http.ResponseWriter, r *http.Request) {
	if status, err := getStatusByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		if status.ID == 0 {
			models.SendNotFound(w)
			return
		}
		models.SendData(w, status)
	}
}

//CreateStatus function
func CreateStatus(w http.ResponseWriter, r *http.Request) {
	var status models.Status
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&status); err != nil {
		models.SendUnprocessableEntity(w)
	} else {
		status.Save()
		models.SendData(w, status)
	}
}

//UpdateStatus function
func UpdateStatus(w http.ResponseWriter, r *http.Request) {
	status, err := getStatusByRequest(r)
	if err != nil {
		models.SendNotFound(w)
		return
	}
	var statusResponse models.Status
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&statusResponse); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}
	statusResponse.ID = status.ID
	statusResponse.Save()
	models.SendData(w, statusResponse)
}

//DeleteStatus function
func DeleteStatus(w http.ResponseWriter, r *http.Request) {
	if status, err := getStatusByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		status.Delete()
		models.SendNoContent(w)
	}
}

func getStatusByRequest(r *http.Request) (*models.Status, error) {
	vars := mux.Vars(r)
	statusID, _ := strconv.Atoi(vars["id"])
	status, err := models.GetStatus(statusID)
	if err != nil {
		return status, err
	}
	return status, err
}
