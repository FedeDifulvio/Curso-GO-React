package jwt

import (
	"time"

	"github.com/FedeDifulvio/Curso-GO-React/models"
	jwt "github.com/dgrijalva/jwt-go"
)

func GeneroJWT(user models.Usuario) (string, error) {

	clave := []byte("mastersdeldesarrollopassword")

	payload := jwt.MapClaims{
		"email":            user.Email,
		"nombre":           user.Nombre,
		"apellidos":        user.Apellidos,
		"fecha_nacimiento": user.FechaNacimiento,
		"biografia":        user.Biografia,
		"ubicacion":        user.Ubicacion,
		"sitioweb":         user.SitioWeb,
		"_id":              user.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(clave)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
