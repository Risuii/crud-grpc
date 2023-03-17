package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "crud-grpc/proto"
	"crud-grpc/server/handler"
	"crud-grpc/server/usecase"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to liste: %v", err)
	}

	userUseCase := usecase.NewUserClass()
	userHandler := handler.NewUserServer(*userUseCase)

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, userHandler)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("error when server grpc", err.Error())
	}

	fmt.Println("server live !")
}
