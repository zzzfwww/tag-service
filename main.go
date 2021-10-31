package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	pb "tag-service/proto"
	"tag-service/server"
)

var port string

func init(){
	port = "8004"
}

func main() {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("server.Serve err: %v", err)
	}
}
