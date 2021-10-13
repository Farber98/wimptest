package manejadores

import (
	"encoding/json"
	"net/http"

	"github.com/Farber98/WIMP/db"
	"github.com/Farber98/WIMP/structs"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ListarSrcMacMayorEmision(w http.ResponseWriter, r *http.Request) {

	results := db.DameSrcMacMayorEmision()

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(results)
}

func ListarSrcIpMayorEmision(w http.ResponseWriter, r *http.Request) {

	results := db.DameSrcIpMayorEmision()

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

/* Metodo que lista la emision y las ip destino dada una $srcMac */
func ListarSrcMacDetalle(w http.ResponseWriter, r *http.Request) {
	var mac structs.Paquetes
	var results []primitive.M
	err := json.NewDecoder(r.Body).Decode(&mac)
	if err != nil {
		http.Error(w, "Datos incorrectos."+err.Error(), http.StatusBadRequest)
		return
	}

	results = db.DameSrcMacEmision(mac.SrcMac)
	//results = append(results, db.DameSrcMacProtoIp(mac.SrcMac)...)
	//results = append(results, db.DameSrcMacProtoTp(mac.SrcMac)...)
	//results = append(results, db.DameSrcMacProtoApp(mac.SrcMac)...)
	//results = append(results, db.DameDstPort(mac.SrcMac)...)
	results = append(results, db.DameDstIp(mac.SrcMac)...)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(results)
}
