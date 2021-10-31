package main

import (
	manejadores "github.com/Farber98/WIMP/handlers"
	midl "github.com/Farber98/WIMP/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func configuroRouter() (handlers.CORSOption, handlers.CORSOption, handlers.CORSOption, *mux.Router) {
	r := mux.NewRouter()

	/* Configuracion del router */
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE", "PATCH"})
	origins := handlers.AllowedOrigins([]string{"*"})

	/* Rutas */

	/* Usuarios */
	r.HandleFunc("/usuarios/crear", midl.ValidarJwt(midl.ChequeoDB(manejadores.CrearUsuario))).Methods("POST")
	r.HandleFunc("/usuarios/iniciar-sesion", midl.ChequeoDB((manejadores.IniciarSesion))).Methods("POST")
	r.HandleFunc("/usuarios/cambiar-password", midl.ValidarJwt(midl.ChequeoDB(manejadores.CambiarPassword))).Methods("PATCH")

	/* Switches */
	//r.HandleFunc("/switches/crear", midl.ValidarJwt(midl.ChequeoDB(manejadores.CrearSwitch))).Methods("POST")
	//r.HandleFunc("/switches/modificar", midl.ValidarJwt(midl.ChequeoDB(manejadores.ModificarSwitch))).Methods("PUT")
	//r.HandleFunc("/switches/borrar", midl.ValidarJwt(midl.ChequeoDB(manejadores.BorrarSwitch))).Methods("DELETE")
	//r.HandleFunc("/switches/activar", midl.ValidarJwt(midl.ChequeoDB(manejadores.ActivarSwitch))).Methods("PUT")
	//r.HandleFunc("/switches/desactivar", midl.ValidarJwt(midl.ChequeoDB(manejadores.DesactivarSwitch))).Methods("PUT")
	r.HandleFunc("/switches", midl.ValidarJwt(midl.ChequeoDB(manejadores.ListarSwitches))).Methods("GET")
	r.HandleFunc("/switches/ubicar", midl.ValidarJwt(midl.ChequeoDB(manejadores.UbicarSwitch))).Methods("PATCH")

	/* Alertas */
	//r.HandleFunc("/alertas/confirmar", midl.ValidarJwt(midl.ChequeoDB(manejadores.ConfirmarAlerta))).Methods("DELETE")
	//r.HandleFunc("/alertas/mac", midl.ValidarJwt(midl.ChequeoDB(manejadores.AlertasPorMac))).Methods("GET")
	r.HandleFunc("/alertas", midl.ValidarJwt(midl.ChequeoDB(manejadores.ListarAlertas))).Methods("GET")
	r.HandleFunc("/alertas/ranking", midl.ValidarJwt(midl.ChequeoDB(manejadores.RankingAlertasPorMac))).Methods("GET")

	/* Paquetes */
	r.HandleFunc("/paquetes/srcmac-emision", midl.ValidarJwt(midl.ChequeoDB(manejadores.RankingSrcMacTransmision))).Methods("GET")
	r.HandleFunc("/paquetes/protoapp-emision", midl.ValidarJwt(midl.ChequeoDB(manejadores.RankingProtoApp))).Methods("GET")
	r.HandleFunc("/paquetes/prototp-emision", midl.ValidarJwt(midl.ChequeoDB(manejadores.RankingProtoTransporte))).Methods("GET")
	r.HandleFunc("/paquetes/protoip-emision", midl.ValidarJwt(midl.ChequeoDB(manejadores.RankingProtoRed))).Methods("GET")
	r.HandleFunc("/paquetes/srcmac-detalle", midl.ValidarJwt(midl.ChequeoDB(manejadores.DetalleSrcMac))).Methods("GET")

	/* Topologia */
	r.HandleFunc("/topologia", midl.ValidarJwt(midl.ChequeoDB(manejadores.ListarTopologia))).Methods("GET")

	/* Anomalias */
	//r.HandleFunc("/anomalias/mac", midl.ValidarJwt(midl.ChequeoDB(manejadores.AnomaliasSrcMac))).Methods("GET")
	r.HandleFunc("/anomalias", midl.ValidarJwt(midl.ChequeoDB(manejadores.ListarAnomalias))).Methods("GET")
	r.HandleFunc("/anomalias/ranking", midl.ValidarJwt(midl.ChequeoDB(manejadores.RankingAnomalias))).Methods("GET")

	return headers, origins, methods, r
}
