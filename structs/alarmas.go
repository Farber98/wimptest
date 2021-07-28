package structs

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Alarmas struct {
	IdAlarma primitive.ObjectID `bson:"_id,omitempty" json:"idHost"`
	IdSwitch primitive.ObjectID `bson:"_pid,omitempty" json:"idSwitch"`
}
