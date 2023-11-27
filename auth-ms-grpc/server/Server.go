package server

import (
	"auth-ms/auth-ms-grpc/protos/pb"
	"context"
	"fmt"
)

type Server struct {
	pb.AuthServer
}

func (a *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	fmt.Println(req)
	return &pb.LoginResponse{
		Status: true,
		Msg:    "dpfgoeirh osehg ehir;gos hoeigh ",
		Data: &pb.Tokens{
			AccessToken:  "sgergerh",
			RefreshToken: "rthrth",
			Uid:          "rthwrthw",
		},
		Errors: nil,
	}, nil
}

func (a *Server) Signup(ctx context.Context, req *pb.SignupRequest) (*pb.SignupResponse, error) {
	return &pb.SignupResponse{
		Status: true,
		Msg:    "",
		Data: &pb.Tokens{
			AccessToken:  "rthw45hwerh",
			RefreshToken: "werhtw45h",
			Uid:          "rthw45hw",
		},
		Errors: nil,
	}, nil
}
