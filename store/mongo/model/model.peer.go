package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//UserModel модель пользователя в БД
type PeerModel struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Title    string             `bson:"title"`
	Host     string             `bson:"host"`
	Port     int32              `bson:"port"`
	Login    string             `bson:"login"`
	Password string             `bson:"password"`
}

type FilterPeerModel struct {
	Limit  int64  `json:"limit"`
	Offset int64  `json:"offset"`
	Title  string `json:"title"`
	Host   string `json:"host"`
	Port   int32  `json:"port"`
}

//UserModel модель пользователя в БД
type RequestPeerModel struct {
	ID       primitive.ObjectID `json:"id"`
	Title    string             `json:"title"`
	Host     string             `json:"host"`
	Port     int32              `json:"port"`
	Login    string             `json:"login"`
	Password string             `json:"password"`
}
