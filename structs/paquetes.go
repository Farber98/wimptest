package structs

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Paquetes struct {
	IdPaquete primitive.ObjectID `bson:"_id,omitempty" json:"idPaquete"`
	SrcMac    string             `bson:"srcMac" json:"srcMac,omitempty"`
	DstMac    string             `bson:"dstMac" json:"dstMac,omitempty"`
	ProtoIp   string             `bson:"protoIp" json:"protoIp,omitempty"`
	SrcIp     string             `bson:"srcIp" json:"srcIp,omitempty"`
	DstIp     string             `bson:"dstIp" json:"dstIp,omitempty"`
	ProtoTp   string             `bson:"protoTp" json:"protoTp,omitempty"`
	SrcTp     string             `bson:"srcTp" json:"srcTp,omitempty"`
	DstTp     string             `bson:"dstTp" json:"dstTp,omitempty"`
	ProtoApp  string             `bson:"protoApp" json:"protoApp,omitempty"`
	Length    int                `bson:"length" json:"length,omitempty"`
	Timestamp time.Time          `bson:"timestamp" json:"timestamp,omitempty"`
}
