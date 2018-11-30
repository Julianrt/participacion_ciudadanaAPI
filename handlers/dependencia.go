package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../models"
	"github.com/gorilla/mux"
)

//GetDependencias method
func GetDependencias(w http.ResponseWriter, r *http.Request) {
	models.SendData(w, models.GetDependencias())
}

//GetDependencia method
func GetDependencia(w http.ResponseWriter, r *http.Request) {
	if dependencia, err := getDependenciaByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		if dependencia.ID == 0 {
			models.SendNotFound(w)
			return
		}
		models.SendData(w, dependencia)
	}
}

//CreateDependencia method
func CreateDependencia(w http.ResponseWriter, r *http.Request) {
	var dependencia models.Dependencia
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dependencia); err != nil {
		models.SendUnprocessableEntity(w)
	} else {
		dependencia.Save()
		models.SendData(w, dependencia)
	}
}

//UpdateDependencia method
func UpdateDependencia(w http.ResponseWriter, r *http.Request) {
	dependencia, err := getDependenciaByRequest(r)
	if err != nil {
		models.SendNotFound(w)
		return
	}
	var dependenciaResponse models.Dependencia
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dependenciaResponse); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}
	dependenciaResponse.ID = dependencia.ID
	dependenciaResponse.Save()
	models.SendData(w, dependenciaResponse)
}

//DeleteDependencia method
func DeleteDependencia(w http.ResponseWriter, r *http.Request) {
	if dependencia, err := getDependenciaByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		dependencia.Delete()
		models.SendNoContent(w)
	}
}

func getDependenciaByRequest(r *http.Request) (*models.Dependencia, error) {
	vars := mux.Vars(r)
	dependenciaID, _ := strconv.Atoi(vars["id"])

	dependencia, err := models.GetDependencia(dependenciaID)
	if err != nil {
		return dependencia, err
	}
	return dependencia, err
}
