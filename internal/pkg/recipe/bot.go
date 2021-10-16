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
	botDependency := Module("github.com/raf924/bot", b.Version)
	return append([]Dependency{botDependency, b.BotRelay}, b.CmdModules...)
}
