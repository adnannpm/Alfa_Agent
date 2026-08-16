package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	configPath string
)

var rootCmd = &cobra.Command{
	Use:   "alfa_agent",
	Short: "Manage and run the Alfa Agent",
	Long: `Alfa Agent is a command-line application for managing an Alfa Agent node.

Use it to initialize the agent configuration and database, enroll the agent
with a master server, or run the agent service.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "/etc/alfa/", "config file)")
}
