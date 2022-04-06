package main

import (
	"log"
	"net"

	pb "github.com/VieiraGabrielAlexandre/gRPCWithGoLang/pb"
	services "github.com/VieiraGabrielAlexandre/gRPCWithGoLang/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &services.UserService{})
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}
}
