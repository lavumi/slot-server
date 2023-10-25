package models

type SavedFeature struct {
	UUID        string `bson:"uuid"`
	SaveData    []byte `bson:"save-data"`
	Collectable bool   `bson:"cash"`
}
