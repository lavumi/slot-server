package db

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"net/url"
	"os"
)

//const uri = "mongodb+srv://lavumi:<password>@cluster0.kuovpbb.mongodb.net/?retryWrites=true&w=majority"

type MongoDb struct{}

var client *mongo.Client

func Initialize() {

	cluster := os.Getenv("CLUSTER")
	username := os.Getenv("USER")
	password := os.Getenv("PASS")

	uri := "mongodb://" + url.QueryEscape(username) + ":" +
		url.QueryEscape(password) + "@" + cluster +
		"/admin"

	fmt.Println(uri)
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

}

func DisConnect() {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
