package main

import (
	"bufio"
	"fmt"
	"os"

	util "git.mhkr.xyz/go-utils"
	"github.com/spf13/cobra"
)

var days int
var weeks int
var months int
var dir string

var rootCmd = &cobra.Command{
	Use: "<search terms separated by new lines> | bulksearch <directories>",
	Run: func(cmd *cobra.Command, args []string) {
		scanner := bufio.NewScanner(os.Stdin)

		terms := []string{}
		for scanner.Scan() {
			terms = append(terms, scanner.Text())
		}
		results := util.Search(terms, args)
		for _,v := range results {
			fmt.Println(v.Term, ",", v.FileCount)
		}
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

