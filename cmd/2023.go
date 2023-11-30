package cmd

import (
	twenty23 "advent-of-code/2023"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

func init() {
	rootCmd.AddCommand(_2023Cmd)
}

var _2023Cmd = &cobra.Command{
	Use:   "2023",
	Short: "Run 2023 challenges",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		nonFlags := cmd.Flags().Args()
		day, err := strconv.Atoi(nonFlags[0])
		if err != nil {
			log.Fatal(err)
		}
		err = twenty23.DaySelector(day, fPartTwo)
		if err != nil {
			log.Fatal(err)
		}
	},
}
