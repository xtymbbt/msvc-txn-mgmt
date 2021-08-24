package server

import (
	"data-center-v2/handleMessage"
	"data-center-v2/proto/execTxnRpc"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// server.go

type Server struct{}

func (s *Server) ExecTxn(ctx context.Context, in *execTxnRpc.TxnMessage) (*execTxnRpc.TxnStatus, error) {
	log.Infof("server received message: %#v", in)
	err := handleMessage.HandleMessage(in)
	if err != nil {
		return &execTxnRpc.TxnStatus{
			Status:  500,
			Message: "Transaction execute failed.",
		}, err
	}
	return &execTxnRpc.TxnStatus{
		Status:  200,
		Message: "Transaction execute success",
	}, err
}
