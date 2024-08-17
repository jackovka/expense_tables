package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// GetConfig инициализирует и заполняет структуру конфигурационного файла
func GetConfig() (*Config, error) {
	var cfg Config

	congifFile, err := os.Open("./config.json")
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	r, err := io.ReadAll(congifFile)
	if err != nil {
		return nil, fmt.Errorf("cannot read config file with error: %v", err)
	}

	err = json.Unmarshal(r, &cfg)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal config file with error: %v", err)
	}

	return &cfg, nil
}
