package cli

import (
	"fmt"

	"github.com/self-sasi/mnemosyne/internal/config"
	"github.com/spf13/cobra"
)

var restoreConfigPath string

var restoreCommand = &cobra.Command{
	Use:     "restore",
	Aliases: []string{"r"},
	Short:   "restores the database from a backup",
	Long: `The restore command restores the database from an existing backup copy. 
	
	It uses the configuration file specified by the --configPath flag to determine database connection details and backup settings. The config file can be of json or yaml types. If no path is provided, the command defaults to the current directory.
	
	Use this command with caution, as the existing database state will be overridden with the contents of the backup file.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return config.SetupConfig(restoreConfigPath)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(`configPath passed to restore command: %v\n`, restoreConfigPath)
		fmt.Println(config.GetConfig())
	},
}

func init() {
	rootCommand.AddCommand(restoreCommand)
	restoreCommand.Flags().StringVarP(&restoreConfigPath, "configPath", "p", ".", "Path to a configuration file (YAML or JSON), or a directory containing a file named mnemo-config.")
}
