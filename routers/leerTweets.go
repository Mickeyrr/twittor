package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Mickeyrr/twittor/bd"
)

/* LeerTweets leer los tweets */
func LeerTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro id", http.StatusBadRequest)
		return
	}
	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parametro pagina", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Debe enviar el parametro pagina con un valor mayoe a 0", http.StatusBadRequest)
		return
	}
	// convertinos pagiga a int64
	pag := int64(pagina)
	respuesta, correct := bd.LeerTweet(ID, pag)
	if !correct {
		http.Error(w, "Error al leer los Tweets", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
