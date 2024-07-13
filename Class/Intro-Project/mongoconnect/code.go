package mongoconnect

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	//? url neveshto
	var ConnectUrl = options.Client().ApplyURI("mongodb://127.0.0.1:27018")

	// ? COnnect click chido
	var Client,_ = mongo.Connect(context.TODO(), ConnectUrl)
	return Client
}