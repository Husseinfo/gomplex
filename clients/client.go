package main

import (
	"context"
	"fmt"
	"gomplex/service"
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

	var real float32
	var imag float32
	fmt.Print("Real part: ")
	fmt.Scanf("%f", &real)
	fmt.Print("Imaginary part: ")
	fmt.Scanf("%f", &imag)

	number := service.ComplexNumber{
		Real:      real,
		Imaginary: imag,
	}

	tp, err := client.Type(context.Background(), &number)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Type of \n{\n\tReal:\t %.2f \n\tImaginary:\t %.2f\n}\nis %v\n", number.Real, number.Imaginary, tp.Type)
}
