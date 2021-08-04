package manejadores

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Farber98/WIMP/db"
	"github.com/Farber98/WIMP/helpers"
	"github.com/Farber98/WIMP/models"
	"github.com/Farber98/WIMP/structs"
)

/* Crea un usuario nuevo. */
func CrearUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario structs.Usuarios

	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		http.Error(w, " Error: "+err.Error(), http.StatusBadRequest)
		return
	}

	/* Sanitizamos */
	usuario.Usuario = strings.TrimSpace(usuario.Usuario)
	usuario.Email = strings.TrimSpace(usuario.Email)
	usuario.Password = strings.TrimSpace(usuario.Password)

	if len(usuario.Email) == 0 {
		http.Error(w, "Email requerido.", http.StatusBadRequest)
		return
	}

	if len(usuario.Usuario) == 0 {
		http.Error(w, "Nombre de usuario requerido.", http.StatusBadRequest)
		return
	}

	if len(usuario.Usuario) < 4 {
		http.Error(w, "Debe especificar un nombre de usuario de al menos 4 caracteres.", http.StatusBadRequest)
		return
	}

	if len(usuario.Password) < 8 {
		http.Error(w, "Debe especificar una contraseña de al menos 8 caracteres.", http.StatusBadRequest)
		return
	}

	_, duplicateEmail, _ := db.EmailDuplicado(usuario.Email)
	if duplicateEmail {
		http.Error(w, "Ya existe un usuario registrado con ese mail.", http.StatusBadRequest)
		return
	}

	_, duplicateUsername, _ := db.UsuarioDuplicado(usuario.Usuario)
	if duplicateUsername {
		http.Error(w, "Ya existe un usuario registrado con ese nombre de usuario.", http.StatusBadRequest)
		return
	}

	/* Encriptamos la PW. */
	usuario.Password, _ = helpers.EncriptarPassword(usuario.Password)
	_, status, err := db.CrearUsuario(usuario)
	if err != nil {
		http.Error(w, "Error al realizar el registro de usuario."+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "Fallo al insertar el registro.", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

/* Inicia la sesion de un usuario en la BD. Genera el TOKEN. */
func IniciarSesion(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var u structs.Usuarios

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Nombre de usuario y/o contraseña invalidos."+err.Error(), http.StatusBadRequest)
		return
	}

	if len(u.Usuario) == 0 {
		http.Error(w, "El usuario es requerido.", http.StatusBadRequest)
		return
	}

	if len(u.Password) == 0 {
		http.Error(w, "La contraseña es requerida.", http.StatusBadRequest)
		return
	}

	document, exists := db.IniciarSesion(u.Usuario, u.Password)
	if !exists {
		http.Error(w, "Nombre de usuario y/o contraseña erroneos.", http.StatusBadRequest)
		return
	}

	jwtKey, err := helpers.GenerarJwt(document)
	if err != nil {
		http.Error(w, "Error en la generacion de Token."+err.Error(), http.StatusBadRequest)
		return
	}

	response := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
