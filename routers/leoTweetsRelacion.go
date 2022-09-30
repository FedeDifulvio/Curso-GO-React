package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/FedeDifulvio/Curso-GO-React/bd"
)

func LeoTweetsSeguidores(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("pagina")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro id ", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))

	if err != nil {
		http.Error(w, "Debe enviar el parámetro como entero mayor a cero ", http.StatusBadRequest)
		return
	}

	respuesta, correcto := bd.LeoTweetsSeguidores(IdUsuario, pagina)

	if !correcto {
		http.Error(w, "Error al leer los tweets ", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)

}
