package main

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "dep-checker",
}

func main() {
	_ = rootCmd.Execute()
}
