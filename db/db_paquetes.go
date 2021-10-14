package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

/* Devuelve la cantidad de paquetes y bytes transmitidos con $srcMac. Ordena por bytes y paquetes desc. Limita 20.  */
func DameSrcMacMayorEmision() []primitive.M {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_PAQUETES)

	var results []primitive.M

	projectStage := bson.D{{"$project", bson.D{{"srcmac", 1}, {"length", 1}}}}
	groupStage := bson.D{{"$group", bson.D{{"_id", "$srcmac"}, {"paquetes", bson.D{{"$sum", 1}}}, {"bytes", bson.D{{"$sum", "$length"}}}}}}
	sortStage := bson.D{{"$sort", bson.D{{"bytes", -1}, {"paquetes", -1}}}}
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

/* Devuelve la cantidad de paquetes y bytes transmitidos con $srcIp. Ordena por bytes y paquetes desc. Limita 20.  */
func DameSrcIpMayorEmision() []primitive.M {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_PAQUETES)

	var results []primitive.M

	projectStage := bson.D{{"$project", bson.D{{"srcip", 1}, {"length", 1}}}}
	groupStage := bson.D{{"$group", bson.D{{"_id", "$srcip"}, {"paquetes", bson.D{{"$sum", 1}}}, {"bytes", bson.D{{"$sum", "$length"}}}}}}
	sortStage := bson.D{{"$sort", bson.D{{"bytes", -1}, {"paquetes", -1}}}}
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

/* Devuelve la cantidad de paquetes con $ProtoApp . Ordena por total desc.*/
func DameProtocolosAplicacionMayorEmision() []primitive.M {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_PAQUETES)

	var results []primitive.M

	projectStage := bson.D{{"$project", bson.D{{"protoapp", 1}}}}
	groupStage := bson.D{{"$group", bson.D{{"_id", "$protoapp"}, {"total", bson.D{{"$sum", 1}}}}}}
	sortStage := bson.D{{"$sort", bson.D{{"total", -1}}}}
	cursor, err := coll.Aggregate(ctx, mongo.Pipeline{projectStage, groupStage, sortStage})
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

/* Devuelve la cantidad de paquetes con $ProtoTp. Ordena por total desc. */
func DameProtocolosTransporteMayorEmision() []primitive.M {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_PAQUETES)

	var results []primitive.M

	projectStage := bson.D{{"$project", bson.D{{"prototp", 1}}}}
	groupStage := bson.D{{"$group", bson.D{{"_id", "$prototp"}, {"total", bson.D{{"$sum", 1}}}}}}
	sortStage := bson.D{{"$sort", bson.D{{"total", -1}}}}
	cursor, err := coll.Aggregate(ctx, mongo.Pipeline{projectStage, groupStage, sortStage})
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

/* Devuelve la cantidad de paquetes con $ProtoIp. Ordena por total desc.  */
func DameProtocolosRedMayorEmision() []primitive.M {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_PAQUETES)

	var results []primitive.M

	projectStage := bson.D{{"$project", bson.D{{"protoip", 1}}}}
	groupStage := bson.D{{"$group", bson.D{{"_id", "$protoip"}, {"total", bson.D{{"$sum", 1}}}}}}
	sortStage := bson.D{{"$sort", bson.D{{"total", -1}}}}
	cursor, err := coll.Aggregate(ctx, mongo.Pipeline{projectStage, groupStage, sortStage})
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

/* Devuelve la cantidad de paquetes dada una $srcMac */
func DameSrcMacEmision(mac string) []primitive.M {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_PAQUETES)

	var results []primitive.M

	matchStage := bson.D{{"$match", bson.D{{"srcmac", mac}}}}
	groupStage := bson.D{{"$group", bson.D{{"_id", "total"}, {"paquetes", bson.D{{"$sum", 1}}}, {"bytes", bson.D{{"$sum", "$length"}}}}}}
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
func DameDstIp(mac string) []primitive.M {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_PAQUETES)

	var results []primitive.M

	matchStage := bson.D{{"$match", bson.D{{"srcmac", mac}}}}
	groupStage := bson.D{{"$group", bson.D{{"_id", "$dstip"}, {"paquetes", bson.D{{"$sum", 1}}}, {"bytes", bson.D{{"$sum", "$length"}}}}}}
	sortStage := bson.D{{"$sort", bson.D{{"bytes", -1}, {"paquetes", -1}}}}
	cursor, err := coll.Aggregate(ctx, mongo.Pipeline{matchStage, groupStage, sortStage})
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
