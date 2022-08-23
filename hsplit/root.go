package main

import (
	"fmt"
	"git.mhkr.xyz/go-utils"
	"github.com/spf13/cobra"
	"os"
	"bufio"
)

var rootCmd = &cobra.Command{
	Use: "hsplit",
	Args: cobra.MinimumNArgs(2),
  Run: func(cmd *cobra.Command, args []string) {
		terms := util.GetTerms(args)
    // filename
    // dir
		fmt.Println(util.Hsplit(terms[0], args[1]))
  },
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Stdin() []string {
  var lines []string
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines
}
