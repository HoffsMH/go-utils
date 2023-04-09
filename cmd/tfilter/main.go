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

    util.PrintList(util.TFilter(terms, tr, ignore, reject))
  },
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}


func init() {
  rootCmd.Flags().IntVarP(&days, "days", "d", 2, "the amount of days to look back.")
  rootCmd.Flags().IntVarP(&weeks, "weeks", "w", 0, "the amount of weeks to look back.")
  rootCmd.Flags().IntVarP(&months, "months", "m", 0, "the amount of months to look back.")
  rootCmd.Flags().BoolVarP(&ignore, "ignore", "i", true, "Ignore filenames with errors")
  rootCmd.Flags().BoolVarP(&reject, "reject", "r", false, "alternatively reject any files that meet the criteria")
}


