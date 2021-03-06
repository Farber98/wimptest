package main

import (
	stdlog "log"
	"net/http"
	"os"

	midl "github.com/Farber98/WIMP/middleware"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/handlers"
)

func main() {

	/* Configuracion del router. */
	headers, methods, origins, router := configuroRouter()

	/* Configuracion logger */
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	stdlog.SetOutput(log.NewStdlibAdapter(logger))
	logger = log.With(logger, "ts", log.DefaultTimestamp, "loc", log.DefaultCaller)

	/* Asignacion middleware */
	loggingMiddleware := midl.LoggingMiddleware(logger)

	/* Vinculamos logger al router */
	loggedRouter := loggingMiddleware(router)

	//go db.TriggerAlerta()

	/* Corremos SV. */
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = SV_PORT
	}
	if err := http.ListenAndServe(":"+PORT, handlers.CORS(headers, methods, origins)(loggedRouter)); err != nil {
		logger.Log("status", "fatal", "err", err)
		os.Exit(1)
	}
}
