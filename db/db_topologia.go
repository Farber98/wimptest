package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* Trae todos los switches de la DB */
func ListarTopologia() ([]primitive.M, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_TOPOLOGIA)

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
