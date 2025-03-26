package create

import (
	"github.com/geekswamp/zen/cmd/genz/internal/template"
	"github.com/spf13/cobra"
)

var (
	dir string
	tm  = new(template.Make)
)

var CreateCmd = &cobra.Command{
	Use:       "create [type]",
	Short:     "Create a new handler, repository, route, service or model.",
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{"handler", "repo", "route", "model", "service"},
}

func init() {
	CreateCmd.AddCommand(modelCmd, repoCmd)
}
