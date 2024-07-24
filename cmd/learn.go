package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var educateCmd = &cobra.Command{
    Use:   "learn",
    Short: "Learn about building CLI applications with Cobra",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("📚 Educating on Cobra CLI 📚")
        fmt.Println()
        fmt.Println("Cobra is a library for creating powerful modern CLI applications.")
        fmt.Println("It provides a simple interface to define commands, flags, and subcommands.")
        fmt.Println()
        fmt.Println("Key Concepts:")
        fmt.Println("1. Commands: The basic building blocks of a Cobra application.")
        fmt.Println("2. Flags: Options that modify the behavior of commands.")
        fmt.Println("3. Subcommands: Commands within commands, allowing for complex CLI structures.")
        fmt.Println()
        fmt.Println("Example Structure:")
        fmt.Println("  rootCmd")
        fmt.Println("  ├── cmd1")
        fmt.Println("  └── cmd2")
        fmt.Println("      └── subcmd1")
        fmt.Println()
        fmt.Println("Basic Usage:")
        fmt.Println("1. Define a command using `&cobra.Command{}`.")
        fmt.Println("2. Implement the `Run` function to specify what the command does.")
        fmt.Println("3. Add the command to the root command using `rootCmd.AddCommand()`.")
        fmt.Println()
        fmt.Println("For more details, visit the Cobra documentation: https://github.com/spf13/cobra")
        fmt.Println()
        fmt.Println("Happy Coding with Cobra! 🚀")
    },
}

func init() {
    rootCmd.AddCommand(educateCmd)
}