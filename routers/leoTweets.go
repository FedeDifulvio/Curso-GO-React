package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/FedeDifulvio/Curso-GO-React/bd"
)

func LeoTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro id ", http.StatusBadRequest)
	}

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parámetro página", http.StatusBadRequest)
	}

	// convierte pagina a int
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))

	if err != nil {
		http.Error(w, "Debe enviar el parámetro página con valor mayor a cero", http.StatusBadRequest)
	}

	// convierte pagina en int 64
	pag := int64(pagina)

	respuesta, correcto := bd.LeoTweets(ID, pag)

	if !correcto {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)

}
