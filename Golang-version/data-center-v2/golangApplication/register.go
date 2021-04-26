package golangApplication

import (
	"data-center-v2/proto/commonInfo"
	"data-center-v2/server"
	"google.golang.org/grpc"
)

func register(s *grpc.Server) {
	commonInfo.RegisterCommonInfoServer(s, &server.Server{})
}
