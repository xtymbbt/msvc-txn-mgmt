package golangApplication

import (
	"data-center-v2/config"
	_ "data-center-v2/database"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"strconv"
)

func Run() {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(config.PORT))
	if err != nil {
		log.Fatalf("DataCenterApplication failed to listen at PORT : %d.\nerror is:\n%v\n", config.PORT, err)
	}
	s := grpc.NewServer()
	register(s)
	reflection.Register(s)
	log.Infof("DataCenterApplication has successfully started at PORT : %d", config.PORT)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("DataCenterApplication has encountered some problems and failed to serve.\nerror is:\n%v\n", err)
	}
	//s.Serve(lis)
}
