package manejadores

import (
	"encoding/json"
	"net/http"

	"github.com/Farber98/WIMP/db"
)

func ListarSrcMacMayorEmision(w http.ResponseWriter, r *http.Request) {

	results := db.DameSrcMacMayorEmision()

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(results)
}

func ListarProtocolosAplicacionMayorEmision(w http.ResponseWriter, r *http.Request) {

	results := db.DameProtocolosAplicacionMayorEmision()

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(results)
}

func ListarProtocolosTransporteMayorEmision(w http.ResponseWriter, r *http.Request) {

	results := db.DameProtocolosTransporteMayorEmision()

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(results)
}

func ListarProtocolosRedMayorEmision(w http.ResponseWriter, r *http.Request) {

	results := db.DameProtocolosRedMayorEmision()

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(results)
}
