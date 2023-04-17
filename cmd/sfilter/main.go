package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	util "git.mhkr.xyz/go-utils"
	"github.com/spf13/cobra"
)

var lessThan string
var greaterThan string
var kb int = 1000
var mb int = 1000 * kb
var gb int = 1000 * mb

var rootCmd = &cobra.Command{
	Use: "sfilter",
	Long: "outputs a file name to std out with a date prefix if it does not already have one",
	Args:    cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
    util.PrintList(util.SFilter(util.Stdin(), convert(greaterThan), convert(lessThan)))
  },
}

func convert(sizeStr string) int {
  var num int = -1
  if sizeStr == "" {
    return num
  }

  num, err := strconv.Atoi(sizeStr[:len(sizeStr) - 1])
  if err != nil {
    log.Fatalf("cannot conver %q to num", sizeStr)
  }

  if (strings.HasSuffix(sizeStr, "G") ) {
    return num * gb
  }

  if (strings.HasSuffix(sizeStr, "M") ) {
    return num * mb
  }

  if (strings.HasSuffix(sizeStr, "K") ) {
    return num * kb
  }
  return num
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}


func init() {
  rootCmd.Flags().StringVarP(&lessThan, "less-than", "l", "", "the amount of days to look back.")
  rootCmd.Flags().StringVarP(&greaterThan, "greater-than", "g", "500M", "the amount of weeks to look back.")
}


