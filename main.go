package main

import (
	"log"
	"net/http"

	"github.com/Julianrt/participacion_ciudadana/routers"
	"github.com/gorilla/mux"
)

func main() {

	mux := mux.NewRouter()

	endpoints(mux)

	log.Println("El servidor está escuchando por el puerto :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func endpoints(mux *mux.Router) {
	routers.StatusEndpoints(mux)
	routers.CiudadesEndpoints(mux)
	routers.ColoniasEndpoints(mux)
	routers.DependenciasEndpoinst(mux)
	routers.UsuariosEndpoints(mux)
	routers.ProblemasEndpoints(mux)
	routers.VotosEndpoints(mux)
	routers.LoginEndpoints(mux)
}
