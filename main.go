package main

import (
	"log"

	"github.com/FedeDifulvio/Curso-GO-React/bd"
	"github.com/FedeDifulvio/Curso-GO-React/handlers"
)

func main() {
	if bd.ChequeoConexion() == 0 {
		log.Fatal("Sin Conexión a BD")
		return
	}

	handlers.Manejadores()

}
