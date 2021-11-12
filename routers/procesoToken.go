package routers

import (
	"errors"
	"strings"

	"github.com/Mickeyrr/twittor/bd"
	"github.com/Mickeyrr/twittor/key"
	"github.com/Mickeyrr/twittor/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/* Email valor usado en todos los EndPoint */
var Email string

/* IDUsuario valor usado en todos los EndPoint */
var IDUsuario string

/* ProcesoToken proceso token para extraer sus valores */
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	myKey := key.MyKey()

	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invaldo")
	}
	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err == nil {
		_, exit, _ := bd.CheckUserExist(claims.Email)
		if exit {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, exit, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}
	return claims, false, string(""), err
}
