package services

import (
	"App/database"
	"App/database/models"

	"github.com/AboutGoods/go-utils/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PhraseService struct {
}

func (PhraseService) Collection() *mongo.Collection {
	return database.Collection(models.PHRASES)
}

func (PhraseService) FindAll() (*[]models.Phrase, error) {
	phrases := []models.Phrase{}
	cur, err := PhraseService{}.Collection().Find(database.Context, bson.D{})
	err = cur.All(database.Context, &phrases)
	
	log.NoticeOnError(err, "cannot find all phrases in database")

	return &phrases, err

}

func (PhraseService) GetRandom(lang string) (*models.Phrase, error) {
	pipeline := []bson.M{
		bson.M{"$match": bson.M{"lang": lang}},
		bson.M{"$sample": bson.M{"size": 1}},
	}

	cur, err := PhraseService{}.Collection().Aggregate(database.Context, pipeline)
	log.NoticeOnError(err, "cannot find random phrase in database")

	var phrases []models.Phrase
	err = cur.All(database.Context, &phrases)
	log.NoticeOnError(err, "cannot decode random phrase in database")

	if len(phrases) > 0 {
		return &phrases[0], err
	}
	
	return nil, err
}

func (PhraseService) FindById(phraseId *primitive.ObjectID) (*models.Phrase, error) {
	filter := bson.M{"_id": phraseId}
	var phrase models.Phrase
	err := PhraseService{}.Collection().FindOne(database.Context, filter).Decode(&phrase)

	log.NoticeOnError(err, "cannot find phrase in database")

	return &phrase, err
}

func (PhraseService) CountFromLanguage(lang string) (int64, error) {
	filter := bson.M{"lang": lang}
	count, err := PhraseService{}.Collection().CountDocuments(database.Context, filter, nil)

	log.NoticeOnError(err, "cannot find phrase in database")

	return count, err
}

func (PhraseService) Insert(phrase *models.Phrase) (*models.Phrase, error) {
	response, err := PhraseService{}.Collection().InsertOne(database.Context, phrase)

	log.NoticeOnError(err, "cannot insert phrase to database")

	id := response.InsertedID.(primitive.ObjectID)
	phrase.ID = &id
	return phrase, err
}

func (PhraseService) InsertMany(phrases *[]interface{}) error {
	_, err := PhraseService{}.Collection().InsertMany(database.Context, *phrases)

	log.NoticeOnError(err, "cannot insert phrase to database")
	return err
}

func (PhraseService) Delete(phraseId *primitive.ObjectID) error {
	_, err := PhraseService{}.Collection().DeleteOne(database.Context, bson.M{"_id": phraseId})
	log.NoticeOnError(err, "cannot delete phrase to database")
	return err
}