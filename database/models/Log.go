package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Log struct {
	PhraseID   *primitive.ObjectID  `bson:"_id,omitempty"`
	Timestamp  int                  `json:"Timestamp",bson:"Timestamp"`
	Attempts  int                   `json:"Attempts",bson:"Attempts"`
}