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

	// Check if the auth.yaml file exists, if not, create it
	if _, err := os.Stat(projectConfigFile); os.IsNotExist(err) {
		err = createProjectConfig(cfg, projectConfigFile)
		if err != nil {
			return nil, err
		}
	}

	cfg.SetConfigFile(projectConfigFile)

	// Read config from file
	err := cfg.ReadInConfig()
	if err != nil {
		return nil, errors.New("Failed to read user config file")
	}

	return cfg, nil
}

// Create the project config file in the current directory
func createProjectConfig(cfg *viper.Viper, path string) error {
	// Create the file in current directory
	_, err := os.Create(filepath.Join(path))
	if err != nil {
		return errors.New("Failed to create paastech.yaml config file")
	}

	// Write config with no real value
	cfg.Set("project", "")

	// Default buildpacks config
	cfg.Set("buildpacks", BuildpackConfig{
		Path:      ".",
		Builder:   "heroku/builder:22",
		Buildpack: "",
		Env:       []string{},
	})
	err = cfg.WriteConfigAs(path)
	if err != nil {
		return err
	}

	return nil
}
