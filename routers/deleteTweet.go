package routers

import (
	"net/http"

	"github.com/Mickeyrr/twittor/bd"
)

/* DeleteTweet permite borrar un Tweet determinado*/
func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe de enviar el parametro ID", http.StatusBadRequest)
		return
	}
	err := bd.DeleteTweet(ID, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar borrar el tweet "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
