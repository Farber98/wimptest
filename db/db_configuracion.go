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

/* Para usar variables de entorno y cargar la URI. */
var clientOptions = options.Client().ApplyURI(os.Getenv("MONGODB_URI"))

/* Reutilizacion de conexion a la BD. */
func ConectarBD() *mongo.Client {
	/* Para pruebas locales. */
	if os.Getenv("MONGODB_URI") == "" {
		clientOptions = options.Client().ApplyURI(DB_HOST)
	}
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	/* Chequeamos que la BD este activa. */
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("MongoDB live")
	return client
}

/* Metodo utilizado por el Middleware ChequeoDb. Revisa que la db este activa antes de ejecutar un Handler. */
func ChequeoConexion() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
