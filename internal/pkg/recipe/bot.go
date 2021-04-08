package recipe

type BotRecipe struct {
	Version    string       `yaml:"version"`
	BotRelay   Dependency   `yaml:"botRelay"`
	CmdModules []Dependency `yaml:"cmdModules"`
}

func (b *BotRecipe) TemplatePattern() string {
	return "assets/bot.go.tpl"
}

func (b *BotRecipe) Deps() []Dependency {
	return append([]Dependency{{yamlDependency{
		Path:    "github.com/raf924/bot",
		Version: b.Version,
	}}, b.BotRelay}, b.CmdModules...)
}
