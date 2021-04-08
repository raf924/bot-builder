package pkg

import (
	"bytes"
	"fmt"
	botbuilder "github.com/raf924/bot-builder"
	"github.com/raf924/bot-builder/internal/pkg/recipe"
	"github.com/spf13/cobra"
	"go/format"
	"gopkg.in/yaml.v2"
	"io"
	"os"
	"text/template"
)

var DependencyTemplate = template.Must(template.ParseFS(botbuilder.Templates, "assets/deps.go.tpl"))

func Build(r recipe.Recipe) ([]byte, error) {
	tpl, err := template.ParseFS(botbuilder.Templates, r.TemplatePattern())
	if err != nil {
		return nil, fmt.Errorf("template error: %s", err.Error())
	}
	buffer := bytes.NewBuffer([]byte{})
	err = BuildDependencies(buffer, r)
	if err != nil {
		return nil, fmt.Errorf("template error: %s", err.Error())
	}
	err = tpl.Execute(buffer, nil)
	if err != nil {
		return nil, fmt.Errorf("template error: %s", err.Error())
	}
	b, err := format.Source(buffer.Bytes())
	if err != nil {
		return nil, fmt.Errorf("formatting error: %s", err.Error())
	}
	return b, nil
}

func BuildDependencies(writer io.Writer, recipe recipe.Recipe) error {
	return DependencyTemplate.ExecuteTemplate(writer, "dependencies", recipe)
}

func MakeBuilderCommand(command string, recipe recipe.Recipe) *cobra.Command {
	cmd := &cobra.Command{
		Use: command,
		RunE: func(cmd *cobra.Command, args []string) error {
			configFile := cmd.Flag("recipe").Value.String()
			output := cmd.Flag("output").Value.String()
			f, err := os.Open(configFile)
			if err != nil {
				return err
			}
			if err := yaml.NewDecoder(f).Decode(recipe); err != nil {
				return err
			}
			source, err := Build(recipe)
			if err != nil {
				return err
			}
			f, err = os.Create(output)
			if err != nil {
				return err
			}
			_, err = f.Write(source)
			if err != nil {
				return err
			}
			_ = f.Close()
			return nil
		},
	}
	return cmd
}
