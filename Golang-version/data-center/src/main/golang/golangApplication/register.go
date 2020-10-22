package golangApplication

import (
	"../proto"
	"../server"
	"google.golang.org/grpc"
)

func register(s *grpc.Server) {
	commonInfo.RegisterCommonInfoServer(s, &server.Server{})
}
