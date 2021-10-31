package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* Devuelve todas las Anomalias de la bd*/
func ListarAnomalias() ([]primitive.M, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_ANOMALIAS)

	pipeline := make([]bson.M, 0)
	var results []primitive.M

	lookupStage := bson.M{
		"$lookup": bson.M{
			"from":         "dispositivos",
			"localField":   "mac",
			"foreignField": "mac",
			"as":           "device"}}

	unwindStage := bson.M{
		"$unwind": bson.M{
			"path":                       "$device",
			"preserveNullAndEmptyArrays": true}}

	projectStage := bson.M{
		"$project": bson.M{
			"mac":          1,
			"device.ip":    1,
			"device.name":  1,
			"device.model": 1,
			"device.swmac": 1,
			"device.type":  1,
			"timestamp":    1}}

	sortStage := bson.M{
		"$sort": bson.M{
			"timestamp": -1}}

	pipeline = append(pipeline, lookupStage, unwindStage, projectStage, sortStage)

	cursor, err := coll.Aggregate(ctx, pipeline)
	if err != nil {
		log.Println(err.Error())
		return nil, false
	}

	for cursor.Next(context.Background()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			return nil, false
		}
		results = append(results, result)
	}

	if err := cursor.Err(); err != nil {
		return nil, false
	}

	cursor.Close(context.Background())

	return results, true
}

/* Devuelve Ranking de Anomalias segun $Srcmac. Ordena por cantidad desc. Limita 20. */
func RankingAnomalias() []primitive.M {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_ANOMALIAS)

	var results []primitive.M

	pipeline := make([]bson.M, 0)

	lookupStage := bson.M{
		"$lookup": bson.M{
			"from":         "dispositivos",
			"localField":   "mac",
			"foreignField": "mac",
			"as":           "dispositivo"}}

	unwindStage := bson.M{
		"$unwind": bson.M{
			"path":                       "$dispositivo",
			"preserveNullAndEmptyArrays": true}}

	groupStage := bson.M{
		"$group": bson.M{
			"_id": bson.M{
				"mac":  "$mac",
				"ip":   "$dispositivo.ip",
				"name": "$dispositivo.name",
			},
			"cant": bson.M{
				"$sum": 1}}}

	projectStage := bson.M{
		"$project": bson.M{
			"mac":              1,
			"dispositivo.ip":   1,
			"dispositivo.name": 1}}

	sortStage := bson.M{
		"$sort": bson.M{
			"cant": -1}}

	limitStage := bson.M{"$limit": 20}

	pipeline = append(pipeline, lookupStage, unwindStage, projectStage, groupStage, sortStage, limitStage)

	data, err := coll.Aggregate(ctx, pipeline)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	err = data.All(ctx, &results)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return results
}

/* Devuelve anomalias dada una $Srcmac. Limita 20. */
/* func AnomaliasSrcMac(s structs.Switches) []primitive.M {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_ANOMALIAS)

	var results []primitive.M

	matchStage := bson.D{{"$match", bson.D{{"mac", s.Mac}}}}
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
*/
