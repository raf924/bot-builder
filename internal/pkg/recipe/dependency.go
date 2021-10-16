package recipe

import (
	"gopkg.in/yaml.v2"
	"strings"
)

var _ yaml.Unmarshaler = (*Dependency)(nil)

type Dependency struct {
	yamlDependency
}

type yamlDependency struct {
	Path    string `yaml:"path"`
	Version string `yaml:"version"`
}

func (d *Dependency) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var dependency yamlDependency
	err := unmarshal(&dependency)
	if err != nil {
		var path string
		err = unmarshal(&path)
		if err != nil {
			return err
		}
		parts := strings.Split(path, "@")
		if len(parts) == 1 {
			parts = append(parts, "latest")
		}
		dependency = yamlDependency{
			Path:    parts[0],
			Version: parts[1],
		}
	}
	d.yamlDependency = dependency
	return nil
}

func Module(path string, version string) Dependency {
	if version == "" {
		version = "latest"
	}
	return Dependency{yamlDependency{
		Path:    path,
		Version: version,
	}}
}
