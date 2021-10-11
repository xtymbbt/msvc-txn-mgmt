package server

import (
	"RegCenter/core"
	"RegCenter/proto/cluster"
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc/peer"
	"net"
	"strconv"
	"strings"
)

// dataCenterServer.go

type RegCenterServer struct{}

func (s *RegCenterServer) HealthCheck(ctx context.Context, in *cluster.ClientStatus) (*cluster.RegCenterStatus, error) {
	log.Debugf("server received message: %v", in)
	clientIP, err := getClientIP(ctx)
	clientIP = clientIP + ":" + strconv.Itoa(int(in.Port))
	log.Infof("clientIP is %s\n", clientIP)
	if err != nil {
		log.Errorf("error occured in HealthCheck: %#v\n", err)
		return nil, err
	}
	err = core.RegMsgHandle(in, clientIP)
	if err != nil {
		return &cluster.RegCenterStatus{Online: false}, err
	}
	return &cluster.RegCenterStatus{
		Online: true,
	}, nil
}

func getClientIP(ctx context.Context) (string, error) {
	pr, ok := peer.FromContext(ctx)
	if !ok {
		return "", fmt.Errorf("[getClientIP] invoke FromContext() failed")
	}
	if pr.Addr == net.Addr(nil) {
		return "", fmt.Errorf("[getClientIP] peer.Addr is nil")
	}
	ip := strings.Split(pr.Addr.String(), ":")[0]
	return ip, nil
}
