package config

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Load the auth config (jwt) from $HOME/.config/paastech/auth.yaml
func LoadAuthConfig() error {
	// Define config path
	authConfigFile := filepath.Join(os.Getenv("HOME"), ".config", "paastech", "auth.yaml")
	viper.SetConfigFile(authConfigFile)

	// Check if the auth.yaml file exists, if not, create it
	if _, err := os.Stat(authConfigFile); os.IsNotExist(err) {
		err = createAuthConfig(authConfigFile)
		if err != nil {
			return err
		}
	}

	// Read config from file
	err := viper.ReadInConfig()
	if err != nil {
		return errors.New("Failed to read user config file")
	}

	return nil
}

// Set the jwt in config file
func SetJWT(jwt string) {
	// Define config path
	authConfigFile := filepath.Join(os.Getenv("HOME"), ".config", "paastech", "auth.yaml")
	viper.SetConfigFile(authConfigFile)

	// Change jwt value in config file
	viper.Set("jwt", jwt)
	viper.WriteConfig()
}

// Create the auth config file in the home config directory
func createAuthConfig(path string) error {
	// Create the $HOME/.config/paastech directory if it doesn't exist
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return errors.New("Failed to create auth.yaml config file")
	}

	// Write config with no real value
	viper.Set("jwt", "")
	err = viper.WriteConfigAs(path)
	if err != nil {
		return err
	}

	return nil
}
