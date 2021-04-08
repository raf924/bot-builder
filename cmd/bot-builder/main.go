package main

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "bot-builder",
}

func main() {
	flags := rootCmd.PersistentFlags()
	flags.StringP("recipe", "r", "recipe.yaml", "-r </path/to/recipe.yaml>")
	flags.StringP("output", "o", "bot-builder-output.go", "-o <path/to/file.go>")
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
