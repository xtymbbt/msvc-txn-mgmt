package golangApplication

import (
	"data-center-v2/proto/execTxnRpc"
	"data-center-v2/server"
	"google.golang.org/grpc"
)

func register(s *grpc.Server) {
	execTxnRpc.RegisterExecTxnRpcServer(s, &server.Server{})
}
