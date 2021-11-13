package bd

import (
	"context"
	"log"
	"time"

	"github.com/Mickeyrr/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* LeerTweet lee los tweets de un perfil*/
func LeerTweet(ID string, pagina int64) ([]*models.ObtenerTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	var results []*models.ObtenerTweets

	condicion := bson.M{
		"userID": ID,
	}

	opciones := options.Find()
	opciones.SetLimit(20)                               // Numero de registros a obtener
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}}) // Ordenados Desc por fecha
	opciones.SetSkip((pagina - 1) * 20)                 // a la pagina se le resta 1 y se multiblica por el limit que es 20

	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	for cursor.Next(context.TODO()) {
		var registro models.ObtenerTweets
		err := cursor.Decode(&registro)
		if err != nil {
			return results, false
		}
		results = append(results, &registro)
	}
	return results, true
}
