package manejadores

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Farber98/WIMP/db"
	"github.com/Farber98/WIMP/structs"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RankingSrcMacTransmision(w http.ResponseWriter, r *http.Request) {

	results := db.RankingSrcMacTransmision()

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(results)
}

func RankingProtoApp(w http.ResponseWriter, r *http.Request) {

	results, status := db.RankingProtoApp()
	if !status {
		http.Error(w, "Error al traer el ranking. ", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(results)
}

func RankingProtoTransporte(w http.ResponseWriter, r *http.Request) {

	results := db.RankingProtoTransporte()

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(results)
}

func RankingProtoRed(w http.ResponseWriter, r *http.Request) {

	results := db.RankingProtoRed()

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(results)
}

func DetalleSrcMac(w http.ResponseWriter, r *http.Request) {
	var s structs.Switches
	var results []primitive.M
	s.Mac = r.URL.Query().Get("mac")
	s.Mac = strings.TrimSpace(s.Mac)
	if s.Mac == "" {
		http.Error(w, "Debe enviar la direccion MAC para armar el detalle.", http.StatusBadRequest)
		return
	}

	/* Sanitizamos */
	s.Mac = strings.TrimSpace(s.Mac)
	if s.Mac == "" {
		http.Error(w, "Debe especificar un dispositivo.", http.StatusBadRequest)
		return
	}

	//results = append(results, db.DameSrcMacProtoIp(mac.SrcMac)...)
	//results = append(results, db.DameSrcMacProtoTp(mac.SrcMac)...)
	//results = append(results, db.DameSrcMacProtoApp(mac.SrcMac)...)
	//results = append(results, db.DameDstPort(mac.SrcMac)...)
	results = db.DetalleSrcMacEmision(s)
	results = append(results, db.DetalleSrcMacDstIp(s)...)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(results)
}
