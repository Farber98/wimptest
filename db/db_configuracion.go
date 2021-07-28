package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* Variable que exporta la conexion a la BD */
var MongoCN = ConectarBD()

/* clientOptions for MongoDB. */
var clientOptions = options.Client().ApplyURI(os.Getenv(DB_HOST))

/* Funcion para conectarse a la BD. */
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("MongoDB live")
	return client
}

/* Chequea que la conexion este activa antes de ejecutar un Handler. */
func ChequeoConexion() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
