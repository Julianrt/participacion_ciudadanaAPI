package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"../models"
	"github.com/gorilla/mux"
)

//GetUsuarios method
func GetUsuarios(w http.ResponseWriter, r *http.Request) {
	models.SendData(w, models.GetUsuarios())
}

//GetUsuario method
func GetUsuario(w http.ResponseWriter, r *http.Request) {
	if usuario, err := getUsuarioByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		if usuario.ID == 0 {
			models.SendNotFound(w)
			return
		}
		models.SendData(w, usuario)
	}
}

//CreateUsuario method
func CreateUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&usuario); err != nil {
		log.Println(err)
		models.SendUnprocessableEntity(w)
	} else {
		if err := usuario.Save(); err != nil {
			models.SendUnprocessableEntity(w)
		}
		models.SendData(w, usuario)
	}
}

//UpdateUsuario method
func UpdateUsuario(w http.ResponseWriter, r *http.Request) {
	usuario, err := getUsuarioByRequest(r)
	if err != nil {
		models.SendNotFound(w)
		return
	}
	var usuarioResponse models.Usuario
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&usuarioResponse); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}
	usuarioResponse.ID = usuario.ID
	usuarioResponse.Save()
	models.SendData(w, usuarioResponse)
}

//DeleteUsuario method
func DeleteUsuario(w http.ResponseWriter, r *http.Request) {
	if usuario, err := getUsuarioByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		usuario.Delete()
		models.SendNoContent(w)
	}
}

func getUsuarioByRequest(r *http.Request) (*models.Usuario, error) {
	vars := mux.Vars(r)
	usuarioID, _ := strconv.Atoi(vars["id"])
	usuario, err := models.GetUsuario(usuarioID)
	if err != nil {
		return usuario, err
	}
	return usuario, err
}
