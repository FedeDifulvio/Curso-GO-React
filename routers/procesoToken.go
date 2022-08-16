package routers

import (
	"errors"
	"strings"

	"github.com/FedeDifulvio/Curso-GO-React/bd"
	"github.com/FedeDifulvio/Curso-GO-React/models"
	jwt "github.com/dgrijalva/jwt-go"
)

var Email string

var IdUsuario string

func ProcesoToken(tk string) (*models.Claim, bool, string, error) {

	clave := []byte("mastersdeldesarrollopassword")

	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) { return clave, nil })

	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado {

			Email = claims.Email
			IdUsuario = claims.ID.Hex()
		}

		return claims, encontrado, IdUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token inválido")
	}

	return claims, false, string(""), err

}
