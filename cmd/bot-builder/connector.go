package main

import (
	"github.com/raf924/bot-builder/v2/internal/pkg"
	"github.com/raf924/bot-builder/v2/internal/pkg/recipe"
)

func init() {
	rootCmd.AddCommand(pkg.MakeBuilderCommand("connector", &recipe.ConnectorRecipe{}))
}
