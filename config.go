package main

import (
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

type DirConfig struct {
	Path    string `yaml:"path"`
	LogPath string `yaml:"logPath"`
}

type Config struct {
	Magick string      `yaml:"magick"`
	Dirs   []DirConfig `yaml:"dirs"`
}

var config *Config = &Config{}

func parseConfig() {
	cfgPath := path.Join(wd, "date-mark.yml")
	cfgFile, err := os.Open(cfgPath)
	if err != nil {
		log.Panic("Failed to open config file,", err)
	}
	d := yaml.NewDecoder(cfgFile)
	err = d.Decode(config)
	if err != nil {
		log.Panic("Failed to parse config file at", cfgPath, ",", err)
	}
}
