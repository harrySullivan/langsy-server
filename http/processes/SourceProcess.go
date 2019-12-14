package processes

import (
	"App/database/models"
	"time"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"math/rand"
	"strings"
	"html"
	"github.com/dlclark/regexp2"
	"github.com/AboutGoods/go-utils/log"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

var LangData []models.LangData

func PreloadLangData() {
	data, err := ioutil.ReadFile("http/controllers/lang-codes/languages-with-random.json")
	if err != nil {
		panic(err.Error())
	}

	json.Unmarshal(data, &LangData)
}

func FindLangData(lang string) *models.LangData {
	for _, langData := range LangData {
		if langData.LangCode == lang {
			return &langData
		}
	}

	return nil
}

func RandomDomain() string {
	domains := []string{"wikipedia", "wikinews"}
	rand.Seed(time.Now().UnixNano())
	return domains[rand.Intn(len(domains))]
}

func ConstructRandomUrl(lang string) (string, string, string) {
	langData := FindLangData(lang)

	subdomain := langData.LangCode
	domain := RandomDomain()
	path := langData.RandLink

	return fmt.Sprintf("https://%s.%s.org%s", subdomain, domain, path), subdomain, domain
}

func HttpGETNoRedirect(url string) (http.Header, error) {
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

func httpGet(url string) ([]byte) {
	resp, err := http.Get(url)
	log.NoticeOnError(err, "Can't get")
	defer resp.Body.Close()
	
	body, rErr := ioutil.ReadAll(resp.Body)
	log.NoticeOnError(rErr, "Can't read body")
	return body
}

func parsedToDocument(page parsedPage) *goquery.Document {
	content := page.Parse.Text.All
	unescapedContent := html.UnescapeString(content)

	formatted := fmt.Sprintf(`<html><head></head><body>%s</body></html>`, unescapedContent)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(formatted))
	log.NoticeOnError(err, "Can't parse")
	return doc
}

func selectionToPhrases(s *goquery.Selection) ([]string, string) {
	text := s.Text()
	noNewLine := strings.Replace(text, "\n", "", -1)
	trimmed := strings.Trim(noNewLine, " ")

	re := regexp2.MustCompile(`(?<!\w\.\w.)(?<![A-Z][a-z]\.)(?<=\.|\?)\s`, regexp2.Multiline)
	phraseSplitByNewLine, err := re.Replace(trimmed, "\n", -1, -1)
	log.NoticeOnError(err, "Can't build sentence split regex")
	chunks := strings.Split(phraseSplitByNewLine, "\n")

	return chunks, noNewLine
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

func GetPhrasesFromPage(subdomain string, domain string, page string) []string {
	url := fmt.Sprintf("https://%s.%s.org/w/api.php?action=parse&format=json&page=%s", subdomain, domain, page)
	body := httpGet(url)
	var parsed parsedPage
	json.Unmarshal(body, &parsed)

	doc := parsedToDocument(parsed)

	var phrases []string
	done := false

	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		if !done {
			if s.Find("br").Length() > 0 {
				done = true
			}
			chunks, noNewLine := selectionToPhrases(s)
			if len(noNewLine) > 0 {
				phrases = append(phrases, chunks...)
			}
		}
	})

	return phrases
}