package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//UserClaims ...
type UserClaims struct {
	jwt.StandardClaims
	ID   primitive.ObjectID
	Time time.Time
}

//UserModel модель пользователя в БД
type UserModel struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Login    string             `bson:"login"`
	Password string             `bson:"password"`
}
