package sql_test

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Configuration struct {
	CORE_DB_HOST     string `yaml:"CORE_DB_HOST"`
	CORE_DB_PORT     string `yaml:"CORE_DB_PORT"`
	CORE_DB_NAME     string `yaml:"CORE_DB_NAME"`
	CORE_DB_USER     string `yaml:"CORE_DB_USER"`
	CORE_DB_PASSWORD string `yaml:"CORE_DB_PASSWORD"`
	SKU_DB_HOST      string `yaml:"SKU_DB_HOST"`
	SKU_DB_PORT      string `yaml:"SKU_DB_PORT"`
	SKU_DB_NAME      string `yaml:"SKU_DB_NAME"`
	SKU_DB_USER      string `yaml:"SKU_DB_USER"`
	SKU_DB_PASSWORD  string `yaml:"SKU_DB_PASSWORD"`
	SERVICE_API_URL  string `yaml:"SERVICE_API_URL"`
	KHAOS_VERSION    string `yaml:"KHAOS_VERSION"`
	NURSE_VERSION    string `yaml:"NURSE_VERSION"`
	SYS_DOVER        string `yaml:"SYS_DOVER"`
}

var configuration *Configuration

func LoadConfiguration(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}
	var config Configuration
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalln(err)
	}
	configuration = &config
	return err
}

func GetConfiguration() *Configuration {
	return configuration
}
