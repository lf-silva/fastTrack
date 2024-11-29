/*
Copyright © 2024 Luis Silva silva.luisfilipe@hotmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fastTrack",
	Short: "fastrack is a quiz about Fast Track",
	Long: `This cli is integrated with a REST API and to start your quiz, please
type 'cli start' and it will prompt a bunch of questions that you can answer.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
