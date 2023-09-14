package database

type User struct {
	Name  string `bson:"name"`
	Count int16  `bson:"count"`
}

type SpinResult struct {
}
