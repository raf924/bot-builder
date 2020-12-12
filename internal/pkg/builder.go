package pkg

import (
	"bytes"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go/format"
	"gopkg.in/yaml.v2"
	"os"
	"text/template"
)

func MakeCommand(command, templateFile string, config interface{}) *cobra.Command {
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
			tpl, err := template.ParseFiles(templateFile)
			if err != nil {
				return err
			}
			_ = f.Close()
			b := bytes.NewBuffer([]byte{})
			f, err = os.Create(output)
			if err != nil {
				return err
			}
			err = tpl.Execute(b, config)
			if err != nil {
				return err
			}
			btes, err := format.Source(b.Bytes())
			if err != nil {
				return err
			}
			_, err = f.Write(btes)
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
