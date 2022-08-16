package bd

import (
	"github.com/FedeDifulvio/Curso-GO-React/models"
	"golang.org/x/crypto/bcrypt"
)

func IntentoLogin(email string, password string) (models.Usuario, bool) {

	user, existe, _ := ChequeoYaExisteUsuario(email)

	if !existe {
		return user, false
	}

	passwordBytes := []byte(password)

	passwordBD := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		return user, false
	}

	return user, true

}
