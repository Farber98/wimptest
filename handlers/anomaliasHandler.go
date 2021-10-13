package manejadores

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Farber98/WIMP/db"
	"github.com/Farber98/WIMP/structs"
)

func VerAnomalias(w http.ResponseWriter, r *http.Request) {

	results, status := db.DameAnomalias()
	if !status {
		http.Error(w, "Error al traer las anomalias. ", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(results)

}

func RankingAnomaliasPorMac(w http.ResponseWriter, r *http.Request) {

	results := db.RankingAnomaliasPorMac()

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(results)

}

func AnomaliasPorMac(w http.ResponseWriter, r *http.Request) {
	var s structs.Switches
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, "Datos incorrectos."+err.Error(), http.StatusBadRequest)
		return
	}

	//Sanitizamos
	s.Mac = strings.TrimSpace(s.Mac)
	if s.Mac == "" {
		http.Error(w, "Debe especificar un dispositivo.", http.StatusBadRequest)
		return
	}

	results := db.DameSrcMacAnomalias(s.Mac)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(results)
}
