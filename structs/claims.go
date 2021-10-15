package structs

import (
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* Usado para procesar JWT */
type Claim struct {
	IdUsuario primitive.ObjectID `bson:"_id" json:"idUsuario,omitempty"`
	Email     string             `bson:"email" json:"email,omitempty"`
	Usuario   string             `bson:"usuario" json:"usuario,omitempty"`
	Admin     string             `bson:"admin" json:"admin,omitempty"`
	jwt.StandardClaims
}
