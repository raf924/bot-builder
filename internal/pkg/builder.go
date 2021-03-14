package pkg

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go/format"
	"gopkg.in/yaml.v2"
	"os"
)

type ConfigTemplate interface {
	Template() ([]byte, error)
}

func Build(config ConfigTemplate) ([]byte, error) {
	b, err := config.Template()
	if err != nil {
		return nil, err
	}
	source, err := format.Source(b)
	if err != nil {
		return nil, err
	}
	return source, nil
}

func MakeBuilderCommand(command string, config ConfigTemplate) *cobra.Command {
	cmd := &cobra.Command{
		Use: command,
		RunE: func(cmd *cobra.Command, args []string) error {
			_ = viper.BindPFlags(cmd.Flags())
			configFile := viper.GetString("config")
			output := viper.GetString("output")
			f, err := os.Open(configFile)
			if err != nil {
				return err
			}
			if err := yaml.NewDecoder(f).Decode(config); err != nil {
				return err
			}
			source, err := Build(config)
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
	fS := cmd.Flags()
	fS.StringP("config", "c", "config.yaml", "")
	fS.StringP("output", "o", "dist/connector.go", "")
	return cmd
}
