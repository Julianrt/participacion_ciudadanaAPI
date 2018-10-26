package main

import (
	"log"
	"net/http"

	"github.com/Julianrt/participacion_ciudadana/handlers"
	"github.com/gorilla/mux"
)

func main() {

	mux := mux.NewRouter()

	endpoints(mux)

	log.Println("El servidor está escuchando por el puerto :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func endpoints(mux *mux.Router) {
	statusEndpoint(mux)
	ciudadesEndpoints(mux)
	coloniasEndpoints(mux)
	dependenciasEndpoinst(mux)
	usuariosEndpoints(mux)
}

func statusEndpoint(mux *mux.Router) {
	mux.HandleFunc("/api/v1/status/", handlers.Status).Methods("GET")
}

func ciudadesEndpoints(mux *mux.Router) {
	mux.HandleFunc("/api/v1/ciudades/", handlers.GetCiudades).Methods("GET")
	mux.HandleFunc("/api/v1/ciudades/{id:[0-9]+}", handlers.GetCiudad).Methods("GET")
	mux.HandleFunc("/api/v1/ciudades/", handlers.CreateCiudad).Methods("POST")
	mux.HandleFunc("/api/v1/ciudades/{id:[0-9]+}", handlers.UpdateCiudad).Methods("PUT")
	mux.HandleFunc("/api/v1/ciudades/{id:[0-9]+}", handlers.DeleteCiudad).Methods("DELETE")
}

func coloniasEndpoints(mux *mux.Router) {
	mux.HandleFunc("/api/v1/colonias/", handlers.GetColonias).Methods("GET")
	mux.HandleFunc("/api/v1/colonias/{id:[0-9]+}", handlers.GetColonia).Methods("GET")
	mux.HandleFunc("/api/v1/colonias/", handlers.CreateColonia).Methods("POST")
	mux.HandleFunc("/api/v1/colonias/{id:[0-9]+}", handlers.UpdateColonia).Methods("PUT")
	mux.HandleFunc("/api/v1/colonias/{id:[0-9]+}", handlers.DeleteColonia).Methods("DELETE")
}

func dependenciasEndpoinst(mux *mux.Router) {
	mux.HandleFunc("/api/v1/dependencias/", handlers.GetDependencias).Methods("GET")
	mux.HandleFunc("/api/v1/dependencias/{id:[0-9]+}", handlers.GetDependencia).Methods("GET")
	mux.HandleFunc("/api/v1/dependencias/", handlers.CreateDependencia).Methods("POST")
	mux.HandleFunc("/api/v1/dependencias/{id:[0-9]+}", handlers.UpdateDependencia).Methods("PUT")
	mux.HandleFunc("/api/v1/dependencias/{id:[0-9]+}", handlers.DeleteDependencia).Methods("DELETE")
}

func usuariosEndpoints(mux *mux.Router) {
	mux.HandleFunc("/api/v1/usuarios/", handlers.GetUsuarios).Methods("GET")
	mux.HandleFunc("/api/v1/usuarios/{id:[0-9]+}", handlers.GetUsuario).Methods("GET")
	mux.HandleFunc("/api/v1/usuarios/", handlers.CreateUsuario).Methods("POST")
	mux.HandleFunc("/api/v1/usuarios/{id:[0-9]+}", handlers.UpdateUsuario).Methods("PUT")
	mux.HandleFunc("/api/v1/usuarios/{id:[0-9]+}", handlers.DeleteUsuario).Methods("DELETE")
}
