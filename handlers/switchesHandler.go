package manejadores

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Farber98/WIMP/db"
	"github.com/Farber98/WIMP/structs"
)

/* Handler para crear un nuevo switch. */
func CrearSwitch(w http.ResponseWriter, r *http.Request) {
	var s structs.Switches
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, " Error: "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(s.Nombre) == 0 {
		http.Error(w, "Nombre requerido.", http.StatusBadRequest)
		return
	}

	if len(s.Modelo) == 0 {
		http.Error(w, "Modelo requerido.", http.StatusBadRequest)
		return
	}

	if s.Lat == 0 || s.Lng == 0 {
		http.Error(w, "Debe especificar latitud y longitud.", http.StatusBadRequest)
		return
	}

	s.Fecha = time.Now()
	s.Estado = true

	_, duplicado, _ := db.NombreDuplicado(s.Nombre)
	if duplicado {
		http.Error(w, "Ya existe un switch con ese nombre.", http.StatusBadRequest)
		return
	}

	/* 000... es el valor vacio de _pid */
	if s.IdPadre.Hex() != "000000000000000000000000" {
		_, parentID, _ := db.ExisteId(s.IdPadre)
		if !parentID {
			http.Error(w, "No existe un padre con ese ID.", http.StatusBadRequest)
			return
		}
	}
	_, status, err := db.CrearSwitch(s)
	if err != nil {
		http.Error(w, "Error al realizar el registro del switch."+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "Fallo al insertar el switch.", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

/* /switches handler that shows switches topology */
func VerTopologia(w http.ResponseWriter, r *http.Request) {

	results, status := db.DameTopologia()
	if !status {
		http.Error(w, "Error al traer la topologia. ", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(results)

}
