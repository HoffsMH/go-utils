package main

import (
	"fmt"
	"os"
	"time"

	util "git.mhkr.xyz/go-utils"
	"github.com/jmhodges/clock"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "prefix",
}

var nameCmd = &cobra.Command{
	Use:  "name",
	Long: "outputs a file name to std out with a date prefix if it does not already have one",

	Aliases: []string{"n"},
	Args:    cobra.MinimumNArgs(0),
}

var ISOCmd = &cobra.Command{
	Use:  "iso",
	Long: "outputs a file name to std out with a date prefix if it does not already have one",

	Aliases: []string{"n"},
	Args:    cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		terms := util.GetTerms(args)
    prefixer := &util.Prefixer{
      Clock: clock.New(),
      Format: time.RFC3339,
    }
		util.PrintList(prefixer.Names(terms))
	},
}

var DateCmd = &cobra.Command{
	Use:  "date",
	Long: "outputs a file name to std out with a date prefix if it does not already have one",

	Aliases: []string{"n"},
	Args:    cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		terms := util.GetTerms(args)
    prefixer := &util.Prefixer{
      Clock: clock.New(),
      Format: "2006-01-02",
    }
		util.PrintList(prefixer.Names(terms))
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
  nameCmd.AddCommand(ISOCmd)
  nameCmd.AddCommand(DateCmd)
	rootCmd.AddCommand(nameCmd)
}
