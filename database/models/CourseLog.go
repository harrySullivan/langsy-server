package models

type CourseLog struct {
	Timestamp  int    `json:"timestamp" bson:"timestamp" binding:"required"`
	Attempts   int    `json:"attempts" bson:"attempts" binding:"required"`
}