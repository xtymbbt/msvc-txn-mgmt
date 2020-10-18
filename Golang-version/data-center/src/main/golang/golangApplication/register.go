package golangApplication

import (
	"../proto"
	"google.golang.org/grpc"
	"../controller"
)

func register(s *grpc.Server) {
	commonInfo.RegisterCommonInfoServer(s, &controller.Server{})
}
