package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lt",
	Short: "LeetQ is a CLI tool for fetching random LeetCode problems and test cases.",
	Long:  "LeetQ is a CLI tool that retrieves random LeetCode problems, including descriptions and example test cases, for efficient problem-solving.",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing LeetQ '%s'\n", err)
		os.Exit(1)
	}
}
