package jwt

import (
	"time"

	"github.com/Mickeyrr/twittor/key"
	"github.com/Mickeyrr/twittor/models"
	jwt "github.com/dgrijalva/jwt-go"
)

func CreateJWT(t models.Usuario) (string, error) {
	// myKey := []byte("CursoGolangMiguelRomero2@21")
	myKey := key.MyKey()

	payload := jwt.MapClaims{
		"email":           t.Email,
		"nombre":          t.Nombre,
		"apellidos":       t.Apellidos,
		"fechaNacimiento": t.FechaNacimiento,
		"biografia":       t.Biografia,
		"ubicacion":       t.Ubicacion,
		"sitioWeb":        t.SitioWeb,
		"_id":             t.ID.Hex(),
		"exp":             time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
