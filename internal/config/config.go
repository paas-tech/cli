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

// Set and return auth config as config file
func authConfigFile() string {
	// Define config path
	authConfigFile := filepath.Join(os.Getenv("HOME"), ".config", "paastech", "auth.yaml")
	viper.SetConfigFile(authConfigFile)
	viper.ReadInConfig()
	return authConfigFile
}

// Load the auth config (jwt) from $HOME/.config/paastech/auth.yaml
func LoadAuthConfig() error {
	// Set auth config file as config
	authConfigFile := authConfigFile()

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
	// Set auth config file as config
	authConfigFile()

	// Change jwt value in config file
	viper.Set("jwt", jwt)
	viper.WriteConfig()
}

// Extract payload from jwt and return it as UserInfos struct
func ExtractJWTInfos() (JWTInfos, error) {
	// Set auth config file as config
	authConfigFile()

	// Get jwt from config
	tokenString := viper.GetString("jwt")
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
