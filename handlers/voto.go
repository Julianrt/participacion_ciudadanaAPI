package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Julianrt/participacion_ciudadana/models"
	"github.com/gorilla/mux"
)

//GetVotos function
func GetVotos(w http.ResponseWriter, r *http.Request) {
	models.SendData(w, models.GetVotos())
}

//GetVoto function
func GetVoto(w http.ResponseWriter, r *http.Request) {
	if voto, err := getVotoByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		if emptyResult(voto.ID) {
			models.SendNotFound(w)
			return
		}
		models.SendData(w, voto)
	}
}

//CreateVoto function
func CreateVoto(w http.ResponseWriter, r *http.Request) {
	var voto *models.Voto
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&voto); err != nil {
		models.SendUnprocessableEntity(w)
	} else {
		voto.Save()
		models.SendData(w, voto)
	}
}

//UpdateVoto function
func UpdateVoto(w http.ResponseWriter, r *http.Request) {
	voto, err := getVotoByRequest(r)
	if err != nil {
		models.SendNotFound(w)
		return
	}
	var votoResponse models.Voto
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&votoResponse); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}
	votoResponse.ID = voto.ID
	votoResponse.Save()
	models.SendData(w, votoResponse)
}

//DeleteVoto function
func DeleteVoto(w http.ResponseWriter, r *http.Request) {
	if voto, err := getVotoByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		if emptyResult(voto.ID) {
			models.SendNotFound(w)
			return
		}
		voto.Delete()
		models.SendNoContent(w)
	}
}

func getVotoByRequest(r *http.Request) (*models.Voto, error) {
	vars := mux.Vars(r)
	votoID, _ := strconv.Atoi(vars["id"])
	voto, err := models.GetVoto(votoID)
	return voto, err
}

func emptyResult(id int64) bool {
	if id == 0 {
		return true
	}
	return false
}
