package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	util "git.mhkr.xyz/go-utils"
	"github.com/spf13/cobra"
)

var chars int

var rootCmd = &cobra.Command{
	Use: "wrap",
	Long: "outputs a file name to std out with a date prefix if it does not already have one",
	Args:    cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
    var num int
    if len(args) > 0 {
      num, _ = strconv.Atoi(args[0])
    } else {
      num = 80
    }

    inputString := util.StdinAll()
    if inputString != "" {
       fmt.Print(util.WrapString(inputString, uint(num)))
    } else {
      cmd := exec.Command("xclip", "-selection", "clipboard", "-o")
      clipContents, _ := cmd.Output()

     fmt.Print(util.WrapString(string(clipContents), uint(num)))
    }
  },
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

