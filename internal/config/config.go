package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	TWITTER struct {
		TWITTER_BEARER_TOKEN string `yaml:"TWITTER_BEARER_TOKEN"`
	} `yaml:"TWITTER"`
	TARGET []string `yaml:"TARGET"`
}

var config Config

func init() {
	buf, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		log.Fatalln(err)
	}
}

func GetConfig() Config {
	return config
}
