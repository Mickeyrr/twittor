package routers

import (
	"net/http"

	"github.com/Mickeyrr/twittor/bd"
	"github.com/Mickeyrr/twittor/models"
)

/* DeleteRelacion Elimina la relacion entrs usuarios */
func DeleteRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}
	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID
	status, err := bd.DeleteRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar eliminar relacion "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado eliminar relacion "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
