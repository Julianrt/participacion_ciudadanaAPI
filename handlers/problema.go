package handlers

import (
	"fmt"
	"net/http"
)

//GetProblemas function
func GetProblemas(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PROBLEMAS")
}

//GetProblema function
func GetProblema(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PROBLEMA")
}

//CreateProblema function
func CreateProblema(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CREATE PROBLEMA")
}

//UpdateProblema function
func UpdateProblema(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UPDATE PROBLEMA")
}

//DeleteProblema function
func DeleteProblema(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DELETE PROBLEMA")
}
