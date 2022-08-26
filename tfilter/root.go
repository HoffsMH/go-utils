package main

import (
	"bufio"
	"fmt"
	"git.mhkr.xyz/go-utils"
	"github.com/spf13/cobra"
	"os"
)

var days int
var weeks int
var months int
var ignore bool

var rootCmd = &cobra.Command{
	Use: "tfilter",
	Long: "outputs a file name to std out with a date prefix if it does not already have one",
	Args:    cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		terms := getTerms(args)
    tr := &util.TimeRange{
      Months: months,
      Weeks: weeks,
      Days: days,
    }

    printlist(util.TFilter(terms, tr, ignore))
  },
}

func getTerms(args []string) []string {
	terms := []string{}

	if len(args) > 0 {
		terms = args
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			terms = append(terms, scanner.Text())
		}
	}

	return terms
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printlist(list []string) {
  for _, str := range list {
    fmt.Println(str)
  }
}

func init() {
  rootCmd.Flags().IntVarP(&days, "days", "d", 2, "the amount of days to look back.")
  rootCmd.Flags().IntVarP(&weeks, "weeks", "w", 0, "the amount of weeks to look back.")
  rootCmd.Flags().IntVarP(&months, "months", "m", 0, "the amount of months to look back.")
  rootCmd.Flags().BoolVarP(&ignore, "ignore", "i", true, "ignore filenames with errors")
}


