package helpers

import (
	"time"

	"github.com/Farber98/WIMP/structs"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

/* Genera JWT. */
func GenerarJwt(t structs.Usuarios) (string, error) {

	/* Privilegios en payload */
	payload := jwt.MapClaims{
		"_id":      t.IdUsuario.Hex(),
		"username": t.Usuario,
		"email":    t.Email,
		"exp":      time.Now().Add(time.Hour * 12).Unix(),
	}

	/* Define jwt con privilegios y  algoritmo de firmado */
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	/* Firmo con la clave JWT_KEY */
	clave := []byte(JWT_KEY)
	tokenStr, err := token.SignedString(clave)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}

/* Encriptar password, usada por CrearUsuario.  */
func EncriptarPassword(password string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}
