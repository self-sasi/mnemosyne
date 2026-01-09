package cli

import (
	"fmt"

	"github.com/self-sasi/mnemosyne/internal/config"
	"github.com/spf13/cobra"
)

var backupConfigPath string

var backupCommand = &cobra.Command{
	Use:     "backup",
	Aliases: []string{"b"},
	Short:   "backs up the database",
	Long: `The backup command creates a backup copy of the database. 
	
	It uses the configuration file specified by the --configPath flag to determine database connection details and backup settings. The config file can be of json or yaml types. If no path is provided, the command defaults to the current directory.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return config.SetupConfig(backupConfigPath)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(`configPath passed to backup command: %v\n`, backupConfigPath)
		fmt.Println(config.GetConfig())
	},
}

func init() {
	rootCommand.AddCommand(backupCommand)
	backupCommand.Flags().StringVarP(&backupConfigPath, "configPath", "p", ".", "Path to a configuration file (YAML or JSON), or a directory containing a file named mnemo-config.")
}
