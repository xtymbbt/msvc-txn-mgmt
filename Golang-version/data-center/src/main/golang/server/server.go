package server

import (
	"../handleMessage"
	"../proto"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// server.go

type Server struct{}

func (s *Server) SendToDataCenter(ctx context.Context, in *commonInfo.HttpRequest) (*commonInfo.HttpResponse, error) {
	log.Infof("server received message: %#v", in)
	err := handleMessage.HandleMessage(in)
	return &commonInfo.HttpResponse{Success: err == nil}, err
}
