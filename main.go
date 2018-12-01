package main

import (
	"log"
	"net/http"

	"./routers"
	"github.com/gorilla/mux"
)

func main() {

	mux := mux.NewRouter()

	routers.Endpoints(mux)

	log.Println("El servidor está escuchando por el puerto :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
