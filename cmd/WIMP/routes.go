package main

import (
	manejadores "github.com/Farber98/WIMP/handlers"
	midl "github.com/Farber98/WIMP/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func configuroRouter() (handlers.CORSOption, handlers.CORSOption, handlers.CORSOption, *mux.Router) {
	r := mux.NewRouter()

	/* Configuracion general */
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	/* Rutas */

	/* Usuarios */
	r.HandleFunc("/usuarios/crear", midl.ChequeoDB((manejadores.CrearUsuario))).Methods("POST")
	r.HandleFunc("/usuarios/iniciar-sesion", midl.ChequeoDB((manejadores.IniciarSesion))).Methods("POST")

	/* Switches */
	r.HandleFunc("/switches/crear", midl.ValidarJwt(midl.ChequeoDB(manejadores.CrearSwitch))).Methods("POST")
	r.HandleFunc("/switches/topologia", midl.ValidarJwt(midl.ChequeoDB(manejadores.VerTopologia))).Methods("GET")
	r.HandleFunc("/switches/modificar", midl.ValidarJwt(midl.ChequeoDB(manejadores.ModificarSwitch))).Methods("PUT")
	r.HandleFunc("/switches/borrar", midl.ValidarJwt(midl.ChequeoDB(manejadores.BorrarSwitch))).Methods("DELETE")
	r.HandleFunc("/switches/activar", midl.ValidarJwt(midl.ChequeoDB(manejadores.ActivarSwitch))).Methods("PUT")
	r.HandleFunc("/switches/desactivar", midl.ValidarJwt(midl.ChequeoDB(manejadores.DesactivarSwitch))).Methods("PUT")

	/* Alertas */
	r.HandleFunc("/alertas", midl.ValidarJwt(midl.ChequeoDB(manejadores.VerAlertas))).Methods("GET")
	r.HandleFunc("/alertas/confirmar", midl.ValidarJwt(midl.ChequeoDB(manejadores.ConfirmarAlerta))).Methods("DELETE")
	return headers, origins, methods, r
}
