package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	util "git.mhkr.xyz/go-utils"
	"github.com/spf13/cobra"
)

var lessThan int
var greaterThan int
var kb int64 = 1000
var mb int64 = 1000 * kb
var gb int64 = 1000 * mb

var rootCmd = &cobra.Command{
	Use: "sfilter",
	Long: "outputs a file name to std out with a date prefix if it does not already have one",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		terms := util.GetTerms(args)


    util.PrintList(util.SFilter(convert(terms[0], greaterThan, lessThan)))
  },
}

func convert(sizeStr string) int64 {
  var num int64
  if (strings.HasSuffix(sizeStr, "G") ) {
    num, _ = strconv.ParseInt(sizeStr[:len(sizeStr) - 1], 10, 64)

    return num * gb
  }

  if (strings.HasSuffix(sizeStr, "M") ) {
    num, _ = strconv.ParseInt(sizeStr[:len(sizeStr) - 1], 10, 64)

    return num * mb
  }

  if (strings.HasSuffix(sizeStr, "K") ) {
    num, _ = strconv.ParseInt(sizeStr[:len(sizeStr) - 1], 10, 64)

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
  rootCmd.Flags().IntVarP(&lessThan, "less-than", "lt", 2, "the amount of days to look back.")
  rootCmd.Flags().IntVarP(&greaterThan, "greater-than", "gt", 0, "the amount of weeks to look back.")
}


