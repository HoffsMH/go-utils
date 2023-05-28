package main

import (
	"fmt"
	"os"

	util "git.mhkr.xyz/go-utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "append_title",
	Long: "appends html title to links",
	Args:    cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		terms := util.GetTerms(args)
    title_appender := &util.HtmlTitleAppender{
      TitleGetter: &util.GetTitleWrapper{},
    }

    fmt.Println(title_appender.Call(terms))
  },
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
}

