package models

import (
	"App/database"

	"github.com/AboutGoods/go-utils/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const COURSES = "Courses"

type Course struct {
	ID             *primitive.ObjectID  `bson:"_id,omitempty"`
	Language       string               `json:"language",bson:"language"`
	NextMotivator  int                  `json:"nextmotivator",bson:"nextmotivator"`
	Logs           []Log                `json:"logs",bson:"logs"`
}

type CoursePatch struct {
	NextMotivator *int `json:"NextMotivator"`
	NewLog *Log `json:"NewLog"`
}

func (course *Course) ApplyPatch(patch CoursePatch) {
	if patch.NextMotivator != nil {
		course.NextMotivator = *patch.NextMotivator
	}

	if patch.NewLog != nil {
		course.Logs = append(course.Logs, *patch.NewLog)
	}
}

func (course *Course) Load(id string) error {
	courseDB, err := CourseRepository{}.FindById(id)
	*course = courseDB
	return err
}

func (course *Course) Drop() error {
	_, err := CourseRepository{}.Collection().DeleteOne(database.Context, bson.M{"_id": course.ID})
	if err == nil {
		course.ID = nil
	}
	return err
}

func (course *Course) Store() {
	collection := CourseRepository{}.Collection()
	if course.ID == nil {
		response, err := collection.InsertOne(database.Context, course)
		log.PanicOnError(err, "cannot insert course to database")
		id := response.InsertedID.(primitive.ObjectID)
		course.ID = &id
		return
	}
	_, err := collection.UpdateOne(database.Context, bson.M{"_id": *course.ID},
		bson.M{"$set": course})
	log.PanicOnError(err, "cannot update course to database")
}

type CourseRepository struct {
}

func (CourseRepository) Collection() *mongo.Collection {
	return database.Collection(COURSES)
}

func (CourseRepository) FindAll() []Course {

	cur, _ := database.Collection(COURSES).Find(database.Context, bson.D{})
	courses := []Course{}
	if cur != nil {
		for cur.Next(database.Context) {
			course := Course{}
			cur.Decode(&course)
			courses = append(courses, course)
		}
		cur.Close(database.Context)
	}
	return courses

}

func (CourseRepository) FindById(id string) (Course, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}
	var course Course
	err := CourseRepository{}.Collection().FindOne(database.Context, filter).Decode(&course)
	return course, err
}
