package controllers

import (
	"App/database/models"
	"App/http/services"
	"App/http/processes"
	"App/utils"

	"net/url"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SourceController struct {
	Controller
}

func (SourceController) List(c *gin.Context) {
	sources, err := services.SourceService{}.FindAll()
	if utils.HttpError(c, 500, err) { return }

	c.JSON(200, sources)
}

func (SourceController) Get(c *gin.Context) {
	sourceId := c.Param("sourceId")
	sourceOid, err := primitive.ObjectIDFromHex(sourceId)
	if utils.HttpError(c, 400, err) { return }

	source, err := services.PhraseService{}.FindById(&sourceOid)
	if utils.HttpError(c, 500, err) { return }

	c.JSON(200, source)
}

type addSourceReq struct {
	Language string `json:"language"`
}

func convStringsToPhrases(phraseStrs []string, lang string) *[]interface{} {
	var final []interface{}
	for _, str := range phraseStrs {
		re := regexp.MustCompile(`\[.*\]`)
		noBrackets := re.ReplaceAllString(str, "")
		trimmed := strings.Trim(noBrackets, " ")
		phrase := models.Phrase {
			Content: trimmed,
			Lang: lang,
		}
		if len(trimmed) > 0 {
			final = append(final, phrase)
		}
		
	}

	return &final
}

func AddPhrasesAndSource(lang string, c *gin.Context) error {
	randomUrl, subdomain, domain := processes.ConstructRandomUrl(lang)

	redirectResp, rErr := processes.HttpGETNoRedirect(randomUrl)
	utils.HttpError(c, 500, rErr)

	redirectLocation := redirectResp["Location"][0]

	redirectUrl, pErr := url.Parse(redirectLocation)
	utils.HttpError(c, 500, pErr)

	articleTitle := strings.Replace(redirectUrl.Path, "/wiki/", "", 1)

	phraseStrs := processes.GetPhrasesFromPage(subdomain, domain, articleTitle)

	phrases := convStringsToPhrases(phraseStrs, lang)

	PhraseController{}.InsertMany(phrases)

	source := models.Source {
		Url: articleTitle,
	}

	_, err := services.SourceService{}.Insert(&source)
	return err
}

func (SourceController) Post(c *gin.Context) {
	var req addSourceReq
	err := c.Bind(&req)
	if utils.HttpError(c, 400, err) { return }

	lang := req.Language

	err = AddPhrasesAndSource(lang, c)
	if utils.HttpError(c, 500, err) { return }

	c.JSON(200, "done")
}
