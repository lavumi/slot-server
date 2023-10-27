package models

type SavedFeature struct {
	UUID        string `bson:"uuid"`
	SaveData    []byte `bson:"save"`
	Collectable bool   `bson:"c"`
}
