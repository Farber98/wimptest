package manejadores

import (
	"encoding/json"
	"net/http"

	"github.com/Farber98/WIMP/db"
)

func VerTopologia(w http.ResponseWriter, r *http.Request) {

	results, status := db.DameTopologia()
	if !status {
		http.Error(w, "error al decodificar el JSON de la peticion: ", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(results)

}
