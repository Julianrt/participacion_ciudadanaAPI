package handlers

import (
	"encoding/json"
	"net/http"

	"../models"
)

//Login handler
func Login(w http.ResponseWriter, r *http.Request) {
	var usuarioRequest models.Usuario
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&usuarioRequest); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}
	usuario, err := models.LoginUsuario(usuarioRequest.Username, usuarioRequest.Pass)
	if err != nil {
		models.SendNotFound(w)
		return
	}
	if usuario.ID == 0 {
		models.SendNotFound(w)
		return
	}
	uuid := models.GetUuidByUserID(usuario.ID)
	if uuid == "" {
		uuid = models.CreateUuid(usuario.ID)
	}

	tokenResponse := models.GetTokenWithUuid(uuid)
	models.SendData(w, tokenResponse)
}
