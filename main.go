package main

import (
	"grpc_mongo_crud/db"
	"grpc_mongo_crud/server.go"
	"log"
	"net"

	pb "grpc_mongo_crud/grpc_mongo_crud/proto"

	"google.golang.org/grpc"
)

func main() {
	err := db.InitMongo()
	if err != nil {
		log.Fatal("MongoConnection failed", err)
	}
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to listen  ", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server.UserServer{})
	log.Println("GRPC server running at port : 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to server ", err)
	}
}
