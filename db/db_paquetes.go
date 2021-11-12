package db

import (
	"context"
	"log"
	"time"

	"github.com/Farber98/WIMP/structs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* Devuelve la cantidad de paquetes y bytes transmitidos con $srcMac. Ordena por bytes y paquetes desc. Limita 20.  */
func RankingSrcMacTransmision() []primitive.M {
	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Second)
	defer cancel()
	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_PAQUETES)
	var results []primitive.M
	pipeline := make([]bson.M, 0)

	projectStage := bson.M{
		"$project": bson.M{
			"srcip":  1,
			"srcmac": 1,
			"length": 1}}

	groupStage := bson.M{
		"$group": bson.M{
			"_id": bson.M{
				"srcmac": "$srcmac",
				"srcip":  "$srcip",
			},
			"bytes":    bson.M{"$sum": "$length"},
			"paquetes": bson.M{"$sum": 1}}}

	sortStage := bson.M{
		"$sort": bson.M{
			"bytes":    -1,
			"paquetes": -1}}

	limitStage := bson.M{"$limit": 20}

	pipeline = append(pipeline, projectStage, groupStage, sortStage, limitStage)

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

/* Devuelve la cantidad de paquetes con $ProtoApp . Ordena por total desc.*/
func RankingProtoApp() ([]primitive.M, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_PAQUETES)

	var results []primitive.M
	pipeline := make([]bson.M, 0)

	projectStage := bson.M{
		"$project": bson.M{"protoapp": 1}}
	groupStage := bson.M{"$group": bson.M{
		"_id":   "$protoapp",
		"total": bson.M{"$sum": 1}}}
	sortStage := bson.M{
		"$sort": bson.M{"total": -1}}
	pipeline = append(pipeline, projectStage, groupStage, sortStage)

	data, err := coll.Aggregate(ctx, pipeline)
	if err != nil {
		log.Println(err.Error())
		return nil, false
	}

	err = data.All(ctx, &results)
	if err != nil {
		log.Println(err.Error())
		return nil, false
	}

	return results, true

}

/* Devuelve la cantidad de paquetes con $ProtoTp. Ordena por total desc. */
func RankingProtoTransporte() []primitive.M {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_PAQUETES)

	var results []primitive.M
	pipeline := make([]bson.M, 0)

	projectStage := bson.M{
		"$project": bson.M{
			"prototp": 1}}
	groupStage := bson.M{
		"$group": bson.M{
			"_id":   "$prototp",
			"total": bson.M{"$sum": 1}}}
	sortStage := bson.M{
		"$sort": bson.M{"total": -1}}
	pipeline = append(pipeline, projectStage, groupStage, sortStage)

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

/* Devuelve la cantidad de paquetes con $ProtoIp. Ordena por total desc.  */
func RankingProtoRed() []primitive.M {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_PAQUETES)

	var results []primitive.M
	pipeline := make([]bson.M, 0)

	projectStage := bson.M{
		"$project": bson.M{
			"protoip": 1}}
	groupStage := bson.M{
		"$group": bson.M{
			"_id":   "$protoip",
			"total": bson.M{"$sum": 1}}}
	sortStage := bson.M{
		"$sort": bson.M{"total": -1}}
	pipeline = append(pipeline, projectStage, groupStage, sortStage)

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

/* Devuelve la cantidad de paquetes dada una $srcMac */
func DetalleSrcMacEmision(s structs.Switches) []primitive.M {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_PAQUETES)

	var results []primitive.M
	pipeline := make([]bson.M, 0)

	matchStage := bson.M{
		"$match": bson.M{
			"srcmac": s.Mac}}
	groupStage := bson.M{
		"$group": bson.M{
			"_id":      "total",
			"bytes":    bson.M{"$sum": "$length"},
			"paquetes": bson.M{"$sum": 1}}}

	pipeline = append(pipeline, matchStage, groupStage)

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

/* Devuelve la cantidad de paquetes Ip dada una $srcMac */
/* func DameSrcMacProtoIp(mac string) []primitive.M {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_PAQUETES)

	var results []primitive.M

	matchStage := bson.D{{"$match", bson.D{{"srcmac", mac}}}}
	groupStage := bson.D{{"$group", bson.D{{"_id", "$protoip"}, {"cant", bson.D{{"$sum", 1}}}}}}
	cursor, err := coll.Aggregate(ctx, mongo.Pipeline{matchStage, groupStage})
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

/* Devuelve la cantidad de paquetes App dada una $srcMac */
/* func DameSrcMacProtoApp(mac string) []primitive.M {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_PAQUETES)

	var results []primitive.M

	matchStage := bson.D{{"$match", bson.D{{"srcmac", mac}}}}
	groupStage := bson.D{{"$group", bson.D{{"_id", "$protoapp"}, {"cant", bson.D{{"$sum", 1}}}}}}
	cursor, err := coll.Aggregate(ctx, mongo.Pipeline{matchStage, groupStage})
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

/* Devuelve la cantidad de paquetes Tp dada una $srcMac */
/* func DameSrcMacProtoTp(mac string) []primitive.M {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_PAQUETES)

	var results []primitive.M

	matchStage := bson.D{{"$match", bson.D{{"srcmac", mac}}}}
	groupStage := bson.D{{"$group", bson.D{{"_id", "$prototp"}, {"cant", bson.D{{"$sum", 1}}}}}}
	cursor, err := coll.Aggregate(ctx, mongo.Pipeline{matchStage, groupStage})
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

/* func DameDstPort(mac string) []primitive.M {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_PAQUETES)

	var results []primitive.M

	matchStage := bson.D{{"$match", bson.D{{"srcmac", mac}}}}
	groupStage := bson.D{{"$group", bson.D{{"_id", "$dsttp"}, {"cant", bson.D{{"$sum", 1}}}}}}
	cursor, err := coll.Aggregate(ctx, mongo.Pipeline{matchStage, groupStage})
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

/* Devuelve la cantidad de paquetes y bytes enviados a una $srcIp. Ordena por bytes desc paquetes desc. */
func DetalleSrcMacDstIp(s structs.Switches) []primitive.M {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_PAQUETES)

	var results []primitive.M
	pipeline := make([]bson.M, 0)

	matchStage := bson.M{
		"$match": bson.M{
			"srcmac": s.Mac}}
	groupStage := bson.M{
		"$group": bson.M{
			"_id":      "$dstip",
			"paquetes": bson.M{"$sum": 1},
			"bytes":    bson.M{"$sum": "$length"}}}
	sortStage := bson.M{
		"$sort": bson.M{
			"bytes":    -1,
			"paquetes": -1}}
	pipeline = append(pipeline, matchStage, groupStage, sortStage)

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
