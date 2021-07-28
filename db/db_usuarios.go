package db

import (
	"context"
	"time"

	"github.com/Farber98/WIMP/structs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

/* Chequea si un mail ya existe en la db */
func EmailDuplicado(email string) (structs.Usuarios, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_USUARIOS)

	condition := bson.M{"email": email}

	var result structs.Usuarios

	err := coll.FindOne(ctx, condition).Decode(&result)
	ID := result.IdUsuario.Hex()
	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}

/* Chequea si un usuario ya existe en la db */
func UsuarioDuplicado(nombre string) (structs.Usuarios, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	coll := db.Collection(COL_USUARIOS)

	condition := bson.M{"usuario": nombre}

	var result structs.Usuarios

	err := coll.FindOne(ctx, condition).Decode(&result)
	ID := result.IdUsuario.Hex()
	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}

/* Crea usuario en la BD. */
func CrearUsuario(usuario structs.Usuarios) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database(DB_NOMBRE)
	col := db.Collection(COL_USUARIOS)

	result, err := col.InsertOne(ctx, usuario)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}

/* Autentica usuario ante la BD. Utilizado por IniciarSesion. */
func IniciarSesion(usuario string, password string) (structs.Usuarios, bool) {

	/* Chequea que el nombre de usuario exista. Misma operacion que usuario duplicado. */
	u, exists, _ := UsuarioDuplicado(usuario)
	if !exists {
		return u, false
	}

	/*  Comparamos la contrasena de la BD con el hash de la password ingresada. */
	inputPassword := []byte(password)
	dbPassword := []byte(u.Password)
	err := bcrypt.CompareHashAndPassword(dbPassword, inputPassword)
	if err != nil {
		return u, false
	}

	return u, true
}
