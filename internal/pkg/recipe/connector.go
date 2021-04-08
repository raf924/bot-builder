package recipe

type ConnectorRecipe struct {
	Version         string     `yaml:"version"`
	BotRelay        Dependency `yaml:"botRelay"`
	ConnectionRelay Dependency `yaml:"connectionRelay"`
}

func (c *ConnectorRecipe) TemplatePattern() string {
	return "assets/connector.go.tpl"
}

func (c *ConnectorRecipe) Deps() []Dependency {
	return []Dependency{{yamlDependency{
		Path:    "github.com/raf924/bot",
		Version: c.Version,
	}}, c.BotRelay, c.ConnectionRelay}
}
