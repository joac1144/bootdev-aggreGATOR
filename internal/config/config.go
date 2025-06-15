package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return Config{}, err
	}

	jsonConfigData, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err = json.Unmarshal(jsonConfigData, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func (config *Config) SetUser(user string) error {
	config.CurrentUserName = user

	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	jsonData, err := json.MarshalIndent(config, "", "	")
	if err != nil {
		return err
	}

	if err = os.WriteFile(configPath, jsonData, 0644); err != nil {
		return err
	}

	return nil
}

func getConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return home + "/.gatorconfig.json", nil
}
