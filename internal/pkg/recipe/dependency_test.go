package recipe

import (
	"gopkg.in/yaml.v2"
	"testing"
)

func TestDependency_UnmarshalYAML(t *testing.T) {
	type args struct {
		yaml string
	}
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		resultWanted Dependency
	}{
		{
			name: "Unmarshal object",
			args: args{
				yaml: `d:
  path: "github.com/golang/time"
  version: "v1.0.0"
`,
			},
			wantErr: false,
			resultWanted: Dependency{yamlDependency{
				Path:    "github.com/golang/time",
				Version: "v1.0.0",
			}},
		},
		{
			name: "Unmarshal string without version",
			args: args{
				yaml: `d: github.com/golang/time`,
			},
			wantErr: false,
			resultWanted: Dependency{yamlDependency{
				Path:    "github.com/golang/time",
				Version: "latest",
			}},
		},
		{
			name: "Unmarshal string with version",
			args: args{
				yaml: `d: github.com/golang/time@v1.0.0`,
			},
			wantErr: false,
			resultWanted: Dependency{yamlDependency{
				Path:    "github.com/golang/time",
				Version: "v1.0.0",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var d struct {
				D Dependency `yaml:"d"`
			}
			if err := yaml.Unmarshal([]byte(tt.args.yaml), &d); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalYAML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if d.D != tt.resultWanted {
				t.Errorf("UnmarshalYAML() result = %v, resultWanted %v", d.D, tt.resultWanted)
			}
		})
	}
}
