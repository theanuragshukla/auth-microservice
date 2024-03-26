package main

import (
	"auth-ms/auth-ms-grpc/protos/pb"
	"auth-ms/auth-ms-grpc/server"
	"auth-ms/data"
	"auth-ms/handlers"
	"auth-ms/utils"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

func main() {
	gs := grpc.NewServer()
	l := utils.CreateLogger()
	_ = l.Sync()

	// reading from env file
	viper.SetConfigName("app")
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		l.Info("Error reading .env file")
		l.Panic(err.Error())
	}

	db, err := data.GetDb()
	if err != nil {
		l.Info("Error connecting to db")
		l.Fatal(err.Error())
	}

	repo := utils.Repository{db}
	repo.DB.AutoMigrate(&data.User{}, &data.Session{})

	c := server.NewAuthGrpcServer(handlers.NewProvider(&repo, l))
	pb.RegisterAuthServer(gs, c)

	reflection.Register(gs)

	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		os.Exit(1)
	}

	gs.Serve(lis)
}
