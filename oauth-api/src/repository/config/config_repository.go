package config

import (
	"encoding/json"
	"github.com/Nicrii/Project/oauth-api/src/domain/configuration"
	"os"
)

func NewConfigRepository() ConfigRepository {
	return &configRepository{}
}

type ConfigRepository interface {
	Get() (*configuration.Configuration, error)
}

type configRepository struct {
}

func (r *configRepository) Get() (config *configuration.Configuration, err error) {
	file, err := os.Open("./oauth-api/config.json")
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)

	if err != nil {
		return nil, err
	}
	return config, err
}
