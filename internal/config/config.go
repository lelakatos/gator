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

const configFileName = ".gatorconfig.json"

func Read() Config {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		fmt.Println(err)
		return Config{}
	}

	fullPath := configFilePath + configFileName

	config, err := os.ReadFile(fullPath)
	if err != nil {
		fmt.Println(err)
		return Config{}
	}

	var cfg Config

	err = json.Unmarshal(config, &cfg)
	if err != nil {
		fmt.Println(err)
		return Config{}
	}

	return cfg
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullPath := homeDir + "/.config/gator/"
	return fullPath, nil
}
