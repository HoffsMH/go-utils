package main

import (
	"fmt"
	"os"

	util "git.mhkr.xyz/go-utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "hsplit",
	Args: cobra.MinimumNArgs(0),
  Run: func(cmd *cobra.Command, args []string) {
    hsplitter := util.NewHsplitter()
    hsplitter.Call(util.StdinLines(), args[0])
  },
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
