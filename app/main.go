package main

import (
	"fmt"
	"goahead/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	var c Conf
	c.GetConf()

	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%d", c.Host, c.Port))
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
