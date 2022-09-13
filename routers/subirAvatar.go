package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/FedeDifulvio/Curso-GO-React/bd"
	"github.com/FedeDifulvio/Curso-GO-React/models"
)

func SubirAvatar(w http.ResponseWriter, r *http.Request) {

	file, handler, error := r.FormFile("avatar")

	if error != nil {
		http.Error(w, "Error al obtener el avatar: "+error.Error(), http.StatusBadRequest)
		return
	}

	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/avatars/" + IdUsuario + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		http.Error(w, "Error al subir el avatar: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)

	if err != nil {
		http.Error(w, "Error al copiar el avatar: "+err.Error(), http.StatusBadRequest)
		return
	}

	var usuario models.Usuario
	var status bool

	usuario.Avatar = IdUsuario + "." + extension

	status, err = bd.ModificoRegistro(usuario, IdUsuario)

	if err != nil || !status {
		http.Error(w, "Error al grabar el avatar: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
