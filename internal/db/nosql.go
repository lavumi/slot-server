package db

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"net/url"
)

//const uri = "mongodb+srv://lavumi:<password>@cluster0.kuovpbb.mongodb.net/?retryWrites=true&w=majority"

type MongoDb struct {
	//Cluster  string
	//Username string
	//password string
	client   *mongo.Client
	database *mongo.Database
}

//var client *mongo.Client

func (m *MongoDb) Initialize(cluster string, username string, password string) {

	//cluster := os.Getenv("CLUSTER")
	//username := os.Getenv("USER")
	//password := os.Getenv("PASS")

	uri := "mongodb://" + url.QueryEscape(username) + ":" +
		url.QueryEscape(password) + "@" + cluster +
		"/admin"

	fmt.Println(uri)
	var err error
	m.client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	var result bson.M
	if err := m.client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	m.database = m.client.Database("slot")
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

}

func (m *MongoDb) GetCollection(name string) *mongo.Collection {
	return m.database.Collection(name)
}

func (m *MongoDb) DisConnect() {
	if err := m.client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
