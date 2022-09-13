package bd

import (
	"context"
	"time"

	"github.com/FedeDifulvio/Curso-GO-React/models"
)

func BorroRelacion(relacion models.Relacion) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("Twittor")

	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, relacion)

	if err != nil {
		return false, err
	}

	return true, err
}
