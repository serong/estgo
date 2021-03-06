package bin

import (
	"io/ioutil"
	"log"
	"strings"
)

func ParseWordsFromFile(filePath string) (words []string) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(content), "\n")
}

func WriteForms(filePath string, vocabs []Vocabulary) {
	var str string

	for _, vocab := range vocabs {
		str = str + vocab.Forms.ToAnki(vocab.word)
	}

	err := ioutil.WriteFile(filePath, []byte(str), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
