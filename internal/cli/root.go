package cli

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:     "mnemo",
	Aliases: []string{"mnemosyne"},
	Short:   "Backup Utility tool for your Database",
	Long: `mnemosyne is a database backup utility tool that can autonomously safeguard your database contents. 
	
	The name 'Mnemosyne' originates from Greek Mythology, where Mnemosyne is the Titaness of Memory who personifies the power of rememberance, much like what this CLI tool aims to do.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCommand.Execute()
	if err != nil {
		os.Exit(1)
	}
}
