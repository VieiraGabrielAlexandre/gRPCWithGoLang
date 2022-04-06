package main

import (
	"context"
	"fmt"
	"log"

	"github.com/VieiraGabrielAlexandre/gRPCWithGoLang/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatal("Could not connect to server: ", err)
	}

	defer connection.Close()

	cliente := pb.NewUserServiceClient(connection)
	AddUser(cliente)

}

func AddUser(cliente pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Gabriel",
		Email: "gabriel@teste.com",
	}

	res, err := cliente.AddUser(context.Background(), req)

	if err != nil {
		log.Fatal("Could not make gRPC server: ", err)
	}

	fmt.Println(res)
}
