package server

import (
	"../handleMessage"
	"../proto"
	"fmt"
	"golang.org/x/net/context"
)

// server.go

type Server struct{}

func (s *Server) SendToDataCenter(ctx context.Context, in *commonInfo.HttpRequest) (*commonInfo.HttpResponse, error) {
	fmt.Printf("server recieved message: %#v\n\n", in)
	err := handleMessage.HandleMessage(in)
	return &commonInfo.HttpResponse{Success: err == nil}, err
}
