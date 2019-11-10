package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Log struct {
	PhraseID   *primitive.ObjectID  `bson:"_id,omitempty"`
	Timestamp  int                  `json:"timestamp",bson:"timestamp"`
	Attempts  int                   `json:"attempts",bson:"attempts"`
}