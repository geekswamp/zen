package create

import (
	"github.com/geekswamp/zen/cmd/genz/internal/template"
	"github.com/spf13/cobra"
)

var repoCmd = &cobra.Command{
	Use:     "repo <name> [-d dir]",
	Short:   "Create a new repository",
	Args:    cobra.MinimumNArgs(1),
	Example: "genz create repo user",
	RunE:    runRepoE,
}

func init() {
	repoCmd.Flags().StringVarP(&dir, "dir", "d", "", "Specify the repository directory")
}

func runRepoE(_ *cobra.Command, args []string) error {
	m := new(template.Make)
	m.FeatureName = args[0]
	m.FileType = template.Repository
	m.SuffixFile = template.RepoSuffix

	if dir != "" {
		m.FilePath = template.FilePath(dir)
	} else {
		m.FilePath = template.RepositoryPath
	}

	err := m.Generate()
	if err != nil {
		return err
	}

	return nil
}
