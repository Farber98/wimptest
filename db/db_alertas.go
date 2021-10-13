package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

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
func DameAlertas() ([]primitive.M, bool) {
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
}

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
