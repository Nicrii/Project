package configuration

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Postgres struct {
		Host     string
		Port     string
		User     string
		Password string
		Dbname   string
	}
}

func (config *Configuration) ReadConfiguration() error {
	file, err := os.Open("./users-api/config.json")
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)

	if err != nil {
		return err
	}
	return err
}
