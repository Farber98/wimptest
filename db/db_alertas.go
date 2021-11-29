package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* Trae todos los Sw de la BD. */
func ListarAlertas() ([]primitive.M, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_ALERTAS)

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
			"evento":       1,
			"device.ip":    1,
			"device.name":  1,
			"device.model": 1,
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

/* Devuelve Ranking de alertas segun $Srcmac. Ordena por cantidad desc. Limita 20. */
func RankingAlertasPorMac() []primitive.M {
	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_ALERTAS)

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

/* Devuelve todas las anomalias dada una $Srcmac. Ordena por timestamp desc. */
/* func AlertasPorMac(s structs.Switches) []primitive.M {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_DISPOSITIVOS)

	var results []primitive.M
	pipeline := make([]bson.M, 0)

	matchStage := bson.M{
		"$match": bson.M{
			"mac": s.Mac}}

	lookupStage := bson.M{
		"$lookup": bson.M{
			"from":         "dispositivos",
			"localField":   "mac",
			"foreignField": "mac",
			"as":           "dispositivo"}}

	projectStage := bson.M{
		"$project": bson.M{
			"mac":               1,
			"evento":            1,
			"mensaje":           1,
			"timestamp":         1,
			"dispositivo.ip":    1,
			"dispositivo.name":  1,
			"dispositivo.model": 1,
			"dispositivo.swmac": 1,
		}}

	sortStage := bson.M{
		"$sort": bson.M{
			"timestamp": -1}}

	pipeline = append(pipeline, matchStage, lookupStage, projectStage, sortStage)

	cursor, err := coll.Aggregate(ctx, pipeline)

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

} */

/* func Iterador(routineCtx context.Context, stream *mongo.ChangeStream) {
	defer stream.Close(routineCtx)
	for stream.Next(routineCtx) {
		var data bson.M = nil
		fmt.Println(data)
		if err := stream.Decode(&data); err != nil {
			panic(err)
		}
		fmt.Println(data)
	}
}

func TriggerAlerta() {
	db := MongoCN.Database(DB_NOMBRE)
	col := db.Collection(COL_SWITCHES)
	alertasStream, err := col.Watch(context.TODO(), mongo.Pipeline{})
	if err != nil {
		panic(err)
	}
	go Iterador(context.TODO(), alertasStream)
} */

/* Trae todas las alertas de la DB, ordenadas en fecha descendente. */
/* func DameAlertas() ([]primitive.M, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_ALERTAS)

	var results []primitive.M
	// opciones := options.Find().SetSort(bson.D{{Key: "fecha", Value: -1}})
	lookupStage := bson.D{{"$lookup", bson.D{{"from", "switches"}, {"localField", "_pid"}, {"foreignField", "_id"}, {"as", "switch"}}}}
	unwindStage := bson.D{{"$unwind", bson.D{{"path", "$switch"}, {"preserveNullAndEmptyArrays", false}}}}
	projectStage := bson.D{{"$project", bson.D{{"_pid", 0}, {"switch.lat", 0}, {"switch.lng", 0}, {"switch.fecha", 0}}}}
	sortStage := bson.D{{"$sort", bson.D{{"fecha", -1}}}}
	cursor, err := coll.Aggregate(ctx, mongo.Pipeline{lookupStage, unwindStage, projectStage, sortStage})
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
} */

/* Borra un switch por ID. */
/* func ConfirmarAlerta(ID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	col := db.Collection(COL_ALERTAS)

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{"_id": objID}

	_, err := col.DeleteOne(ctx, condition)
	return err

} */

/* Chequea si el ID de la alerta ya existe en la BD.*/
/* func ExisteIdAlertas(ID primitive.ObjectID) (structs.Alertas, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var condition primitive.M
	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_ALERTAS)
	condition = bson.M{"_id": ID}
	var result structs.Alertas
	err := coll.FindOne(ctx, condition).Decode(&result)
	PID := result.IdAlerta.Hex()
	if err != nil {
		return result, false, PID
	}

	return result, true, PID
} */
