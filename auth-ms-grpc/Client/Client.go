package main

import (
	"auth-ms/auth-ms-grpc/protos/pb"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	client := pb.NewAuthClient(conn)

	req := pb.LoginRequest{
		Email:    "www.anuragshukla@gmail.com",
		Password: "asdf1234",
	}

	res, err := client.Login(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("response: %v", res)
}
