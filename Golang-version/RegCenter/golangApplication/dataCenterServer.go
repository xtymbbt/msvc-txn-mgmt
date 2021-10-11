package golangApplication

import (
	"RegCenter/config"
	"RegCenter/proto/execTxnRpc"
	"RegCenter/server"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func dataCenterServer(lis net.Listener) {
	s := grpc.NewServer()
	execTxnRpc.RegisterExecTxnRpcServer(s, &server.DataCenterServer{})
	reflection.Register(s)
	log.Infof("DataCenterApplication has successfully started at PORT : %d", config.PORT)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("DataCenterApplication has encountered some problems and failed to serve.\nerror is:\n%v\n", err)
	}
	//s.Serve(lis)
	wg.Done()
}
