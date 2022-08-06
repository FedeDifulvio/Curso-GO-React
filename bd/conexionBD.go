package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConectarBD()

var ClientOptions = options.Client().ApplyURI("mongodb+srv://fdmongouser:bCCFrbKjexubuIFQ@cluster0.oov8a.mongodb.net/test?authSource=admin&replicaSet=atlas-qc5aau-shard-0&readPreference=primary&appname=MongoDB%20Compass&ssl=true")

/*Conecta con MongoDB*/

func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), ClientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Printf("Conexión existosa con BD")
	return client

}

func ChequeoConexion() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return 0
	}
	return 1
}
