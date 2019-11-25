package services

import (
	"App/database"
	"App/database/models"

	"github.com/AboutGoods/go-utils/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
}

func (UserService) Collection() *mongo.Collection {
	return database.Collection(models.USERS)
}

func (UserService) FindAll() (*[]models.User, error) {
	users := []models.User{}
	cur, err := UserService{}.Collection().Find(database.Context, bson.D{})
	err = cur.All(database.Context, &users)
	
	log.PanicOnError(err, "cannot find all users in database")

	return &users, err

}

func (UserService) FindById(userId *primitive.ObjectID) (*models.User, error) {
	filter := bson.M{"_id": userId}
	var user models.User
	err := UserService{}.Collection().FindOne(database.Context, filter).Decode(&user)

	log.PanicOnError(err, "cannot find user in database")

	return &user, err
}

func (UserService) Insert(user *models.User) (*models.User, error) {
	response, err := UserService{}.Collection().InsertOne(database.Context, user)

	log.PanicOnError(err, "cannot insert user to database")

	id := response.InsertedID.(primitive.ObjectID)
	user.ID = &id
	return user, err
}

func (UserService) Update(userId *primitive.ObjectID, userPatch *map[string]interface{}) (error) {
	_, err := UserService{}.Collection().UpdateOne(database.Context, bson.M{"_id": *userId},
		*userPatch)

	log.PanicOnError(err, "cannot update user to database")

	return err
}

func (UserService) Delete(userId *primitive.ObjectID) error {
	_, err := UserService{}.Collection().DeleteOne(database.Context, bson.M{"_id": userId})
	log.PanicOnError(err, "cannot delete user to database")
	return err
}