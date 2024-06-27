package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type friendStruct struct {
	FriendName    string
	FriendSurname string
	FriendPhone   string
}

func main() {
	
	GetUser()

}
func AddUser()  {
	//? url neveshto
	var ConnectUrl = options.Client().ApplyURI("mongodb://127.0.0.1:27018")

	// ? COnnect click chido
	var Client,_ = mongo.Connect(context.TODO(), ConnectUrl)
	
	// !===========
	// ? Ar basayat collection  dedow
	var connection = Client.Database("MasrurDB").Collection("Friends")

	var SomeFriendTemp = friendStruct{
		FriendName:    "Dean",
		FriendSurname: "Winchester",
		FriendPhone:   "+992934567123",
	}

	// ? User ar dasa dob chid
	connection.InsertOne(context.TODO(),SomeFriendTemp)
}

func GetUser()  {
	//? url neveshto
	var ConnectUrl = options.Client().ApplyURI("mongodb://127.0.0.1:27018")

	// ? COnnect click chido
	var Client,_ = mongo.Connect(context.TODO(), ConnectUrl)
	
	// !===========
	// ? Ar basayat collection  dedow
	var connection = Client.Database("MasrurDB").Collection("Friends")

	

	// ? User ar dasa dob chid
	ResultFromDB := connection.FindOne(context.TODO(),bson.M{
		"friendname":"Dean",
	})
	var DataFromDB friendStruct
	ResultFromDB.Decode(&DataFromDB)

	fmt.Printf("DataFromDB.FriendPhone: %v\n", DataFromDB.FriendPhone)
}