package bin

import (
	"fmt"
	"github.com/gocolly/colly"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetTTS(phrase string) (rawMP3 []byte) {
	url := fmt.Sprintf("https://translate.google.com/translate_tts?ie=UTF-8&tl=et&client=tw-ob&q=%s", phrase)

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	rawMP3, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return rawMP3
}

func FetchForms(word string) (vocabulary Vocabulary) {
	url := fmt.Sprintf("http://www.eki.ee/dict/psv/index.cgi?Q=%s&F=M", word)
	vocabulary.word = word

	var elements []string

	collector := colly.NewCollector()
	collector.OnHTML("span[class=mvf]", func(e *colly.HTMLElement) {
		element := strings.Replace(e.Text, "`", "", -1)
		element = strings.TrimSpace(element)
		elements = append(elements, element)
	})

	collector.Visit(url)

	if len(elements) > 2 {
		vocabulary.Forms.nominative = elements[0]
		vocabulary.Forms.genitive = elements[1]
		vocabulary.Forms.partitive = elements[2]
	}

	return vocabulary
}

func FetchMultipleForms(words []string) (vocabs []Vocabulary) {
	for _, word := range words {
		vocabs = append(vocabs, FetchForms(word))
	}

	return vocabs
}

func FetchDefinitions(word string) (voc Vocabulary) {
	url := fmt.Sprintf("https://www.dict.com/estonian-english/%s", word)

	collector := colly.NewCollector()
	voc.word = word

	// Word itself
	collector.OnHTML("span[class='lex_ful_entr l1']", func(e *colly.HTMLElement) {
		element := strings.TrimSpace(e.Text)
		voc.word = voc.word + " - " + strings.ToUpper(element)
	})

	// Word itself
	collector.OnHTML("span[class='lex_ful_morf']", func(e *colly.HTMLElement) {
		element := strings.TrimSpace(e.Text)
		if voc.wordType == "" {
			voc.wordType = strings.ToLower(element)
		}

	})

	// Definitions
	collector.OnHTML("span[class='lex_ful_tran w l2']", func(e *colly.HTMLElement) {
		element := strings.TrimSpace(e.Text)
		voc.definitions = append(voc.definitions, element)
	})

	// Related
	collector.OnHTML("span[class='lex_ful_coll2']", func(e *colly.HTMLElement) {
		element := strings.TrimSpace(e.Text)
		voc.related = append(voc.related, element)
	})

	// Examples
	collector.OnHTML("span[class='lex_ful_samp2']", func(e *colly.HTMLElement) {
		element := strings.TrimSpace(e.Text)
		voc.examples = append(voc.examples, element)
	})

	collector.Visit(url)

	return voc
}
