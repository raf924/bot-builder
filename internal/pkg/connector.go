package pkg

import (
	"bytes"
	bot_builder "github.com/raf924/bot-builder"
	"text/template"
)

type ConnectorConfig struct {
	BotRelay        string `yaml:"botRelay"`
	ConnectionRelay string `yaml:"connectionRelay"`
}

func (c *ConnectorConfig) Deps() []string {
	return []string{c.BotRelay, c.ConnectionRelay}
}

func (c *ConnectorConfig) Template() ([]byte, error) {
	tpl, err := template.ParseFS(bot_builder.Templates, "assets/connector.go.tpl")
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer([]byte{})
	err = tpl.Execute(buffer, c)
	return buffer.Bytes(), nil
}
