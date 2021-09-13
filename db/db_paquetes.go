package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

/* Devuelve la cantidad de paquetes con $srcMac */
func DameSrcMacMayorEmision() []primitive.M {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_PAQUETES)

	var results []primitive.M

	projectStage := bson.D{{"$project", bson.D{{"srcmac", 1}}}}
	groupStage := bson.D{{"$group", bson.D{{"_id", "$srcmac"}, {"total", bson.D{{"$sum", 1}}}}}}
	cursor, err := coll.Aggregate(ctx, mongo.Pipeline{projectStage, groupStage})
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

/* Devuelve la cantidad de paquetes con $ProtoApp */
func DameProtocolosAplicacionMayorEmision() []primitive.M {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_PAQUETES)

	var results []primitive.M

	projectStage := bson.D{{"$project", bson.D{{"protoapp", 1}}}}
	groupStage := bson.D{{"$group", bson.D{{"_id", "$protoapp"}, {"total", bson.D{{"$sum", 1}}}}}}
	cursor, err := coll.Aggregate(ctx, mongo.Pipeline{projectStage, groupStage})
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

/* Devuelve la cantidad de paquetes con $ProtoTp */
func DameProtocolosTransporteMayorEmision() []primitive.M {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_PAQUETES)

	var results []primitive.M

	projectStage := bson.D{{"$project", bson.D{{"prototp", 1}}}}
	groupStage := bson.D{{"$group", bson.D{{"_id", "$prototp"}, {"total", bson.D{{"$sum", 1}}}}}}
	cursor, err := coll.Aggregate(ctx, mongo.Pipeline{projectStage, groupStage})
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

/* Devuelve la cantidad de paquetes con $ProtoTp */
func DameProtocolosRedMayorEmision() []primitive.M {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_PAQUETES)

	var results []primitive.M

	projectStage := bson.D{{"$project", bson.D{{"protoip", 1}}}}
	groupStage := bson.D{{"$group", bson.D{{"_id", "$protoip"}, {"total", bson.D{{"$sum", 1}}}}}}
	cursor, err := coll.Aggregate(ctx, mongo.Pipeline{projectStage, groupStage})
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
