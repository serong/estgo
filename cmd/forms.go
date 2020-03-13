package cmd

import (
	"estgo/bin"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var fromFile bool

func init() {
	rootCmd.AddCommand(formsCmd)
	formsCmd.Flags().BoolVarP(
		&fromFile, "file", "f", false,
		"Fetch list of words from a file.")
}

var formsCmd = &cobra.Command{
	Use:   "forms [word or filename]",
	Short: "Get first 3 noun cases of given word(s)",
	Long:  `Get nominative, genitive, and partitive cases for given word(s) from an online dictionary.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fgGreen := color.New(color.FgGreen).SprintFunc()
		fgRed := color.New(color.FgRed).SprintFunc()

		fmt.Println(fgRed(">>> Fetching word forms: "))
		defer fmt.Println(fgGreen(">>> Done !"))

		if fromFile != true {
			vocabulary := bin.FetchForms(args[0])
			fmt.Println(">>>", vocabulary.Forms.String())
		} else {
			vocabularies := bin.FetchMultipleForms(bin.ParseWordsFromFile(args[0]))

			for _, vocabulary := range vocabularies {
				fmt.Println(">>>", vocabulary.Forms.String())
			}

			bin.WriteForms("anki-forms.txt", vocabularies)
		}
	},
}
