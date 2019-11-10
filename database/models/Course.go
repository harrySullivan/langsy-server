package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const COURSES = "Courses"

type Course struct {
	ID             *primitive.ObjectID  `bson:"_id,omitempty"`
	Language       string               `json:"language",bson:"language"`
	NextMotivator  int                  `json:"nextmotivator",bson:"nextmotivator"`
	Score          int                  `json:"score",bson:"score"`
	Logs           []Log                `json:"logs",bson:"logs"`
}

type CoursePatch struct {
	NextMotivator  *int  `json:"nextmotivator,omitempty"`
	Score          *int  `json:"newscore,omitempty"`
	NewLog         *Log  `json:"newlog,omitempty"`
}

