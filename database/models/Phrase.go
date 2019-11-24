package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const PHRASES = "Phrases"

type Phrase struct {
	ID             *primitive.ObjectID  `bson:"_id,omitempty"`
	content        string               `json:"content" bson:"content"`
	Souce          *primitive.ObjectID  `json:"source" bson:"source"`
}

