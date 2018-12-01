package handlers

import (
	"encoding/json"
	"log"
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
	token := models.GetToken(usuario.ID)
	if token == "" {
		token = models.CreateToken(usuario.ID)
	}

	if tokenResponse, err := models.GetTokenWithUser(token); err != nil {
		log.Println(err.Error()+" 2 ")
	} else {
		models.SendData(w, tokenResponse)
	}
}
