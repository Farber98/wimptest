package midl

import (
	"errors"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"github.com/Farber98/WIMP/db"
	"github.com/Farber98/WIMP/structs"
	"github.com/dgrijalva/jwt-go"
	log "github.com/go-kit/kit/log"
)

/* Wrappers para agarrar el estado HTTP y logearlo. */
type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true

}

/* Logea las peticiones HTTP y su duracion. */
func LoggingMiddleware(logger log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					logger.Log(
						"err", err,
						"trace", debug.Stack(),
					)
				}
			}()

			start := time.Now()
			wrapped := wrapResponseWriter(w)
			next.ServeHTTP(wrapped, r)
			logger.Log(
				"status", wrapped.status,
				"method", r.Method,
				"path", r.URL.EscapedPath(),
				"duration", time.Since(start),
			)
		}

		return http.HandlerFunc(fn)
	}
}

/* Ejecuta ChequeoConexion() antes de la ejecucion del handler, para ver si la BD sigue activa. */
func ChequeoDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.ChequeoConexion() == 0 {
			http.Error(w, "Conexion con BD perdida.", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	}
}

/* Llama a ProcesarToken() para validar el jwt en una request */
func ValidarJwt(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := ProcesarToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el Token: "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}

var TokenUsuario string
var TokenIdUsuario string

/* Funcion utilizada por middleware ValidateJwt para extraer atributos de un token */
func ProcesarToken(tk string) (*structs.Claim, bool, string, error) {
	clave := []byte(JWT_KEY)
	claims := &structs.Claim{}

	/* Separa el token desde el delimitador Bearer */
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return clave, nil
	})
	if err == nil {
		_, found, ID := db.UsuarioDuplicado(claims.Usuario)
		if found {
			TokenUsuario = claims.Usuario
			TokenIdUsuario = ID
		}

		return claims, found, TokenIdUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err
}
