package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"../models"
)

//GetProblemas function
func GetProblemas(w http.ResponseWriter, r *http.Request) {
	models.SendData(w, models.GetProblemas())
}

//GetProblema function
func GetProblema(w http.ResponseWriter, r *http.Request) {
	if problema, err := getProblemaByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		if problema.ID == 0 {
			models.SendNotFound(w)
			return
		}
		models.SendData(w, problema)
	}
}

//CreateProblema function
func CreateProblema(w http.ResponseWriter, r *http.Request) {
	var problema models.Problema
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&problema); err != nil {
		models.SendUnprocessableEntity(w)
	} else {
		problema.Save()
		models.SendData(w, problema)
	}
}

//UpdateProblema function
func UpdateProblema(w http.ResponseWriter, r *http.Request) {
	problema, err := getProblemaByRequest(r)
	if err != nil {
		models.SendNotFound(w)
		return
	}
	var problemaResponse models.Problema
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&problemaResponse); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}
	problemaResponse.ID = problema.ID
	problemaResponse.Save()
	models.SendData(w, problemaResponse)
}

//DeleteProblema function
func DeleteProblema(w http.ResponseWriter, r *http.Request) {
	if problema, err := getProblemaByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		problema.Delete()
		models.SendNoContent(w)
	}
}

func getProblemaByRequest(r *http.Request) (*models.Problema, error) {
	vars := mux.Vars(r)
	problemaID, _ := strconv.Atoi(vars["id"])

	problema, err := models.GetProblema(problemaID)
	if err != nil {
		return problema, err
	}
	return problema, err
}
