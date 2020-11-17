package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email       string             `json:"email"`
	NickName    string             `json:"nickname"`
	PasswordMD5 string             `json:"-" bson:"password_md5"`
}
