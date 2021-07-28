package structs

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Alertas struct {
	IdAlerta primitive.ObjectID `bson:"_id,omitempty" json:"idAlerta"`
	IdSwitch primitive.ObjectID `bson:"_pid,omitempty" json:"idSwitch"`
	Problema string             `bson:"problema" json:"problema,omitempty"`
	Fecha    time.Time          `bson:"fecha" json:"fecha,omitempty"`
}
