package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// boiler plate
var rootCmd = &cobra.Command{
	Use:   "cmdex",
	Short: "cmdex is a cli tool designed to explain linux commands",
	Long:  "I can do this whenever",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
