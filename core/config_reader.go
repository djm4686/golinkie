package main

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
//	"fmt"
)

type Config struct{
  RootDir string
	Title string
}

func get_config() Config{
	var config Config
	var configfile = "local_config.toml"
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing", &configfile)
		return config
	}
	_, err = toml.DecodeFile(configfile, &config)
	if err != nil {
		log.Fatal(err)
		return config
	}
	return config
}
