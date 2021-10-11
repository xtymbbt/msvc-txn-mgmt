package golangApplication

import (
	"data-center-v2/config"
	_ "data-center-v2/database"
	"data-center-v2/proto/cluster"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"strconv"
	"time"
)

func Run() {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(config.PORT))
	if err != nil {
		log.Fatalf("DataCenterApplication failed to listen at PORT : %d.\nerror is:\n%v\n", config.PORT, err)
	}
	if config.ClusterEnabled { // 此处表示的是事务管理中心集群
		conn, err := grpc.Dial(config.RegisterCenter, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Cannot connect to register center: %v\n"+
				"Error is: %v\n", config.RegisterCenter, err)
		}
		log.Infof("Connect to register center at %s success.", config.RegisterCenter)
		defer func(conn *grpc.ClientConn) {
			err := conn.Close()
			if err != nil {
				log.Fatalf("Close connection to register center failed.\nError is: %#v\n:", err)
			}
		}(conn)
		c := cluster.NewHealthCheckClient(conn)
		go healthCheck(c)
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

func healthCheck(c cluster.HealthCheckClient) {
	for true {
		time.Sleep(time.Second * time.Duration(config.HealthCheckTime/10)) // 默认每3秒进行一次健康检查，可配置
		hltChk(c)
	}
}

func hltChk(c cluster.HealthCheckClient) {
	check, err := c.HealthCheck(context.Background(), &cluster.ClientStatus{
		Online: true,
		Port:   int32(config.PORT),
	})
	if err != nil {
		log.Fatalf("There is an error occured during checking health with register center.\n"+
			"Error is: %v\n", err)
	}
	if !check.Online {
		log.Fatalln("There is an error occured during checking health with register center.")
	}
	log.Infoln("Health check success.")
}
