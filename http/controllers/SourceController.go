// import (
// 	"encoding/xml"
// 	"fmt"
// )

// type div struct {
// 	XMLName  xml.Name `xml:"div"`
// 	PList []p `xml:"p"`
// }

// type p struct {
// 	Text string `xml:",chardata"`
// }

// func main() {
// 	d := div{}
// 	err := xml.Unmarshal(xmlContent, &d)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(d)
// }

package controllers

import (
	"App/database/models"
	"App/http/services"
	"App/utils"

	"io/ioutil"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"math/rand"
	"regexp"
	"html"
	"time"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/AboutGoods/go-utils/log"
	"github.com/PuerkitoBio/goquery"
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

var LangData []models.LangData

func PreloadLangData() {
	data, err := ioutil.ReadFile("http/controllers/lang-codes/languages-with-random.json")
	if err != nil {
		panic(err.Error())
	}

	json.Unmarshal(data, &LangData)
}

func findLangData(lang string) *models.LangData {
	for _, langData := range LangData {
		if langData.LangCode == lang {
			return &langData
		}
	}

	return nil
}

func randomDomain() string {
	domains := []string{"wikipedia", "wikinews"}
	rand.Seed(time.Now().UnixNano())
	return domains[rand.Intn(len(domains))]
}

func constructRandomUrl(lang string) (string, string, string) {
	langData := findLangData(lang)

	subdomain := langData.LangCode
	domain := randomDomain()
	path := langData.RandLink

	return fmt.Sprintf("https://%s.%s.org%s", subdomain, domain, path), subdomain, domain
}

func httpGETNoRedirect(url string) (http.Header, error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Get(url)
	defer resp.Body.Close()

	log.NoticeOnError(err, "Can't get")
	return resp.Header, err
}

type parsedPage struct {
	Parse parseResut `json:"parse"`
}

type parseResut struct {
	Text textResult `json:"text"`
}

type textResult struct {
	All string `json:"*"`
}

func getPhrasesFromPage(subdomain string, domain string, page string) []string {
	url := fmt.Sprintf("https://%s.%s.org/w/api.php?action=parse&format=json&page=%s", subdomain, domain, page)
	resp, err := http.Get(url)
	log.NoticeOnError(err, "Can't get")
	defer resp.Body.Close()
	
	body, rErr := ioutil.ReadAll(resp.Body)
	log.NoticeOnError(rErr, "Can't read body")

	var parsed parsedPage
	json.Unmarshal(body, &parsed)

	content := parsed.Parse.Text.All
	unescapedContent := html.UnescapeString(content)

	formatted := fmt.Sprintf(`<html><head></head><body>%s</body></html>`, unescapedContent)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(formatted))
	log.NoticeOnError(rErr, "Can't parse")

	var phrases []string
	var done bool = false

	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		if !done {
			if s.Find("br").Length() > 0 {
				done = true
			}
			text := s.Text()
			noNewLine := strings.Replace(text, "\n", "", -1)
			trimmed := strings.Trim(noNewLine, " ")

			chunks := regexp.MustCompile(`(\,|\.|\;)`).Split(trimmed, -1)
			if len(noNewLine) > 0 {
				phrases = append(phrases, chunks...)
			}
		}
	})

	return phrases
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

func (SourceController) Post(c *gin.Context) {
	var req addSourceReq
	err := c.Bind(&req)
	if utils.HttpError(c, 400, err) { return }

	lang := req.Language

	randomUrl, subdomain, domain := constructRandomUrl(lang)

	redirectResp, rErr := httpGETNoRedirect(randomUrl)
	if utils.HttpError(c, 500, rErr) { return }

	redirectLocation := redirectResp["Location"][0]

	redirectUrl, pErr := url.Parse(redirectLocation)
	if utils.HttpError(c, 500, pErr) { return }

	articleTitle := strings.Replace(redirectUrl.Path, "/wiki/", "", 1)

	phraseStrs := getPhrasesFromPage(subdomain, domain, articleTitle)

	phrases := convStringsToPhrases(phraseStrs, lang)

	PhraseController{}.InsertMany(phrases)

	c.JSON(200, phrases)
}
