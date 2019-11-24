package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const SOURCES = "Sources"

type Source struct {
	ID             *primitive.ObjectID  `bson:"_id,omitempty"`
	url            string               `json:"url" bson:"url"`
}

