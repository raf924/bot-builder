package recipe

import (
	"strings"
)

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
	return Dependency{yamlDependency{
		Path:    path,
		Version: version,
	}}
}
