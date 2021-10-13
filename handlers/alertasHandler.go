package manejadores

import (
	"encoding/json"
	"net/http"

	"github.com/Farber98/WIMP/db"
)

/* Devuelve todos los switches en formato JSON. */
func VerAlertas(w http.ResponseWriter, r *http.Request) {

	results, status := db.DameAlertas()
	if !status {
		http.Error(w, "Error al traer las alertas. ", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(results)

}

/* Permite borrar un switch siempre y cuando no tenga hijos asociados. */
/* func ConfirmarAlerta(w http.ResponseWriter, r *http.Request) {
	IdAlerta := r.URL.Query().Get("idAlerta")
	if len(IdAlerta) < 1 {
		http.Error(w, "Debe enviar el parametro ID del switch.", http.StatusBadRequest)
		return
	}

	objID, _ := primitive.ObjectIDFromHex(IdAlerta)
	_, deleteID, _ := db.ExisteIdAlertas(objID)
	if !deleteID {
		http.Error(w, "No existe una alerta con ese ID.", http.StatusBadRequest)
		return
	}

	err := db.ConfirmarAlerta(IdAlerta)
	if err != nil {
		http.Error(w, "Error al borrar el switch. "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
} */
