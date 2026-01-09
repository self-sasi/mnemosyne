package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

const defaultConfigFilename string = "mnemo-config"

func SetupConfig(configPath string) error {
	if err := loadConfig(configPath); err != nil {
		return err
	}
	return viper.ReadInConfig()
}

func loadConfig(configPath string) error {
	if configPath == "" {
		viper.AddConfigPath(".")
		viper.SetConfigName(defaultConfigFilename)
		return nil
	}

	info, err := os.Stat(configPath)
	if err != nil {
		return fmt.Errorf("config path %q: %w", configPath, err)
	}

	if info.IsDir() {
		viper.AddConfigPath(configPath)
		viper.SetConfigName(defaultConfigFilename)
	} else {
		viper.SetConfigFile(filepath.Clean(configPath))
	}

	return nil
}
