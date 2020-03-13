package cmd

import (
	"estgo/bin"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(defCmd)
}

var defCmd = &cobra.Command{
	Use:   "def [word]",
	Short: "Get definition and details of a given word",
	Long:  `Get definition and details of a given word`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		definition := bin.FetchDefinitions(args[0])
		fmt.Println(definition)
	},
}
