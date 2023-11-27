package main

import (
	"auth-ms/auth-ms-grpc/protos/pb"
	"auth-ms/auth-ms-grpc/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

func main() {
	gs := grpc.NewServer()

	c := &server.Server{}

	pb.RegisterAuthServer(gs, c)

	reflection.Register(gs)

	l, err := net.Listen("tcp", ":8080")

	if err != nil {
		os.Exit(1)
	}

	gs.Serve(l)
}
