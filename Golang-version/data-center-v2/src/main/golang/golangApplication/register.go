package golangApplication

import (
	"../proto/commonInfo"
	"../server"
	"google.golang.org/grpc"
)

func register(s *grpc.Server) {
	commonInfo.RegisterCommonInfoServer(s, &server.Server{})
}
