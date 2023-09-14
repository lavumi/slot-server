package database

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

var spinCollection *mongo.Collection
var collectCollection *mongo.Collection
var userCollection *mongo.Collection

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

	userCollection = client.Database("slot_data_2").Collection("user")
	spinCollection = client.Database("slot_data_2").Collection("spin")
	collectCollection = client.Database("slot_data_2").Collection("collect")
	//filter := bson.D{{"name", "lavumi2"}}
	//
	//loginCollection := client.Database("user").Collection("login")
	//
	//var user model.Login
	//err = loginCollection.FindOne(context.TODO(), filter).Decode(&user)
	//if err == mongo.ErrNoDocuments {
	//	// Do something when no record was found
	//	fmt.Println("record does not exist")
	//} else if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf("User Query test %s | %s\n", user.Name, user.PlatformId)
	//return client
}

func DisConnect() {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func InsertUserData(user string, count int16) {
	_, err := userCollection.InsertOne(context.TODO(), User{
		Name:  user,
		Count: count,
	})
	if err != nil {
		return
	}
}

func InsertSpinData(spinRawData bson.M) {
	_, err := spinCollection.InsertOne(context.TODO(), spinRawData)
	if err != nil {
		panic(err)
	}
}

func InsertCollectionData(collectRawData bson.M) {
	_, err := collectCollection.InsertOne(context.TODO(), collectRawData)
	if err != nil {
		panic(err)
	}
}

func GetUserList() []User {
	cursor, err := userCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.Background())

	var users []User
	if err = cursor.All(context.TODO(), &users); err != nil {
		panic(err)
	}
	return users
}

func GetSpinDataByUser(userId string) ([]bson.M, error) {
	cursor, err := spinCollection.Find(context.TODO(), bson.D{
		{"uuid", userId},
	})
	if err != nil {
		return nil, err
	}

	var spinResult []bson.M
	if err = cursor.All(context.TODO(), &spinResult); err != nil {
		panic(err)
	}
	return spinResult, nil
}
