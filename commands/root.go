package commands

import (
	"cli/commands/children"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "app",
		Short: "A CLI for interfacing with the ValenciaRateMyProfessor",
		Long: `ValenciaRateMyProfessor is a tool Warren Snipes 
created to find the best professors at Valencia 
College based on Course search in atlas.`,
	}
)

func RootExecute() error {
	return rootCmd.Execute()
}

func RootInit() {
	rootCmd.AddCommand(children.ScrapeCmd, children.FindCmd)
}
