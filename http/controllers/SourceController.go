package controllers

import (
	"App/database/models"
	"App/http/services"
	"App/http/processes"
	"App/utils"

	"net/url"
	"regexp"
	"strings"
	"fmt"

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

func cleanPhrases(strs []string) []string {
	var final []string
	for _, str := range strs {
		re := regexp.MustCompile(`\[.*\]`)
		noBrackets := re.ReplaceAllString(str, "")
		trimmed := strings.Trim(noBrackets, " ")
		if len(trimmed) > 0 {
			final = append(final, trimmed)
		}
	}

	return final
}

func convStringsToPhrases(phraseStrs []string, translationStrs []string, lang string) *[]interface{} {
	cleanPhrasesStrs := cleanPhrases(phraseStrs)
	cleanTranslationStrs := cleanPhrases(translationStrs)

	fmt.Println(cleanPhrasesStrs, "\n\n\n\n")
	fmt.Println(cleanTranslationStrs)

	var final []interface{}
	
	if len(cleanPhrasesStrs) == len(cleanTranslationStrs) {
		for i := 0; i < len(cleanPhrasesStrs); i++ {
			phrase := models.Phrase {
				Content: cleanPhrasesStrs[i],
				Translation: cleanTranslationStrs[i],
				Lang: lang,
			}
			final = append(final, phrase)
		}
	} else {
		for i := 0; i < len(cleanPhrasesStrs); i++ {
			phrase := models.Phrase {
				Content: cleanPhrasesStrs[i],
				Translation: "",
				Lang: lang,
			}
			final = append(final, phrase)
		}
	}

	return &final
}

func AddPhrasesAndSource(lang string, c *gin.Context) error {
	translationExists := false

	var targetLangPhrases []string
	var translationPhrases []string
	var articleTitle string

	for ok := true; ok; ok = !translationExists {
		randomUrl, subdomain, domain := processes.ConstructRandomUrl(lang)

		redirectResp, rErr := processes.HttpGETNoRedirect(randomUrl)
		utils.HttpError(c, 500, rErr)

		redirectLocation := redirectResp["Location"][0]

		redirectUrl, pErr := url.Parse(redirectLocation)
		utils.HttpError(c, 500, pErr)

		articleTitle = strings.Replace(redirectUrl.Path, "/wiki/", "", 1)

		articleApiUrl := processes.ConstructArticleUrl(subdomain, domain, articleTitle)
		phraseStrs := processes.GetPhrasesFromPage(articleApiUrl)

		translationLocation, exists := processes.GetTranslationLocation(subdomain, domain, articleTitle)
		translationExists = exists
		translationUrl, pErr := url.Parse(translationLocation)
		translationTitle := strings.Replace(translationUrl.Path, "/wiki/", "", 1)
		translationApiUrl := processes.ConstructArticleUrl("en", domain, translationTitle)
		translationPhraseStrs := processes.GetPhrasesFromPage(translationApiUrl)

		fmt.Println(translationApiUrl)

		targetLangPhrases = phraseStrs
		translationPhrases = translationPhraseStrs
	}


	phrases := convStringsToPhrases(targetLangPhrases, translationPhrases, lang)

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
