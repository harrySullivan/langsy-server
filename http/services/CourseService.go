package services

import (
	"App/database"
	"App/database/models"

	"github.com/AboutGoods/go-utils/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CourseService struct {
}

func (CourseService) Collection() *mongo.Collection {
	return database.Collection(models.COURSES)
}

func (CourseService) FindAll() (*[]models.Course, error) {
	courses := []models.Course{}
	cur, err := CourseService{}.Collection().Find(database.Context, bson.D{})
	err = cur.All(database.Context, &courses)
	
	log.PanicOnError(err, "cannot find all courses in database")

	return &courses, err

}

func (CourseService) FindById(courseId *primitive.ObjectID) (*models.Course, error) {
	filter := bson.M{"_id": courseId}
	var course models.Course
	err := CourseService{}.Collection().FindOne(database.Context, filter).Decode(&course)

	log.PanicOnError(err, "cannot find course in database")

	return &course, err
}

func (CourseService) Insert(course *models.Course) (*models.Course, error) {
	response, err := CourseService{}.Collection().InsertOne(database.Context, course)

	log.PanicOnError(err, "cannot insert course to database")

	id := response.InsertedID.(primitive.ObjectID)
	course.ID = &id
	return course, err
}

func (CourseService) Update(courseId *primitive.ObjectID, coursePatch *map[string]interface{}) (error) {
	_, err := CourseService{}.Collection().UpdateOne(database.Context, bson.M{"_id": *courseId},
		*coursePatch)

	log.PanicOnError(err, "cannot update course to database")

	return err
}

func (CourseService) Delete(courseId *primitive.ObjectID) error {
	_, err := CourseService{}.Collection().DeleteOne(database.Context, bson.M{"_id": courseId})
	log.PanicOnError(err, "cannot delete course to database")
	return err
}