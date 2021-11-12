package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* MongoCN esl el objeto de conexion a la BD */
var MongoCN = ConectarDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://usergo:1uGxHBbq24hVmIb7@angularcluster.pzleh.mongodb.net/twittor?retryWrites=true&w=majority")

/* ConectarDB Esta funcion me permite conectarme a la BD */
func ConectarDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion Existosa con la BD")
	return client
}

/* CheckConnection funcion para revisar el status de la conexion*/
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1

}
