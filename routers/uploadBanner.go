package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Mickeyrr/twittor/bd"
	"github.com/Mickeyrr/twittor/models"
)

/* UploadBanner subir el banner al servidor y actualizar el registro en la BD */
func UploadBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("banner")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/banners/" + IDUsuario + "." + extension
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
	usuario.Banner = IDUsuario + "." + extension
	status, err = bd.ModificarRegistro(usuario, IDUsuario)
	if err != nil || !status {
		http.Error(w, "Error al guardar el banner en el BD ! "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
