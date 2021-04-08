package main

import (
	"github.com/raf924/bot-builder/internal/pkg"
	"github.com/raf924/bot-builder/internal/pkg/recipe"
)

func init() {
	rootCmd.AddCommand(pkg.MakeBuilderCommand("connector", &recipe.ConnectorRecipe{}))
}
