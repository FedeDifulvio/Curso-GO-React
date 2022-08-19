package routers

import (
	"encoding/json"
	"net/http"

	"github.com/FedeDifulvio/Curso-GO-React/bd"
	"github.com/FedeDifulvio/Curso-GO-React/models"
)

func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

	var user models.Usuario

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Error en el formato de los datos:"+err.Error(), 400)
		return
	}

	status, error := bd.ModificoRegistro(user, IdUsuario)

	if !status {
		http.Error(w, "Error al querer modificar perfil:"+error.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
