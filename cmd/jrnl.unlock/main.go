package main

import (
	"fmt"
	"os"

	util "git.mhkr.xyz/go-utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "jrnl.unlock",
	Args: cobra.MinimumNArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
		terms := util.GetTerms(args)
		util.JrnlUnlock(terms[0])
  },
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
