package manejadores

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Farber98/WIMP/db"
	"github.com/Farber98/WIMP/helpers"
	midl "github.com/Farber98/WIMP/middleware"
	"github.com/Farber98/WIMP/structs"
)

/* Crea un usuario nuevo, si tiene los permisos necesarios. */
func CrearUsuario(w http.ResponseWriter, r *http.Request) {
	if !midl.TokenEsAdmin {
		http.Error(w, "Solo un administrador puede ejecutar esta accion.", http.StatusBadRequest)
		return
	}

	var usuario structs.Usuarios
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		http.Error(w, "error al decodificar el JSON de la peticion: "+err.Error(), http.StatusBadRequest)
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

/* Inicia la sesion de un usuario. Genera el TOKEN. */
func IniciarSesion(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var u structs.Usuarios

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "error al decodificar el JSON de la peticion: "+err.Error(), http.StatusBadRequest)
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

	response := structs.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

/* Modifica la contraseña de su propio usuario. */
func CambiarPassword(w http.ResponseWriter, r *http.Request) {
	var cambios structs.CambiarPassword
	err := json.NewDecoder(r.Body).Decode(&cambios)
	if err != nil {
		http.Error(w, "error al decodificar el JSON de la peticion: "+err.Error(), http.StatusBadRequest)
		return
	}

	/* Sanitizamos */
	cambios.Password = strings.TrimSpace(cambios.Password)
	cambios.NuevaPassword = strings.TrimSpace(cambios.NuevaPassword)
	cambios.Confirmacion = strings.TrimSpace(cambios.Confirmacion)

	if len(cambios.Password) == 0 {
		http.Error(w, "La contraseña anterior es requerida.", http.StatusBadRequest)
		return
	}

	if len(cambios.NuevaPassword) == 0 {
		http.Error(w, "La contraseña nueva es requerida.", http.StatusBadRequest)
		return
	}

	if len(cambios.Confirmacion) == 0 {
		http.Error(w, "La confirmacion de la contraseña nueva es requerida.", http.StatusBadRequest)
		return
	}

	if len(cambios.NuevaPassword) < 8 {
		http.Error(w, "Debe especificar una nueva contraseña de al menos 8 caracteres.", http.StatusBadRequest)
		return
	}

	if cambios.NuevaPassword != cambios.Confirmacion {
		http.Error(w, "La nueva contraseña no coincide con la confirmacion.", http.StatusBadRequest)
		return
	}

	_, exists := db.IniciarSesion(midl.TokenUsuario, cambios.Password)
	if !exists {
		http.Error(w, "Contraseña anterior erronea.", http.StatusBadRequest)
		return
	}

	/* Encriptamos la nueva PW. */
	cambios.NuevaPassword, _ = helpers.EncriptarPassword(cambios.NuevaPassword)
	errr := db.CambiarPassword(midl.TokenUsuario, cambios.NuevaPassword)
	if errr != nil {
		http.Error(w, "Error al modificar la contraseña."+errr.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
