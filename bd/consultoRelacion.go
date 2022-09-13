package bd

import (
	"context"
	"time"

	"github.com/FedeDifulvio/Curso-GO-React/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ConsultoRelacion(relacion models.Relacion) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("Twittor")

	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         relacion.UsuarioID,
		"usuariorelacionid": relacion.UsuarioRelacionID,
	}

	var resultado models.Relacion

	err := col.FindOne(ctx, condicion).Decode(&resultado)

	if err != nil {
		return false, err
	}

	return true, nil
}
