package golangApplication

import (
	"RegCenter/config"
	log "github.com/sirupsen/logrus"
	"net"
	"strconv"
	"sync"
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
		//TODO:增加注册中心集群，可能做不完
	}
	wg.Add(2)
	go regCenterServer(lis)
	go dataCenterServer(lis)
	wg.Wait()
}
