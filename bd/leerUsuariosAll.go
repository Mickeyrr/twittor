package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/Mickeyrr/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* LeerUsuariosAll obtenemos todos los usuario por un filtro de busqueda y tipo "R" son los que se relacionesn conmigo */
func LeerUsuariosAll(ID string, pagina int64, search string, tipo string) ([]*models.Usuario, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	var results []*models.Usuario

	findOptions := options.Find()
	findOptions.SetSkip((pagina - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	var exist, incluir bool
	for cur.Next(ctx) {
		var u models.Usuario
		err := cur.Decode(&u)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}
		var r models.Relacion
		r.UsuarioID = ID
		r.UsuarioRelacionID = u.ID.Hex()
		incluir = false
		exist, err = ConsultarRelacion(r)
		// validamos a los usuario que no sigo
		if tipo == "new" && !exist {
			incluir = true
		}
		// Validamos a los usuarios que sigo
		if tipo == "follow" && exist {
			incluir = true
		}
		// Validamos que no me siga a mi mismo
		if r.UsuarioRelacionID == ID {
			incluir = false
		}

		if incluir {
			u.Password = ""
			u.Biografia = ""
			u.SitioWeb = ""
			u.Ubicacion = ""
			u.Banner = ""
			u.Email = ""

			results = append(results, &u)
		}

	}

	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	cur.Close(ctx)
	return results, true
}
