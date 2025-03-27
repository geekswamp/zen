package create

import (
	"github.com/geekswamp/zen/cmd/genz/internal/template"
	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Use:     "service <name>",
	Short:   "Create a new service.",
	Args:    cobra.MinimumNArgs(1),
	Example: "genz create service user",
	RunE:    runServiceE,
}

func init() {
	serviceCmd.Flags().StringVarP(&dir, "dir", "d", "", "Specify the service directory.")
}

func runServiceE(_ *cobra.Command, args []string) error {
	tm.FeatureName = args[0]
	tm.FileType = template.Service
	tm.SuffixFile = template.ServiceSuffix

	if dir != "" {
		tm.FilePath = template.FilePath(dir)
	} else {
		tm.FilePath = template.ServicePath
	}

	if err := tm.Generate(); err != nil {
		return err
	}

	return nil
}
