package models

import (
	"App/database"

	"github.com/AboutGoods/go-utils/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const USERS = "Users"

type User struct {
	ID         *primitive.ObjectID `bson:"_id,omitempty"`
	Username   string              `json:"username",bson:"username"`
	Password   string              `json:"password",bson:"password"`
	PublicKey  string              `json:"public_key",bson:"public_key"`
	PrivateKey string              `json:"private_key",bson:"private_key"`
}

type UserPatch struct {
	Username *string `json:"username"`
}

func (User *User) ApplyPatch(patch UserPatch) {
	if patch.Username != nil {
		User.Username = *patch.Username
	}
}

func (user *User) Load(id string) error {
	userDB, err := UserRepository{}.FindById(id)
	*user = userDB
	return err
}

func (user *User) Drop() error {
	_, err := UserRepository{}.Collection().DeleteOne(database.Context, bson.M{"_id": user.ID})
	if err == nil {
		user.ID = nil
	}
	return err
}

func (user *User) Store() {
	collection := UserRepository{}.Collection()
	if user.ID == nil {
		response, err := collection.InsertOne(database.Context, user)
		log.PanicOnError(err, "cannot insert user to database")
		id := response.InsertedID.(primitive.ObjectID)
		user.ID = &id
		return
	}
	_, err := collection.UpdateOne(database.Context, bson.M{"_id": *user.ID},
		bson.M{"$set": user})
	log.PanicOnError(err, "cannot update user to database")
}

type UserRepository struct {
}

func (UserRepository) Collection() *mongo.Collection {
	return database.Collection(USERS)
}

func (UserRepository) FindAll() []User {

	cur, _ := database.Collection(USERS).Find(database.Context, bson.D{})
	users := []User{}
	if cur != nil {
		for cur.Next(database.Context) {
			user := User{}
			cur.Decode(&user)
			users = append(users, user)
		}
		cur.Close(database.Context)
	}
	return users

}

func (UserRepository) FindById(id string) (User, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}
	var user User
	err := UserRepository{}.Collection().FindOne(database.Context, filter).Decode(&user)
	return user, err
}
