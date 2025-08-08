package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version",
	Long:  "Displays the application version number, revision, build date, and other information.Including this information when reporting bugs will help us in our development efforts.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("todogo V0.1 -- HEAD")
	},
}
