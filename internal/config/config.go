package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config holds service list from JSON
type Config struct {
	Services []string `json:"services"`
}

// Load reads and parses the config file
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %w", err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("invalid JSON: %w", err)
	}

	return &cfg, nil
}
