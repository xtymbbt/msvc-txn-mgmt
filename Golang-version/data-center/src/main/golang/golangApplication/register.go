package golangApplication

import (
	"../proto"
	"google.golang.org/grpc"
	"../server"
)

func register(s *grpc.Server) {
	commonInfo.RegisterCommonInfoServer(s, &server.Server{})
}
