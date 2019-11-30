package services

import (
	"App/database"
	"App/database/models"

	"github.com/AboutGoods/go-utils/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SourceService struct {
}

func (SourceService) Collection() *mongo.Collection {
	return database.Collection(models.SOURCES)
}

func (SourceService) FindAll() (*[]models.Source, error) {
	sources := []models.Source{}
	cur, err := SourceService{}.Collection().Find(database.Context, bson.D{})
	err = cur.All(database.Context, &sources)
	
	log.NoticeOnError(err, "cannot find all sources in database")

	return &sources, err
}

func (SourceService) FindByUrl(url string) (*models.Source, error) {
	filter := bson.M{"url": url}
	var source models.Source
	err := SourceService{}.Collection().FindOne(database.Context, filter).Decode(&source)

	log.NoticeOnError(err, "cannot find source in database")

	return &source, err
}

func (SourceService) Insert(source *models.Source) (*models.Source, error) {
	response, err := SourceService{}.Collection().InsertOne(database.Context, source)

	log.NoticeOnError(err, "cannot insert source to database")

	id := response.InsertedID.(primitive.ObjectID)
	source.ID = &id
	return source, err
}