package routers

import (
	"encoding/json"
	"net/http"

	"github.com/FedeDifulvio/Curso-GO-React/bd"
	"github.com/FedeDifulvio/Curso-GO-React/jwt"
	"github.com/FedeDifulvio/Curso-GO-React/models"
)

func Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("content-type", "application/json")

	var user models.Usuario

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Usuario y/ contraseña inválidos:"+err.Error(), 400)
		return
	}
	if len(user.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}

	documento, existe := bd.IntentoLogin(user.Email, user.Password)

	if !existe {
		http.Error(w, "Usuario y/ contraseña inválidos", 400)
		return
	}

	jwtkey, err := jwt.GeneroJWT(documento)

	if err != nil {
		http.Error(w, "Error al generar token:"+err.Error(), 400)
	}

	resp := models.RespuestaLogin{
		Token: jwtkey,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

}
