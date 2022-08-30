package routers

import (
	"net/http"

	"github.com/FedeDifulvio/Curso-GO-React/bd"
)

func EliminarTweet(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro id ", http.StatusBadRequest)
		return
	}

	err := bd.BorroTweet(ID, IdUsuario)

	if err != nil {
		http.Error(w, "Error asl borrar tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
