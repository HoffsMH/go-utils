package main

import (
	"bufio"
	"fmt"
	"git.mhkr.xyz/go-utils"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "prefix",
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

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "prefixes a file with the current iso date, if it does not already have one",
	Long:  "file",

	Aliases: []string{"f"},
	Args:    cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		terms := getTerms(args)
		util.PrefixFiles(terms)
	},
}

var nameCmd = &cobra.Command{
	Use:  "name",
	Long: "outputs a file name to std out with a date prefix if it does not already have one",

	Aliases: []string{"n"},
	Args:    cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		terms := getTerms(args)
		util.PrefixNames(terms)
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(nameCmd)
	rootCmd.AddCommand(fileCmd)
}
