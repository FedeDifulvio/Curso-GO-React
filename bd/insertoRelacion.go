package bd

import (
	"context"
	"time"

	"github.com/FedeDifulvio/Curso-GO-React/models"
)

func InsertoRelacion(relacion models.Relacion) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("Twittor")

	col := db.Collection("relacion")

	_, err := col.InsertOne(ctx, relacion)

	if err != nil {
		return false, err
	}

	return true, nil
}
