package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Mickeyrr/twittor/bd"
	"github.com/Mickeyrr/twittor/models"
)

/*Registro es la funcion para crear en la BD el registro de usuario */
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode((&t))
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido ", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "El password debe ser de al menos 6 caracteres ", 400)
		return
	}

	_, existe, _ := bd.CheckUserExist(t.Email)
	if existe {
		http.Error(w, "Ya existe un usuario registrado con ese email ", 400)
		return
	}
	_, status, err := bd.InsertRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro del usuario "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "No se a logrado insertar el registro del usuario ", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
