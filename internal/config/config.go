package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = "/.gatorconfig.json"

func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	fullPath := configFilePath + configFileName

	config, err := os.ReadFile(fullPath)
	if err != nil {
		return Config{}, err
	}

	var cfg Config

	err = json.Unmarshal(config, &cfg)
	if err != nil {
		fmt.Println(err)
		return Config{}, err
	}

	return cfg, nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	// fullPath := homeDir + "/.config/gator/"
	return homeDir, nil
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	return write(*c)
}

func write(cfg Config) error {
	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	fullPath := configPath + configFileName

	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(fullPath, data, 0666)
	if err != nil {
		return err
	}

	return nil
}
