package config

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type configType struct {
	LogFilePath    string `yaml:"LogFilePath" default:"./"`
	LogFileNameReg string `yaml:"LogFileNameReg"  default:"dtm.log"`
}

// Config 配置
var Config = configType{}

// MustLoadConfig load config from env and file
func MustLoadConfig(confFile string) {
	log.Println(confFile)

	if confFile != "" {
		cont, _ := ioutil.ReadFile(confFile)
		log.Println(cont)
		_ = yaml.UnmarshalStrict(cont, &Config)
	}
	scont, _ := json.MarshalIndent(&Config, "", "  ")
	log.Printf("config file: %s loaded config is: \n%s\n", confFile, scont)
}
