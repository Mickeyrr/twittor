package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/Mickeyrr/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

/* ConsultarRelacion proceso para comprobar la relacion entre 2 usuarios */
func ConsultarRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}
	var result models.Relacion
	fmt.Println(result)
	err := col.FindOne(ctx, condicion).Decode(&result)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}
