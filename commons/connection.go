package commons

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

  var collection *mongo.Collection
  var ctx = context.TODO() //Añadir contexto en futuro

 func initDDBB() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://mariogl:mariogl@cluster0.hlmfl.mongodb.net/pallars")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
  err = client.Ping(ctx, nil)
  if err != nil {
    log.Fatal(err)
  }
  collection = client.Database("pallars").Collection("viatges")
}
