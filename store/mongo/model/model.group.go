package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GroupModel struct {
	ID         primitive.ObjectID   `bson:"_id,omitempty"`
	Title      string               `bson:"title"`
	Dependents []primitive.ObjectID `bson:"dependents"`
	RepLink    string               `bson:"rep_link"`
	Version    string               `bson:"version"`
}

type RequestGroupModel struct {
	ID         primitive.ObjectID   `json:"id"`
	Title      string               `json:"title"`
	Dependents []primitive.ObjectID `json:"dependents"`
	RepLink    string               `json:"rep_link"`
	Version    string               `json:"version"`
}
