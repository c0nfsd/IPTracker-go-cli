package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "IPTracker-go-cli",
		Short: "IP tracker",
		Long:  `IP tracker`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
