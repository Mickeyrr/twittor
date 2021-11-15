package bd

import (
	"context"
	"time"

	"github.com/Mickeyrr/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

/* LeerTweetsFollow Recuperar todos los tweets de las personas que sigo */
func LeerTweetsFollow(ID string, pagina int) ([]models.ObtenerTweetsFollow, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	skip := (pagina - 1) * 20
	condicion := make([]bson.M, 0)
	condicion = append(condicion, bson.M{"$match": bson.M{"usuarioid": ID}}) // filtrar por mi usuario
	condicion = append(condicion, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",             // Tabla con la voy a hacer la entidad relacion de tablas
			"localField":   "usuariorelacionid", // Campo con el que se hace la entidad relacion de la tabla relacion
			"foreignField": "userID",            // Mi usuario id de la tabla Tweet
			"as":           "tweet",             // Alias de la relacion
		},
	})
	condicion = append(condicion, bson.M{"$unwind": "$tweet"})                // quita la opcion de maestro detalle
	condicion = append(condicion, bson.M{"$sort": bson.M{"tweet.fecha": -1}}) // ordenados Desc = -1 | ACS = 1
	condicion = append(condicion, bson.M{"$skip": skip})                      // que pagina va a mostrar
	condicion = append(condicion, bson.M{"$limit": 20})                       // cuantos registros va a mostrar por pagina

	cursor, err := col.Aggregate(ctx, condicion)
	var results []models.ObtenerTweetsFollow
	// Mandamos los resultados a la variable result, pero si hay algun error se manda el mensaje a la variable err
	err = cursor.All(ctx, &results)
	if err != nil {
		return results, false
	}
	return results, true
}
