package db

import (
	"context"
	"time"

	"github.com/Farber98/WIMP/structs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* Checks if a given switch name already exists in db */
func NombreDuplicado(nombre string) (structs.Switches, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_SWITCHES)

	condition := bson.M{"nombre": nombre}

	var result structs.Switches

	err := coll.FindOne(ctx, condition).Decode(&result)
	ID := result.IdSwitch.Hex()
	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}

/* Chequea si el ID ya existe en la BD. */
func ExisteId(ID primitive.ObjectID) (structs.Switches, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_SWITCHES)
	condition := bson.M{"_id": ID}
	var result structs.Switches
	err := coll.FindOne(ctx, condition).Decode(&result)
	PID := result.IdSwitch.Hex()
	if err != nil {
		return result, false, PID
	}

	return result, true, PID
}

/* Crea el Switch en la db. */
func CrearSwitch(s structs.Switches) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	col := db.Collection(COL_SWITCHES)

	result, err := col.InsertOne(ctx, s)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}

/* Trae todos los switches de la DB */
func DameTopologia() ([]primitive.M, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_SWITCHES)

	var results []primitive.M

	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		return results, false
	}

	for cursor.Next(context.Background()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			return results, false
		}
		results = append(results, result)
	}

	if err := cursor.Err(); err != nil {
		return results, false
	}

	cursor.Close(context.Background())

	return results, true
}