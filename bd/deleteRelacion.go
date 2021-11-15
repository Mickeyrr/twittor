package bd

import (
	"context"
	"time"

	"github.com/Mickeyrr/twittor/models"
)

/* DeleteRelacion elimina la relacion entre los usuarios (dejar de seguir) */
func DeleteRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")
	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
