package main

import (
	"fmt"
	"os"

	util "git.mhkr.xyz/go-utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "hcat",
	Args: cobra.MinimumNArgs(0),
  Run: func(cmd *cobra.Command, args []string) {
		terms := util.StdinLines()

    if len(args) == 0 {
      wd, _ := os.Getwd()
      fmt.Print(util.Hcat(terms, wd))
    } else {
      fmt.Print(util.Hcat(terms, args[0]))
    }
  },
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
