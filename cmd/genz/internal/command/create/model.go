package create

import (
	"github.com/geekswamp/zen/cmd/genz/internal/template"
	"github.com/spf13/cobra"
)

var modelCmd = &cobra.Command{
	Use:     "model <name> [-d dir]",
	Short:   "Create a new model.",
	Args:    cobra.MinimumNArgs(1),
	Example: "genz create model user",
	RunE:    runModelE,
}

func init() {
	modelCmd.Flags().StringVarP(&dir, "dir", "d", "", "Specify the model directory")
}

func runModelE(_ *cobra.Command, args []string) error {
	m := new(template.Make)
	m.FeatureName = args[0]
	m.FileType = template.Model

	if dir != "" {
		m.FilePath = template.FilePath(dir)
	} else {
		m.FilePath = template.ModelPath
	}

	err := m.Generate()
	if err != nil {
		return err
	}

	return nil
}
