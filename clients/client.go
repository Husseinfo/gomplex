package main

import (
	"context"
	"fmt"
	"goahead/service"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := service.NewComplexClient(conn)


	number := service.ComplexNumber{
		Real: -1,
		Imaginary:0,
	}

	tp, err := client.Type(context.Background(), &number)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Print(tp)
}
