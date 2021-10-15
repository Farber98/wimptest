package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type Usuarios struct {
	IdUsuario primitive.ObjectID `bson:"_id,omitempty" json:"idUsuario"`
	Usuario   string             `bson:"usuario" json:"usuario,omitempty"`
	Email     string             `bson:"email" json:"email,omitempty"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Admin     bool               `bson:"admin" json:"admin,omitempty"`
}

type CambiarPassword struct {
	Password      string `bson:"password" json:"password,omitempty"`
	NuevaPassword string `bson:"nuevapassword" json:"nuevapassword,omitempty"`
	Confirmacion  string `bson:"confirmacion" json:"confirmacion,omitempty"`
}
