package handlers

import (
	"fmt"
	"net/http"
)

//Status method
func Status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "true")
}
