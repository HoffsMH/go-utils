package main

import (
	"fmt"
	"os"

	util "git.mhkr.xyz/go-utils"
	"github.com/spf13/cobra"
)

var before string
var after string

var rootCmd = &cobra.Command{
	Use: "dedup",
	Long: "dedups all lines that match patterns",
	Args:    cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
    util.PrinList(util.DeDup(util.Stdin(), before, args[0]))
  },
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
  rootCmd.Flags().StringVarP(&before, "before", "b", "", "remove preceding line if it matches")
  rootCmd.Flags().StringVarP(&after, "after", "a", "", "remove after line if it matches")
}



