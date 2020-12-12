package main

import (
	"github.com/raf924/bot-builder/internal/pkg"
)

type connectorConfig struct {
	BotRelay        string `yaml:"botRelay"`
	ConnectionRelay string `yaml:"connectionRelay"`
}

func init() {
	rootCmd.AddCommand(pkg.MakeCommand("connector", "assets/connector.tpl", &connectorConfig{}))
}
