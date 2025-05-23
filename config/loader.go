package config

import (
	"encoding/json"
	"os"
)

type ClusterConfig struct {
	Name     string `json:"name"`
	API      string `json:"api"`
	AuthType string `json:"authType"`
	Token    string `json:"token"`
}

type Config struct {
	Clusters []ClusterConfig `json:"clusters"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
