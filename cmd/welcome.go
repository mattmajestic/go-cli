package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var welcomeCmd = &cobra.Command{
    Use:   "welcome",
    Short: "Welcome Majestic Coding!",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("👋 Welcome Majestic Coders! 🎉")
        fmt.Println("Thrilled to have you here. Let's embark on a coding journey together!")
        fmt.Println()
        fmt.Println("📺 Check out the YouTube channel for tutorials and live coding sessions:")
        fmt.Println("   https://youtube.com/@majesticcoding")
        fmt.Println()
        fmt.Println("🎮 Join me on Twitch for live streams and interactive coding:")
        fmt.Println("   https://twitch.tv/majesticcodingtwitch")
        fmt.Println()
        fmt.Println("Next Run learn in the CLI 🚀")
    },
}

func init() {
    rootCmd.AddCommand(welcomeCmd)
}