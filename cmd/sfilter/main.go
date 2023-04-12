package main

import (
	"fmt"
	"os"

	util "git.mhkr.xyz/go-utils"
	"github.com/spf13/cobra"
)

var days int
var weeks int
var months int
var ignore bool
var reject bool

var rootCmd = &cobra.Command{
	Use: "sfilter",
	Long: "outputs a file name to std out with a date prefix if it does not already have one",
	Args:    cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		terms := util.GetTerms(args)

    util.PrintList(util.SFilter(terms, tr, ignore, reject))
  },
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}


func init() {
  rootCmd.Flags().IntVarP(&lessThan, "less-than", "lt", 2, "the amount of days to look back.")
  rootCmd.Flags().IntVarP(&greaterThan, "greater-than", "gt", 0, "the amount of weeks to look back.")
}


