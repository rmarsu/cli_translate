package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var translatorCmd = &cobra.Command{
	Use:   "translator",
	Short: "Перевод текста , языки писать в 2 буквенном формате.",
	Long:  `Переводит текст. Принимает только 2 аргументa. Языки писать в 2 буквенном формате`,
	Run: func(cmd *cobra.Command, args []string) {
		var textval string
		var langval string

		text, _ := cmd.Flags().GetString("text")

		if text != "" {
			textval = text
		} else {
			textval = "You must specify the text to translate"
			os.Exit(1)
		}

		lang, _ := cmd.Flags().GetString("lang")
		if lang != "" {
			langval = lang
		} else {
			textval = "You must specify language"
		}
		if textval != "" && langval != "" {
			fmt.Println(translatetxt(textval, langval))
		}
	},
}

func init() {
	rootCmd.AddCommand(translatorCmd)
	translatorCmd.Flags().StringP("text", "t", "", "The text to translate")
	translatorCmd.Flags().StringP("lang", "l", "", "Language")
}
