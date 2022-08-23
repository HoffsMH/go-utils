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
	Args: cobra.MinimumNArgs(0),
  Run: func(cmd *cobra.Command, args []string) {
    // filename
    // dir
		fmt.Println(util.Hsplit(stdin(), args[0]))
  },
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func stdin() []string {
  var lines []string
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines
}
