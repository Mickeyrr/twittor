package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Mickeyrr/twittor/bd"
	"github.com/Mickeyrr/twittor/models"
)

/* ConsultarRelacion obtenemos la relacion entre 2 usuario */
func ConsultarRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}
	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	var result models.ConsultarRelacion

	status, err := bd.ConsultarRelacion(t)
	if err != nil || !status {
		result.Status = false
	} else {
		result.Status = true
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
