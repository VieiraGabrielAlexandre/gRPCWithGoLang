package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	"github.com/VieiraGabrielAlexandre/gRPCWithGoLang/pb"
	"github.com/VieiraGabrielAlexandre/gRPCWithGoLang/services"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}