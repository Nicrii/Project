package configuration

import (
	"encoding/json"
	"os"
)

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
