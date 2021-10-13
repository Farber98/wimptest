package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

/* Devuelve todas las Anomalias */
func DameAnomalias() ([]primitive.M, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_ANOMALIAS)

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

/* Devuelve Ranking de Anomalias segun $Srcmac */
func RankingAnomaliasPorMac() []primitive.M {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_ANOMALIAS)

	var results []primitive.M

	projectStage := bson.D{{"$project", bson.D{{"mac", 1}}}}
	groupStage := bson.D{{"$group", bson.D{{"_id", "$mac"}, {"cant", bson.D{{"$sum", 1}}}}}}
	sortStage := bson.D{{"$sort", bson.D{{"cant", -1}}}}
	limitStage := bson.D{{"$limit", 20}}
	cursor, err := coll.Aggregate(ctx, mongo.Pipeline{projectStage, groupStage, sortStage, limitStage})
	if err != nil {
		return results
	}

	for cursor.Next(context.Background()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			return results
		}
		results = append(results, result)
	}

	if err := cursor.Err(); err != nil {
		return results
	}

	cursor.Close(context.Background())

	return results
}

/* Devuelve anomalias de una SrcMac */
func DameSrcMacAnomalias(mac string) []primitive.M {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_ANOMALIAS)

	var results []primitive.M

	matchStage := bson.D{{"$match", bson.D{{"mac", mac}}}}
	projectStage := bson.D{{"$project", bson.D{{"mac", 1}, {"anomaly", 1}, {"timestamp", 1}}}}
	sortStage := bson.D{{"$sort", bson.D{{"timestamp", -1}}}}
	cursor, err := coll.Aggregate(ctx, mongo.Pipeline{projectStage, matchStage, sortStage})
	if err != nil {
		return results
	}

	for cursor.Next(context.Background()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			return results
		}
		results = append(results, result)
	}

	if err := cursor.Err(); err != nil {
		return results
	}

	cursor.Close(context.Background())

	return results

}
