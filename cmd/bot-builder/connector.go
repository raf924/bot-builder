package main

import (
	"github.com/raf924/bot-builder/internal/pkg"
)

func init() {
	rootCmd.AddCommand(pkg.MakeBuilderCommand("connector", &pkg.ConnectorConfig{}))
}
