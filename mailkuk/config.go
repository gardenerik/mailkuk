package main

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Route struct {
	Mail    string            `yaml:"mail"`
	Url     string            `yaml:"url"`
	Headers map[string]string `yaml:"headers"`
}

type Server struct {
	ListenAddr string `yaml:"listen_addr"`
	Domain     string `yaml:"domain"`
	Debug      bool   `yaml:"debug"`
}

type Config struct {
	Routing []Route `yaml:"routing"`
	Server  Server  `yaml:"server"`
}

func loadConfig() (Config, error) {
	cfgData, err := os.ReadFile("config.yml")
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = yaml.Unmarshal(cfgData, &config)
	return config, err
}
