package alenar

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
)

type globalConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DataBase string `yaml:"data_base"`
	UserName string `yaml:"user_name"`
	Password string `yaml:"password"`
}

func getConfig() globalConfig {
	filename := "configs/config.yaml"
	config := globalConfig{
		Host:     "localhost",
		Port:     3306,
		DataBase: "db",
		UserName: "root",
		Password: "password",
	}

	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		saveConfig(filename, config)
		return config
	}

	return readConfig(filename)
}

func readConfig(filename string) globalConfig {
	byteConfig, err := os.ReadFile(filename)
	checkError(err)

	config := globalConfig{}
	err = yaml.Unmarshal(byteConfig, &config)
	checkError(err)

	return config
}

func saveConfig(filename string, config globalConfig) {
	marshal, err := yaml.Marshal(config)
	checkError(err)

	err = os.WriteFile(filename, marshal, 0644)
	checkError(err)
}
