package forms

type ReqGuest struct {
}

type ResGuest struct {
	Key string `bson:"key"`
	Id  string `bson:"id"`
}
