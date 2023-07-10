package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type JWTInfos struct {
	Username       string
	ExpirationTime time.Time
}

// Load the auth config (jwt) from $HOME/.config/paastech/auth.yaml
func LoadAuthConfig() (*viper.Viper, error) {
	// Define config path
	authConfigFile := filepath.Join(os.Getenv("HOME"), ".config", "paastech", "auth.yaml")
	cfg := viper.New()

	// Check if the auth.yaml file exists, if not, create it
	if _, err := os.Stat(authConfigFile); os.IsNotExist(err) {
		err = createAuthConfig(cfg, authConfigFile)
		if err != nil {
			return nil, err
		}
	}

	cfg.SetConfigFile(authConfigFile)
	// Read config from file
	err := cfg.ReadInConfig()
	if err != nil {
		return nil, errors.New("Failed to read user config file")
	}

	return cfg, nil
}

// Extract payload from jwt and return it as UserInfos struct
func ExtractJWTInfos(cfg *viper.Viper) (JWTInfos, error) {
	// Get jwt from config
	tokenString := cfg.GetString("jwt")
	if tokenString == "" {
		return JWTInfos{}, errors.New("Not logged in")
	}

	// Parse payload
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		fmt.Println(tokenString)
		return JWTInfos{}, errors.New("Impossible to parse user jwt")
	}
	claims := token.Claims.(jwt.MapClaims)

	return JWTInfos{
		Username:       claims["username"].(string),
		ExpirationTime: time.Unix(int64(claims["exp"].(float64)), 0),
	}, nil
}

// Create the auth config file in the home config directory
func createAuthConfig(cfg *viper.Viper, path string) error {
	// Create the $HOME/.config/paastech directory if it doesn't exist
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return errors.New("Failed to create auth.yaml config file")
	}

	// Write config with no real value
	cfg.Set("jwt", "")
	cfg.Set("server", "")
	err = cfg.WriteConfigAs(path)
	if err != nil {
		return err
	}

	return nil
}
