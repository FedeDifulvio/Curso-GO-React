package bd

import (
	"context"
	"log"
	"time"

	"github.com/FedeDifulvio/Curso-GO-React/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LeoTweets(ID string, pagina int64) ([]*models.DevuelvoTweets, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("Twittor")

	col := db.Collection("tweet")

	var resultados []*models.DevuelvoTweets

	condicion := bson.M{
		"userid": ID,
	}

	opciones := options.Find()

	opciones.SetLimit(20)

	// el -1 inidica que tiene que traer el listado en orden descendente
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})
	// cuantos registros se saltea. Sirve para la paginación
	opciones.SetSkip((pagina - 1) * 20)

	cursor, err := col.Find(ctx, condicion, opciones)

	if err != nil {
		log.Fatal(err.Error())
		return resultados, false
	}

	//context.TODO() crea un nuevo contexto sin limitaciones

	for cursor.Next(context.TODO()) {

		var registro models.DevuelvoTweets
		err := cursor.Decode(&registro)
		if err != nil {
			return resultados, false
		}

		resultados = append(resultados, &registro)
	}

	return resultados, true
}
