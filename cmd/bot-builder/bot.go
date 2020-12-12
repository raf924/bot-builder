package main

import "github.com/raf924/bot-builder/internal/pkg"

type botConfig struct {
	BotRelay   string   `yaml:"botRelay"`
	CmdModules []string `yaml:"cmdModules"`
}

func init() {
	rootCmd.AddCommand(pkg.MakeCommand("bot", "assets/bot.tpl", &botConfig{}))
}
