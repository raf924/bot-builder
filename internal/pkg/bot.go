package pkg

import (
	"bytes"
	bot_builder "github.com/raf924/bot-builder"
	"text/template"
)

type BotConfig struct {
	BotRelay   string   `yaml:"botRelay"`
	CmdModules []string `yaml:"cmdModules"`
}

func (b *BotConfig) Template() ([]byte, error) {
	tpl, err := template.ParseFS(bot_builder.Templates, "assets/bot.go.tpl")
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer([]byte{})
	err = tpl.Execute(buffer, b)
	return buffer.Bytes(), nil
}

func (b *BotConfig) Deps() []string {
	return append([]string{b.BotRelay}, b.CmdModules...)
}
