package routers

import (
	"github.com/Julianrt/participacion_ciudadana/handlers"
	"github.com/gorilla/mux"
)

func StatusEndpoints(mux *mux.Router) {
	mux.HandleFunc("/api/v1/status/", handlers.GetAllStatus).Methods("GET")
	mux.HandleFunc("/api/v1/status/{id:[0-9]+}", handlers.GetStatus).Methods("GET")
	mux.HandleFunc("/api/v1/status/", handlers.CreateStatus).Methods("POST")
	mux.HandleFunc("/api/v1/status/{id:[0-9]+}", handlers.UpdateStatus).Methods("PUT")
	mux.HandleFunc("/api/v1/status/{id:[0-9]+}", handlers.DeleteStatus).Methods("DELETE")
}

func CiudadesEndpoints(mux *mux.Router) {
	mux.HandleFunc("/api/v1/ciudades/", handlers.GetCiudades).Methods("GET")
	mux.HandleFunc("/api/v1/ciudades/{id:[0-9]+}", handlers.GetCiudad).Methods("GET")
	mux.HandleFunc("/api/v1/ciudades/", handlers.CreateCiudad).Methods("POST")
	mux.HandleFunc("/api/v1/ciudades/{id:[0-9]+}", handlers.UpdateCiudad).Methods("PUT")
	mux.HandleFunc("/api/v1/ciudades/{id:[0-9]+}", handlers.DeleteCiudad).Methods("DELETE")
}

func ColoniasEndpoints(mux *mux.Router) {
	mux.HandleFunc("/api/v1/colonias/", handlers.GetColonias).Methods("GET")
	mux.HandleFunc("/api/v1/colonias/{id:[0-9]+}", handlers.GetColonia).Methods("GET")
	mux.HandleFunc("/api/v1/colonias/", handlers.CreateColonia).Methods("POST")
	mux.HandleFunc("/api/v1/colonias/{id:[0-9]+}", handlers.UpdateColonia).Methods("PUT")
	mux.HandleFunc("/api/v1/colonias/{id:[0-9]+}", handlers.DeleteColonia).Methods("DELETE")
}

func DependenciasEndpoinst(mux *mux.Router) {
	mux.HandleFunc("/api/v1/dependencias/", handlers.GetDependencias).Methods("GET")
	mux.HandleFunc("/api/v1/dependencias/{id:[0-9]+}", handlers.GetDependencia).Methods("GET")
	mux.HandleFunc("/api/v1/dependencias/", handlers.CreateDependencia).Methods("POST")
	mux.HandleFunc("/api/v1/dependencias/{id:[0-9]+}", handlers.UpdateDependencia).Methods("PUT")
	mux.HandleFunc("/api/v1/dependencias/{id:[0-9]+}", handlers.DeleteDependencia).Methods("DELETE")
}

func UsuariosEndpoints(mux *mux.Router) {
	mux.HandleFunc("/api/v1/usuarios/", handlers.GetUsuarios).Methods("GET")
	mux.HandleFunc("/api/v1/usuarios/{id:[0-9]+}", handlers.GetUsuario).Methods("GET")
	mux.HandleFunc("/api/v1/usuarios/", handlers.CreateUsuario).Methods("POST")
	mux.HandleFunc("/api/v1/usuarios/{id:[0-9]+}", handlers.UpdateUsuario).Methods("PUT")
	mux.HandleFunc("/api/v1/usuarios/{id:[0-9]+}", handlers.DeleteUsuario).Methods("DELETE")
}

func ProblemasEndpoints(mux *mux.Router) {
	mux.HandleFunc("/api/v1/problemas/", handlers.GetProblemas).Methods("GET")
	mux.HandleFunc("/api/v1/problemas/{id:[0-9]+}", handlers.GetProblema).Methods("GET")
	mux.HandleFunc("/api/v1/problemas/", handlers.CreateProblema).Methods("POST")
	mux.HandleFunc("/api/v1/problemas/{id:[0-9]+}", handlers.UpdateProblema).Methods("PUT")
	mux.HandleFunc("/api/v1/problemas/{id:[0-9]+}", handlers.DeleteProblema).Methods("DELETE")
}

func VotosEndpoints(mux *mux.Router) {
	mux.HandleFunc("/api/v1/votos/", handlers.GetVotos).Methods("GET")
	mux.HandleFunc("/api/v1/votos/{id:[0-9]+}", handlers.GetVoto).Methods("GET")
	mux.HandleFunc("/api/v1/votos/", handlers.CreateVoto).Methods("POST")
	mux.HandleFunc("/api/v1/votos/{id:[0-9]+}", handlers.UpdateVoto).Methods("PUT")
	mux.HandleFunc("/api/v1/votos/{id:[0-9]+}", handlers.DeleteVoto).Methods("DELETE")
}
