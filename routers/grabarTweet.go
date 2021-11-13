package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Mickeyrr/twittor/bd"
	"github.com/Mickeyrr/twittor/models"
)

/** GrabarTweet permite inserta el tweeet en la BD */
func GrabarTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GrabarTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar indertar el registro"+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "No se a logrado insertar el Tweet", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
