package slot

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"net/url"
	"os"
)

type Database struct {
	client *mongo.Client
	log    *mongo.Collection
	state  *mongo.Collection
}

func Initialize() *Database {
	cluster := os.Getenv("CLUSTER")
	username := os.Getenv("SLOT_DB_USER")
	password := os.Getenv("SLOT_DB_PASS")
	db := os.Getenv("SLOT_DB_NAME")

	uri := "mongodb://" + url.QueryEscape(username) + ":" +
		url.QueryEscape(password) + "@" + cluster +
		"/admin"

	fmt.Println(uri)
	//var err error
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	logCol := client.Database(db).Collection("log")
	stateCol := client.Database("slot_data_2").Collection("spin")

	slotDb := Database{
		client: client,
		log:    logCol,
		state:  stateCol,
	}

	return &slotDb
}
