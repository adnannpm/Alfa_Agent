package helpers

import (
	"alfa_agent/internal/constants"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Auth struct{
		Token 	string	`yaml:token`
	} `yaml:auth`

	HostIp	string `yaml:host_ip`
}

func InsertToken(token, host string) error {
	path := filepath.Join(constants.CONFIG_PATH, "/config.yml")
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}

	config.Auth.Token = token
	config.HostIp = host
	
	newData, err := yaml.Marshal(&config)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, newData, 0644)
	if err != nil {
		return err
	}

	return nil
}