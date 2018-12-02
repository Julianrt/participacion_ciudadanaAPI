package main

import (
	"log"
	"net/http"

	"./routers"
	"./config"
	"github.com/gorilla/mux"
)

func main() {

	mux := mux.NewRouter()

	routers.Endpoints(mux)

	log.Println("El servidor está escuchando por el puerto :", config.ServerPort())
	server := http.Server{
		Addr: 		config.URLServer(),
		Handler: 	mux,
	}
	log.Fatal(server.ListenAndServe())
}
