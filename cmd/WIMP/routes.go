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
	r.HandleFunc("/usuarios/crear", (midl.ChequeoDB((manejadores.CrearUsuario)))).Methods("POST")
	r.HandleFunc("/usuarios/iniciar-sesion", (midl.ChequeoDB((manejadores.IniciarSesion)))).Methods("POST")

	return headers, origins, methods, r
}
