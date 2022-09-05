package main

import (
	"fmt"
	"git.mhkr.xyz/go-utils"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "append_title",
	Long: "appends html title to links",
	Args:    cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		terms := util.GetTerms(args)

    fmt.Println(util.AppendTitle(terms))
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


