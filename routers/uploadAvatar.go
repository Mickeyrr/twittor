package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Mickeyrr/twittor/bd"
	"github.com/Mickeyrr/twittor/models"
)

/* UploadAvatar subir el avatar al servidor y actualizar el registro en la BD */
func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/avatars/" + IDUsuario + "." + extension
	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}
	// Copia el archivo file a f (ya con el nombre del IDUsuario)
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}
	var usuario models.Usuario
	var status bool
	usuario.Avatar = IDUsuario + "." + extension
	status, err = bd.ModificarRegistro(usuario, IDUsuario)
	if err != nil || !status {
		http.Error(w, "Error al guardar el avatar en el BD ! "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
