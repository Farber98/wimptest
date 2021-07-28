package structs

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Hosts struct {
	IdHost   primitive.ObjectID `bson:"_id,omitempty" json:"idHost"`
	IdSwitch primitive.ObjectID `bson:"_pid,omitempty" json:"idSwitch"`
	Email    string             `bson:"email" json:"email,omitempty"`
	Mac      string             `bson:"modelo" json:"modelo,omitempty"`
	Fecha    time.Time          `bson:"fecha" json:"fecha,omitempty"`
	Estado   bool               `bson:"estado" json:"estado,omitempty"`
}
