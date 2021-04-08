package pkg

import (
	"bytes"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/raf924/bot-builder/internal/pkg/recipe"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io"
	"os"
	"strings"
	"sync"
)

func MakeCheckerCommand(command string, config recipe.Recipe) *cobra.Command {
	cmd := &cobra.Command{
		Use: command,
		RunE: func(cmd *cobra.Command, args []string) error {
			configFile := cmd.Flag("recipe").Value.String()
			hashFile := cmd.Flag("hashFile").Value.String()
			f, err := os.Open(configFile)
			if err != nil {
				return err
			}
			if err := yaml.NewDecoder(f).Decode(config); err != nil {
				return err
			}
			//TODO: check dependencies + write last commit hash in timestamp-named file
			hashes, err := CheckDependencies(config)
			if err != nil {
				return err
			}
			hashString := strings.Join(hashes, ";")
			hF, err := os.OpenFile(hashFile, os.O_RDWR|os.O_CREATE, 0755)
			if err != nil {
				return err
			}
			defer hF.Close()
			b, err := io.ReadAll(hF)
			if err != nil {
				return err
			}
			hashBytes := []byte(hashString)
			if bytes.Equal(hashBytes, b) {
				return nil
			}
			_, err = hF.WriteAt(hashBytes, 0)
			if err != nil {
				return err
			}
			return nil
		},
	}
	fS := cmd.Flags()
	fS.StringP("hashFile", "f", "dependencies", "")
	return cmd
}

func CheckDependencies(ctp recipe.Recipe) ([]string, error) {
	deps := ctp.Deps()
	hashes := []plumbing.Hash{}
	var wg sync.WaitGroup
	wg.Add(len(deps))
	auth := &http.BasicAuth{
		Username: "",
		Password: "",
	}
	for _, dep := range deps {
		dep := dep
		go func() {
			defer wg.Done()
			repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
				URL:           dep.Path,
				Auth:          auth,
				ReferenceName: "master",
				SingleBranch:  true,
				NoCheckout:    false,
			})
			if err != nil {
				return
			}
			ref, err := repo.Head()
			if err != nil {
				return
			}
			hashes = append(hashes, ref.Hash())
		}()
	}
	wg.Wait()
	plumbing.HashesSort(hashes)
	hashStrings := []string{}
	for _, hash := range hashes {
		hashStrings = append(hashStrings, hash.String())
	}
	return hashStrings, nil
}
