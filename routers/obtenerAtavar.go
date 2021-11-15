package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/Mickeyrr/twittor/bd"
)

/* ObtenerAvatar recupara la imagen de la base de datos*/
func ObtenerAvatar(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametros id", http.StatusBadGateway)
		return
	}
	perfil, err := bd.BuscarPerfil(ID)
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusBadGateway)
		return
	}
	OpenFile, err := os.Open("uploads/avatars/" + perfil.Avatar)
	if err != nil {
		http.Error(w, "Imagen no encontrada", http.StatusBadGateway)
		return
	}
	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error al copiar la imagen", http.StatusBadGateway)
	}
}
