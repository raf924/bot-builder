package recipe

type Recipe interface {
	Deps() []Dependency
	TemplatePattern() string
}
