package server

import (
	"RegCenter/core"
	"RegCenter/proto/execTxnRpc"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// dataCenterServer.go

type DataCenterServer struct{}

func (s *DataCenterServer) ExecTxn(ctx context.Context, in *execTxnRpc.TxnMessage) (*execTxnRpc.TxnStatus, error) {
	log.Infof("server received message: %#v", in)
	return core.RouteMessage(in)
}
