package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const PHRASES = "Phrases"

type Phrase struct {
	ID             *primitive.ObjectID  `bson:"_id,omitempty"`
	Content        string               `json:"content" bson:"content"`
	Lang           string               `json:"lang" bson:"lang"`
}

