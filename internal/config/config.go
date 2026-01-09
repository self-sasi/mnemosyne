package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	Type     string `mapstructure:"type" yaml:"type" json:"type"`
	Host     string `mapstructure:"host" yaml:"host" json:"host"`
	Port     int    `mapstructure:"port" yaml:"port" json:"port"`
	DbName   string `mapstructure:"dbname" yaml:"dbname" json:"dbname"`
	UserName string `mapstructure:"username" yaml:"username" json:"username"`
	Password string `mapstructure:"password" yaml:"password" json:"password"`
}

func (config Config) String() string {
	return fmt.Sprintf(
		"type=%s host=%s port=%d dbname=%s user=%s",
		config.Type,
		config.Host,
		config.Port,
		config.DbName,
		config.UserName,
	)
}

const defaultConfigFilename string = "mnemo-config"

var config Config

func SetupConfig(configPath string) error {
	if err := loadConfig(configPath); err != nil {
		return err
	}
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return viper.Unmarshal(&config)
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

func GetConfig() Config {
	return config
}
