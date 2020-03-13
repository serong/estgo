package cmd

import (
	"estgo/bin"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(playCmd)
}

var playCmd = &cobra.Command{
	Use:   "play [word]",
	Short: "TTS of given word.",
	Long:  `Play the given word using Google Translate TTS`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		bin.PlayMP3(bin.GetTTS(args[0]))
	},
}
