package bd

import (
	"context"
	"log"
	"time"

	"github.com/FedeDifulvio/Curso-GO-React/models"
	"go.mongodb.org/mongo-driver/bson"
)

func LeoTweetsSeguidores(ID string, pagina int) ([]models.DevuelvoTweetsSeguidores, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("Twittor")

	col := db.Collection("relacion")

	skip := (pagina - 1) * 20

	condiciones := make([]bson.M, 0)

	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})

	// hace el join de tablas. El from indica la tabla con la que se 	quiere ewstablecer relación. Se quiere relacionar relaciones con tweets
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "usuariorelacionid",
			"foreignField": "userid",
			"as":           "tweet",
		}})

	// nos permite traer todos los documentos exactamente iguales
	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})

	//condición para que los tweets vengan ordenados por fecha  del mas nuevo al más viejo
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"tweet.fecha": -1}})

	condiciones = append(condiciones, bson.M{"$skip": skip})

	condiciones = append(condiciones, bson.M{"$limit": 20})

	cursor, err := col.Aggregate(ctx, condiciones)

	if err != nil {
		log.Fatal("error al traer los datos: " + err.Error())
		return nil, false
	}

	var result []models.DevuelvoTweetsSeguidores

	err = cursor.All(ctx, &result)

	if err != nil {
		return result, false
	}

	return result, true
}
