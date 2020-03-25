package bin

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
)

type Vocabulary struct {
	word        string
	wordType    string
	definitions []string
	examples    []string
	related     []string
	Forms       Forms
}

type Forms struct {
	word       string
	nominative string
	genitive   string
	partitive  string
}

func (w Forms) IsEmpty() bool {
	if strings.TrimSpace(w.nominative) == "" {
		return true
	}

	return false
}

func (w Forms) String() string {
	return fmt.Sprintf("Form |> %s - %s - %s", w.nominative, w.genitive, w.partitive)
}

func (w Forms) ToAnki(word string) string {
	if w.IsEmpty() {
		return ""
	}
	return fmt.Sprintf("%s, %s - %s - %s\n", word, w.nominative, w.genitive, w.partitive)
}

func (v Vocabulary) String() string {
	output := ""
	fgRed := color.New(color.FgRed).SprintFunc()
	fgGreen := color.New(color.FgGreen).SprintFunc()

	// output = output + fgGreen(v.word + "\n")
	formattedWord := fmt.Sprintf("%s [%s]", v.word, v.wordType)
	output = output + fgGreen(AddUnderline(formattedWord+"", "="))

	output = output + fgRed(AddUnderline("Definitions:", "-"))
	for ind, definition := range v.definitions {
		output = output + fmt.Sprintf("%d. %s \n", ind, definition)
	}

	output = output + fgRed(AddUnderline("\nExamples:", "-"))
	for ind, definition := range v.examples {
		output = output + fmt.Sprintf("%d. %s \n", ind, definition)
	}

	output = output + fgRed(AddUnderline("\nRelated words:", "-"))
	for ind, definition := range v.related {
		output = output + fmt.Sprintf("%d. %s \n", ind, definition)
	}

	return output
}
