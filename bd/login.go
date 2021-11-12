package bd

import (
	"github.com/Mickeyrr/twittor/models"
	"golang.org/x/crypto/bcrypt"
)

/* Login funcion para validar credenciales de usuario y conraseña */
func Login(email string, password string) (models.Usuario, bool) {
	u, exist, _ := CheckUserExist(email)
	if !exist {
		return u, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(u.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return u, false
	}
	return u, true
}
