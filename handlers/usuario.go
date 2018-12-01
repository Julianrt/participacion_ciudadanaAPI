package handlers

import (
	"strconv"
	"net/http"
	"encoding/json"

	"../models"
	"github.com/gorilla/mux"
)

//GetUsuarios method
func GetUsuarios(w http.ResponseWriter, r *http.Request) {
	if len(models.GetUsuarios()) == 0 {
		models.SendNotFound(w)
		return
	}
	models.SendData(w, models.GetUsuarios())
}

//GetUsuario method
func GetUsuario(w http.ResponseWriter, r *http.Request) {
	if usuario, err := getUsuarioByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		models.SendData(w, usuario)
	}
}

//CreateUsuario method
func CreateUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&usuario); err != nil {
		models.SendUnprocessableEntity(w)
		return
	} 

	if err := usuario.Valid(); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}

	usuario.SetPassword(usuario.Pass)
	if err := usuario.Save(); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}

	models.SendData(w, usuario)
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
	usuario, _ := models.GetUsuarioByID(usuarioID)
	
	if usuario.ID == 0 {
		return usuario, models.ErrorUsuarioNoEncontrado
	}
	return usuario, nil
}
