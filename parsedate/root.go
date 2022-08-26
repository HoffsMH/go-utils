package main

import (
	"fmt"
	"git.mhkr.xyz/go-utils"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "get_date_portion",
	Args: cobra.MinimumNArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println(util.ParseDateISO(args[0]))
  },
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

