package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/FedeDifulvio/Curso-GO-React/bd"
	"github.com/FedeDifulvio/Curso-GO-React/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {

	var mensaje models.Tweet

	err := json.NewDecoder(r.Body).Decode(&mensaje)

	if err != nil {
		http.Error(w, "Error en el formato del mensaje: "+err.Error(), 400)
		return
	}

	registro := models.GraboTweet{
		UserID:  IdUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, _, err = bd.InsertoTweet(registro)

	if err != nil {
		http.Error(w, "Error al insertar tweet: "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
