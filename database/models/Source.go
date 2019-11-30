package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const SOURCES = "Sources"

type Source struct {
	ID             *primitive.ObjectID  `bson:"_id,omitempty"`
	Url            string               `json:"url" bson:"url"`
}

type LangData struct {
	LangCode       string `json:"LangCode"`
	Language       string `json:"Language"`
	LanguageLocal  string `json:"LanguageLocal"`
	RandLink       string `json:"RandLink"`
}

