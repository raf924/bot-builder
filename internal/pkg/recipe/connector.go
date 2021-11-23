package recipe

type ConnectorRecipe struct {
	Version         string       `yaml:"version"`
	BotRelay        Dependency   `yaml:"botRelay"`
	ConnectionRelay Dependency   `yaml:"connectionRelay"`
	CmdModules      []Dependency `yaml:"cmdModules"`
}

func (c *ConnectorRecipe) TemplatePattern() string {
	return "assets/connector.go.tpl"
}

func (c *ConnectorRecipe) Deps() []Dependency {
	return append([]Dependency{Module("github.com/raf924/bot/v2", c.Version), c.BotRelay, c.ConnectionRelay}, c.CmdModules...)
}
