package controllers

import (
	"App/database/models"
	"App/http/services"
	"App/utils"

	"github.com/gin-gonic/gin"
	"github.com/AboutGoods/go-utils/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PhraseController struct {
	Controller
}

func (PhraseController) List(c *gin.Context) {
	phrases, err := services.PhraseService{}.FindAll()
	if utils.HttpError(c, 500, err) { return }

	c.JSON(200, phrases)
}

func (PhraseController) Get(c *gin.Context) {
	phraseId := c.Param("phraseId")
	phraseOid, err := primitive.ObjectIDFromHex(phraseId)
	if utils.HttpError(c, 400, err) { return }

	phrase, err := services.PhraseService{}.FindById(&phraseOid)
	if utils.HttpError(c, 500, err) { return }

	c.JSON(200, phrase)
}

func (PhraseController) GetRandom(c *gin.Context) {
	lang := c.Param("lang")

	phrase, err := services.PhraseService{}.GetRandom(lang)
	if utils.HttpError(c, 500, err) { return }

	c.JSON(200, phrase)
}

func (PhraseController) Post(c *gin.Context) {
	var phrase models.Phrase
	err := c.Bind(&phrase)
	if utils.HttpError(c, 400, err) { return }

	iPhrase, iErr := services.PhraseService{}.Insert(&phrase)
	if utils.HttpError(c, 500, iErr) { return }

	c.JSON(200, iPhrase)
}

func (PhraseController) InsertMany(phrases *[]interface{}) {
	err := services.PhraseService{}.InsertMany(phrases)
	log.NoticeOnError(err, "cannot insert phrases to database")
}

func (PhraseController) Delete(c *gin.Context) {
	phraseId := c.Param("phraseId")
	phraseOid, iErr := primitive.ObjectIDFromHex(phraseId)
	if utils.HttpError(c, 400, iErr) { return }

	phrase, fErr := services.PhraseService{}.FindById(&phraseOid)
	if utils.HttpError(c, 500, fErr) { return }

	count, cErr := services.PhraseService{}.CountFromLanguage(phrase.Lang)
	if utils.HttpError(c, 500, cErr) { return }

	if count < 10 {
		AddPhrasesAndSource(phrase.Lang, c)
	}

	dErr := services.PhraseService{}.Delete(&phraseOid)
	if utils.HttpError(c, 500, dErr) { return }

	c.JSON(200, "good")
}

