package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Log struct {
	PhraseID   *primitive.ObjectID  `json:"phraseid" binding:"required"`
	Timestamp  int                  `json:"timestamp" bson:"timestamp" binding:"required"`
	Attempts  int                   `json:"attempts" bson:"attempts" binding:"required"`
}