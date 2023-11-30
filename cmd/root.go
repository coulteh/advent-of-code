package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var fPartTwo bool

func init() {
	rootCmd.PersistentFlags().BoolVar(&fPartTwo, "part2", false, "Run part two of the challenge")
}

var rootCmd = &cobra.Command{
	Use:   "advent",
	Short: "Advent of Code",
	Long:  `A tool to generate my advent of code entries.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
