package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Majestic Coding CLI",
	Short: "A useful CLI for Devs of the Majestic Coding Channel",
	Long:  `This CLI was made with Go.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Add any global flags or configuration settings here
}
