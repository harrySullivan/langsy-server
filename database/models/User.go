package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
)

const USERS = "User"

type User struct {
	ID             *primitive.ObjectID  `bson:"_id,omitempty"`
	OauthID        string               `json:"oauthid" bson:"oauthid"`
	Avatar         int                  `json:"avatar" bson:"avatar"`
	Username       int                  `json:"username" bson:"username"`
	Courses        []Course             `json:"courses" bson:"courses"`
	Level          int                  `json:"level" bson:"level"`
}

type UserPatch struct {
	Set bson.M `structs:"$set,omitempty"`
	Push bson.M `binding:"required" structs:"$push"`
}

var UserPatchSchema = map[string][]string {
	"$set": {"avatar", "username", "level"},
	"$push": {"courses"},
}

