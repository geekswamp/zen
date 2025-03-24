package main

import (
	"os"

	"github.com/geekswamp/genz/internal/command/create"
	"github.com/spf13/cobra"
)

var mainCmd = &cobra.Command{
	Use:     "genz",
	Short:   "GenZ is a command-line tool for creating and managing projects.",
	Long:    "GenZ is a command-line tool that automates the creation of repositories, services, and models, streamlining project development with minimal setup.",
	Version: "0.0.1",
}

func init() {
	mainCmd.AddCommand(create.CreateCmd)
}

func main() {
	if err := mainCmd.Execute(); err != nil {
		err := mainCmd.Help()
		if err != nil {
			panic(err)
		}
		os.Exit(1)
	}
}
