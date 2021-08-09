package structs

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SwitchWrapper struct {
	Switch Switches
}

type Switches struct {
	IdSwitch primitive.ObjectID `bson:"_id,omitempty" json:"idSwitch"`
	IdPadre  primitive.ObjectID `bson:"_pid,omitempty" json:"idPadre"`
	Nombre   string             `bson:"nombre" json:"nombre,omitempty"`
	Modelo   string             `bson:"modelo" json:"modelo,omitempty"`
	Lat      float64            `bson:"lat" json:"lat,omitempty"`
	Lng      float64            `bson:"lng" json:"lng,omitempty"`
	Fecha    time.Time          `bson:"fecha" json:"fecha,omitempty"`
	Estado   bool               `bson:"estado" json:"estado,omitempty"`
}
