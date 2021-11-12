package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Mickeyrr/twittor/bd"
)

/* VerPerfil permite ver los datos de un perfil */
func VerPerfil(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe de enviar el parameto ID, ", http.StatusBadRequest)
		return
	}
	perfil, err := bd.BuscarPerfil(ID)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar buscar el registro"+err.Error(), 400)
		return
	}
	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}
