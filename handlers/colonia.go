package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Julianrt/participacion_ciudadana/models"
	"github.com/gorilla/mux"
)

//GetColonias method
func GetColonias(w http.ResponseWriter, r *http.Request) {
	models.SendData(w, models.GetColonias())
}

//GetColonia method
func GetColonia(w http.ResponseWriter, r *http.Request) {
	if colonia, err := getColoniaByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		if colonia.ID == 0 {
			models.SendNotFound(w)
			return
		}
		models.SendData(w, colonia)
	}
}

//CreateColonia method
func CreateColonia(w http.ResponseWriter, r *http.Request) {
	var colonia models.Colonia
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&colonia); err != nil {
		models.SendUnprocessableEntity(w)
	} else {
		colonia.Save()
		models.SendData(w, colonia)
	}
}

//UpdateColonia method
func UpdateColonia(w http.ResponseWriter, r *http.Request) {
	colonia, err := getColoniaByRequest(r)
	if err != nil {
		models.SendNotFound(w)
		return
	}
	var coloniaResponse models.Colonia
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&coloniaResponse); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}
	coloniaResponse.ID = colonia.ID
	coloniaResponse.Save()
	models.SendData(w, coloniaResponse)
}

//DeleteColonia method
func DeleteColonia(w http.ResponseWriter, r *http.Request) {
	if colonia, err := getColoniaByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		colonia.Delete()
		models.SendNoContent(w)
	}

}

func getColoniaByRequest(r *http.Request) (*models.Colonia, error) {
	vars := mux.Vars(r)
	coloniaID, _ := strconv.Atoi(vars["id"])

	colonia, err := models.GetColonia(coloniaID)
	if err != nil {
		return colonia, err
	}
	return colonia, nil
}
