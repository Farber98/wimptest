package manejadores

import (
	"encoding/json"
	"net/http"

	"github.com/Farber98/WIMP/db"
)

func ListarAnomalias(w http.ResponseWriter, r *http.Request) {

	results, status := db.ListarAnomalias()
	if !status {
		http.Error(w, "Error al traer las anomalias. ", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(results)

}

func RankingAnomalias(w http.ResponseWriter, r *http.Request) {

	results := db.RankingAnomalias()

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(results)

}

/* func AnomaliasSrcMac(w http.ResponseWriter, r *http.Request) {
	var s structs.Switches
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, "error al decodificar el JSON de la peticion: "+err.Error(), http.StatusBadRequest)
		return
	}

	//Sanitizamos
	s.Mac = strings.TrimSpace(s.Mac)
	if s.Mac == "" {
		http.Error(w, "Debe especificar un dispositivo.", http.StatusBadRequest)
		return
	}

	results := db.AnomaliasSrcMac(s)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(results)
}
*/
