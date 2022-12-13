package main

import (
	"fmt"
	"os"

	util "git.mhkr.xyz/go-utils"
	"github.com/spf13/cobra"
)

var count int
var days int
var weeks int
var months int
var ignore bool

var rootCmd = &cobra.Command{
	Use: "tfilter",
	Long: "outputs a file name to std out with a date prefix if it does not already have one",
	Args:    cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		terms := util.GetTerms(args)

    tr := &util.TimeRange{
      Months: months,
      Weeks: weeks,
      Days: days,
    }

    util.PrintList(util.TFilter(terms, tr, count, ignore))
  },
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
  rootCmd.Flags().IntVarP(&count, "count", "c", 0, "alternatively provide a count of files")
  rootCmd.Flags().IntVarP(&days, "days", "d", 0, "the amount of days to look back.")
  rootCmd.Flags().IntVarP(&weeks, "weeks", "w", 0, "the amount of weeks to look back.")
  rootCmd.Flags().IntVarP(&months, "months", "m", 0, "the amount of months to look back.")
  rootCmd.Flags().BoolVarP(&ignore, "ignore", "i", true, "ignore filenames with errors")
}


