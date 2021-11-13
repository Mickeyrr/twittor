package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Mickeyrr/twittor/bd"
	"github.com/Mickeyrr/twittor/models"
)

/* ModificarPerfil modifica el perfil del usuario */
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos"+err.Error(), 400)
		return
	}
	// Lo creamos asi ya que err ya existe arriba
	var status bool
	status, err = bd.ModificarRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "No se pudo modificar el registro en la BD"+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado modificar el registro del usuariopudo modificar"+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
