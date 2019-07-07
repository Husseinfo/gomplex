package main

import (
	"fmt"
	"goahead/service"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net"
)

type conf struct {
	Host string `yaml:"host"`
	Port int64  `yaml:"port"`
}

func (c *conf) getConf() *conf {

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func main() {
	var c conf
	c.getConf()

	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%d",c.Host, c.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	service.RegisterComplexServer(server, &service.ComplexNumberServer{})
	e := server.Serve(lis)
	if e != nil {
		panic(e)
	}
}
