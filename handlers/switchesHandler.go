package manejadores

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Farber98/WIMP/db"
	"github.com/Farber98/WIMP/structs"
)

func UbicarSwitch(w http.ResponseWriter, r *http.Request) {
	var s structs.Switches

	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, "Datos incorrectos. "+err.Error(), http.StatusBadRequest)
		return
	}

	//Sanitizamos
	s.Mac = strings.TrimSpace(s.Mac)
	if s.Mac == "" {
		http.Error(w, "Debe especificar un switch para ubicar.", http.StatusBadRequest)
		return
	}

	if s.Lat == 0 || s.Lng == 0 {
		http.Error(w, "Debe especificar latitud y longitud.", http.StatusBadRequest)
		return
	}

	errr := db.UbicarSwitch(s)
	if errr != nil {
		http.Error(w, "Error al intentar ubicar el switch."+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

/* Permite crear un nuevo switch. */
/* func CrearSwitch(w http.ResponseWriter, r *http.Request) {
	var s structs.Switches
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, " Error: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Sanitizamos
	s.Modelo = strings.TrimSpace(s.Modelo)
	s.Nombre = strings.TrimSpace(s.Nombre)

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

	//000... es el valor vacio de _pid
	if s.IdPadre.Hex() != "000000000000000000000000" {
		_, parentID, _ := db.ExisteIdSwitches(s.IdPadre, false)
		if !parentID {
			http.Error(w, "No existe un padre con ese ID.", http.StatusBadRequest)
			return
		}

		//No permitir que se creen hijos con padre desactivado.
		_, active, _ := db.EstaActivo(s.IdPadre)
		if !active {
			http.Error(w, "El padre esta desactivado.", http.StatusBadRequest)
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
} */

/* 	Permite modificar nombre, modelo y padre de un Switch existente. Actualiza marca de tiempo. */
/* func ModificarSwitch(w http.ResponseWriter, r *http.Request) {
	var s structs.Switches

	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, "Datos incorrectos."+err.Error(), http.StatusBadRequest)
		return
	}

	//Sanitizamos
	s.Modelo = strings.TrimSpace(s.Modelo)
	s.Nombre = strings.TrimSpace(s.Nombre)

	if len(s.Modelo) == 0 {
		http.Error(w, "Modelo requerido.", http.StatusBadRequest)
		return
	}

	if len(s.Nombre) == 0 {
		http.Error(w, "Nombre requerido.", http.StatusBadRequest)
		return
	}

	if s.Lat == 0 || s.Lng == 0 {
		http.Error(w, "Debe especificar latitud y longitud.", http.StatusBadRequest)
		return
	}

	if s.IdSwitch.Hex() == "000000000000000000000000" {
		http.Error(w, "No se especifico un ID de switch a modificar.", http.StatusBadRequest)
		return
	}

	_, modifyID, _ := db.ExisteIdSwitches(s.IdSwitch, false)
	if !modifyID {
		http.Error(w, "No existe un switch con ese ID.", http.StatusBadRequest)
		return
	}

	switchModificado := make(map[string]interface{})

	if s.IdPadre.Hex() != "000000000000000000000000" {
		_, parentID, _ := db.ExisteIdSwitches(s.IdPadre, false)
		if !parentID {
			http.Error(w, "No existe un padre con ese ID.", http.StatusBadRequest)
			return
		}
		switchModificado["_pid"] = s.IdPadre
	}

	if len(s.Nombre) > 0 {
		switchBD, duplicateSwitch, _ := db.NombreDuplicado(s.Nombre)
		if duplicateSwitch && switchBD.IdSwitch.Hex() != s.IdSwitch.Hex() {
			http.Error(w, "Ya existe un switch con ese nombre.", http.StatusBadRequest)
			return
		}
		switchModificado["nombre"] = s.Nombre
	}

	if len(s.Modelo) > 0 {
		switchModificado["modelo"] = s.Modelo
	}

	switchModificado["lat"] = s.Lat
	switchModificado["lng"] = s.Lng
	switchModificado["fecha"] = time.Now()

	status, err := db.ModificarSwitch(s.IdSwitch, switchModificado)
	if err != nil {
		http.Error(w, "Error al intentar modificar el switch."+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se pudo modificar el registro del usuario. ", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
} */

/* Permite borrar un switch siempre y cuando no tenga hijos asociados. */
/* func BorrarSwitch(w http.ResponseWriter, r *http.Request) {
	IdSwitch := r.URL.Query().Get("idSwitch")
	if len(IdSwitch) < 1 {
		http.Error(w, "Debe enviar el parametro ID del switch.", http.StatusBadRequest)
		return
	}

	objID, _ := primitive.ObjectIDFromHex(IdSwitch)
	_, deleteID, _ := db.ExisteIdSwitches(objID, false)
	if !deleteID {
		http.Error(w, "No existe un switch con ese ID.", http.StatusBadRequest)
		return
	}

	_, parentID, _ := db.ExisteIdSwitches(objID, true)
	if parentID {
		http.Error(w, "No se puede borrar un switch que tiene hijos asociados.", http.StatusBadRequest)
		return
	}

	err := db.BorrarSwitch(IdSwitch)
	if err != nil {
		http.Error(w, "Error al borrar el switch. "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
} */

/* Permite activar un switch (estado = true) siempre y cuando no este activo ya. */
/* func ActivarSwitch(w http.ResponseWriter, r *http.Request) {
	var s structs.Switches
	registry := make(map[string]interface{})

	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, "Datos incorrectos."+err.Error(), http.StatusBadRequest)
		return
	}

	//Check if switch ID is specified
	if s.IdSwitch.Hex() == "000000000000000000000000" {
		http.Error(w, "No se especifico un ID de switch a activar.", http.StatusBadRequest)
		return
	}

	//Check if ID exists
	_, exists, _ := db.ExisteIdSwitches(s.IdSwitch, false)
	if !exists {
		http.Error(w, "No existe un switch con ese ID.", http.StatusBadRequest)
		return
	}

	//Check if switch is active
	_, active, _ := db.EstaActivo(s.IdSwitch)
	if active {
		http.Error(w, "El switch ya esta activo.", http.StatusBadRequest)
		return
	}

	registry["estado"] = true

	status, err := db.ModificarSwitch(s.IdSwitch, registry)
	if err != nil {
		http.Error(w, "Error al intentar modificar el switch."+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se pudo modificar el registro del switch. ", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

} */

/* Permite desactivar un switch (estado = false) siempre y cuando no este desactivado ya y no tenga hijos activos */
/* func DesactivarSwitch(w http.ResponseWriter, r *http.Request) {
	var s structs.Switches
	registry := make(map[string]interface{})

	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, "Datos incorrectos."+err.Error(), http.StatusBadRequest)
		return
	}

	//Check if switch ID is specified
	if s.IdSwitch.Hex() == "000000000000000000000000" {
		http.Error(w, "No se especifico un ID de switch a desactivar.", http.StatusBadRequest)
		return
	}

	//Check if ID exists
	_, exists, _ := db.ExisteIdSwitches(s.IdSwitch, false)
	if !exists {
		http.Error(w, "No existe un switch con ese ID.", http.StatusBadRequest)
		return
	}

	//Check if switch is active
	_, active, _ := db.EstaActivo(s.IdSwitch)
	if !active {
		http.Error(w, "El switch ya esta desactivado.", http.StatusBadRequest)
		return
	}

	//Check parent and active
	_, childActive, _ := db.HijoActivo(s.IdSwitch)
	if childActive {
		http.Error(w, "No se puede desactivar un switch que tiene hijos activos.", http.StatusBadRequest)
		return
	}

	registry["estado"] = false

	status, err := db.ModificarSwitch(s.IdSwitch, registry)
	if err != nil {
		http.Error(w, "Error al intentar modificar el switch."+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se pudo modificar el registro del usuario. ", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
} */
