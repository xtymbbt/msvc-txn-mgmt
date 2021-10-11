package golangApplication

import (
	"RegCenter/config"
	"RegCenter/proto/cluster"
	"RegCenter/server"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func regCenterServer(lis net.Listener) {
	s := grpc.NewServer()
	cluster.RegisterHealthCheckServer(s, &server.RegCenterServer{})
	reflection.Register(s)
	log.Infof("RegCenterApplication has successfully started at PORT : %d", config.PORT)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("RegCenterApplication has encountered some problems and failed to serve.\nerror is:\n%v\n", err)
	}
	//s.Serve(lis)
	wg.Done()
}
