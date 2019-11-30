package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
)

const COURSES = "Courses"

type Course struct {
	ID             *primitive.ObjectID  `bson:"_id,omitempty"`
	Language       string               `json:"language" bson:"language"`
	NextReward     int                  `json:"nextreward" bson:"nextreward"`
	Score          int                  `json:"score" bson:"score"`
	Logs           []CourseLog          `json:"logs" bson:"logs"`
	Completed      bool                 `json:"completed" bson:"completed"`
}

type CoursePatch struct {
	Set bson.M `structs:"$set,omitempty"`
	Push bson.M `binding:"required" structs:"$push"`
}

var CoursePatchSchema = map[string][]string {
	"$set": {"nextreward", "score", "completed"},
	"$push": {"logs"},
}

