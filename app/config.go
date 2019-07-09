package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Conf struct {
	Host string `yaml:"host"`
	Port int64  `yaml:"port"`
}

func (c *Conf) GetConf() *Conf {

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	log.Printf("Loaded config: %v:%d", c.Host, c.Port)

	return c
}
