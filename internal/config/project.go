package config

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type BuildpackConfig struct {
	Path      string   `yaml:"path"`
	Builder   string   `yaml:"builder"`
	Buildpack string   `yaml:"buildpack"`
	Env       []string `yaml:"env"`
}

// Load the project config from ./paastech.yaml
func LoadProjectConfig() (*viper.Viper, error) {
	// Set project config file as config
	projectConfigFile := filepath.Join(".", "paastech.yaml")

	cfg := viper.New()

	// Check if the paastch.yaml file exists
	if _, err := os.Stat(projectConfigFile); os.IsNotExist(err) {
		return nil, errors.New("Project is not initialized")
	}

	cfg.SetConfigFile(projectConfigFile)

	// Read config from file
	err := cfg.ReadInConfig()
	if err != nil {
		return nil, errors.New("Failed to read user config file")
	}

	return cfg, nil
}

func ProjectExists() bool {
	// Set project config file as config
	projectConfigFile := filepath.Join(".", "paastech.yaml")

	// Check if the paastch.yaml file exists
	if _, err := os.Stat(projectConfigFile); os.IsNotExist(err) {
		return false
	}

	return true
}

// Create the project config file in the current directory
func CreateProjectConfig() error {
	// Set project config file as config
	projectConfigFile := filepath.Join(".", "paastech.yaml")

	cfg := viper.New()

	// Create the file in current directory
	_, err := os.Create(filepath.Join(projectConfigFile))
	if err != nil {
		return errors.New("Failed to create paastech.yaml config file")
	}

	cfg.SetConfigFile(projectConfigFile)

	// Write config with no real value
	cfg.Set("project", "")

	// Default buildpacks config
	cfg.Set("buildpacks", BuildpackConfig{
		Path:      ".",
		Builder:   "paketobuildpacks/builder:full",
		Buildpack: "",
		Env:       []string{},
	})
	err = cfg.WriteConfigAs(projectConfigFile)
	if err != nil {
		return err
	}

	return nil
}
