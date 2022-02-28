package golangApplication

import (
	"RegCenter/config"
	"RegCenter/proto/cluster"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"strconv"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func Run() {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(config.PORT))
	if err != nil {
		log.Fatalf("DataCenterApplication failed to listen at PORT : %d.\nerror is:\n%v\n", config.PORT, err)
	}
	if config.ClusterEnabled { // 此处表示的是注册中心集群
		selfIP := "127.0.0.1:" + strconv.Itoa(config.PORT)
		if len(config.RegNodes) > 0 && selfIP != config.RegNodes[0] {
			conn, err := grpc.Dial(config.RegNodes[0], grpc.WithInsecure())
			if err != nil {
				log.Fatalf("Cannot connect to register center: %v\n"+
					"Error is: %v\n", config.RegNodes[0], err)
			}
			log.Infof("Connect to register center at %s success.", config.RegNodes[0])
			defer func(conn *grpc.ClientConn) {
				err := conn.Close()
				if err != nil {
					log.Fatalf("Close connection to register center failed.\nError is: %#v\n:", err)
				}
			}(conn)
			go healthCheck(conn)
		}
	}
	wg.Add(2)
	go regCenterServer(lis)
	go dataCenterServer(lis)
	wg.Wait()
}

func healthCheck(conn *grpc.ClientConn) {
	c := cluster.NewHealthCheckClient(conn)
	check, err := c.HealthCheck(context.Background(), &cluster.ClientStatus{
		Online: true,
		Port:   int32(config.PORT),
	})
	if err != nil {
		log.Errorf("There is an error occured during checking health with register center.\n"+
			"Error is: %v\n", err)
		for err != nil {
			err = conn.Close()
			if err != nil {
				log.Fatalf("Close connection to register center failed.\nError is: %#v\n:", err)
			}
			conn, err = grpc.Dial(config.RegNodes[0], grpc.WithInsecure())
			if err != nil {
				log.Fatalf("Cannot connect to register center: %v\n"+
					"Error is: %v\n", config.RegNodes[0], err)
			}
			log.Infof("Connect to register center at %s success.", config.RegNodes[0])
			c = cluster.NewHealthCheckClient(conn)
			check, err = c.HealthCheck(context.Background(), &cluster.ClientStatus{
				Online: true,
				Port:   int32(config.PORT),
			})
			if err != nil {
				log.Errorf("There is an error occured during checking health with register center.\n"+
					"Error is: %v\n", err)
			}
		}
	}
	if !check.Online {
		log.Fatalln("There is an error occured during checking health with register center.")
	}
	log.Infoln("Health check success.")
	for true {
		time.Sleep(time.Second * time.Duration(config.HealthCheckTime/3)) // 默认每3秒进行一次健康检查，可配置
		hltChk(c)
	}
}

func hltChk(c cluster.HealthCheckClient) {
	check, err := c.HealthCheck(context.Background(), &cluster.ClientStatus{
		Online: true,
		Port:   int32(config.PORT),
	})
	if err != nil {
		log.Errorf("There is an error occured during checking health with register center.\n"+
			"Error is: %v\n", err)
	}
	if !check.Online {
		log.Fatalln("There is an error occured during checking health with register center.")
	}
	log.Infoln("Health check success.")
}
