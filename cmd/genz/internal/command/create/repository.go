package create

import (
	"github.com/geekswamp/zen/cmd/genz/internal/template"
	"github.com/spf13/cobra"
)

var repoCmd = &cobra.Command{
	Use:     "repo <name>",
	Short:   "Create a new repository.",
	Args:    cobra.MinimumNArgs(1),
	Example: "genz create repo user",
	RunE:    runRepoE,
}

func init() {
	repoCmd.Flags().StringVarP(&dir, "dir", "d", "", "Specify the repository directory.")
}

func runRepoE(_ *cobra.Command, args []string) error {
	tm.FeatureName = args[0]
	tm.FileType = template.Repository
	tm.SuffixFile = template.RepoSuffix

	if dir != "" {
		tm.FilePath = template.FilePath(dir)
	} else {
		tm.FilePath = template.RepositoryPath
	}

	if err := tm.Generate(); err != nil {
		return err
	}

	return nil
}
