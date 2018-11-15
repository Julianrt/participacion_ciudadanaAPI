package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Julianrt/participacion_ciudadana/models"
)

//Login handler
func Login(w http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&usuario); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}
	u, err := models.GetUsuarioByUsername(usuario.Username, usuario.Pass)
	if err != nil {
		log.Println(err)
		return
	}
	if u.ID == 0 {
		log.Println("No existe el usuario")
		return
	}
	token := models.GetToken(u.ID)
	if token == "" {
		token = models.CreateToken(u.ID)
	}

	if tokenResponse, err := models.GetTokenWithUser(token); err != nil {
		log.Println(err)
	} else {
		models.SendData(w, tokenResponse)
	}
}
