package routers

import (
	"net/http"

	"github.com/FedeDifulvio/Curso-GO-React/bd"
	"github.com/FedeDifulvio/Curso-GO-React/models"
)

func AltaRelacion(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro id ", http.StatusBadRequest)
		return
	}

	var relacion models.Relacion

	relacion.UsuarioID = IdUsuario
	relacion.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(relacion)

	if !status || err != nil {
		http.Error(w, "Error al guardar la relación: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
