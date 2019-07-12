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
		log.Printf("config.yaml not found; using default conf #%v ", err)
		return &Conf{
			Host: "localhost",
			Port: 8000,
		}
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Printf("failed to parse config.yaml; using deafult conf: %v", err)
		return &Conf{
			Host: "localhost",
			Port: 8000,
		}
	}

	log.Printf("Loaded config: %v:%d", c.Host, c.Port)

	return c
}
