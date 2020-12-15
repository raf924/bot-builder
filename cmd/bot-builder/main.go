package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "bot-builder",
}

func main() {
	_ = rootCmd.Execute()
}
