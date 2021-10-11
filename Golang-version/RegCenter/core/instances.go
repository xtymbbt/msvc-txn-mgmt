package core

import (
	"RegCenter/config"
	myErr "RegCenter/error"
	"RegCenter/proto/cluster"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"sync"
	"time"
)

var (
	instances = make([]string, 0, 0)
	unusedPos = make([]int, 0, 0)
	set       = make(map[string]*IstsInfo, 0)
	mutex     sync.RWMutex
)

type IstsInfo struct {
	pos  int
	ch   chan bool
	conn *grpc.ClientConn
}

func RegMsgHandle(msg *cluster.ClientStatus, clientAddr string) error {
	if msg.Online {
		mutex.Lock()
		_, ok := set[clientAddr]
		if ok {
			clientChan := set[clientAddr].ch
			err := sendChanMsg(clientChan)
			if err != nil { // 代表着channel已经关闭了，即，上一次健康检查已失败
				addInstance(clientAddr)
			}
		} else {
			addInstance(clientAddr)
		}
		mutex.Unlock()
		return nil
	} else {
		return myErr.NewError(400, "client is not online?")
	}
}

func addInstance(clientAddr string) {
	clientChan := make(chan bool, 0)
	conn, err := grpc.Dial(clientAddr, grpc.WithInsecure())
	if err != nil {
		log.Errorf("Cannot connect to instance: %v\n"+
			"Error is: %v\n", clientAddr, err)
	}
	log.Infof("Connect to instance at %s success.", clientAddr)
	go countDownTime(clientChan, clientAddr)
	var istsInfo *IstsInfo
	if len(unusedPos) == 0 {
		istsInfo = &IstsInfo{
			pos:  len(instances),
			ch:   clientChan,
			conn: conn,
		}
		instances = append(instances, clientAddr)
		if len(instances) > 1023 {
			panic(myErr.NewError(500, "Too many instances! Please check your system."))
		}
	} else {
		istsInfo = &IstsInfo{
			pos:  unusedPos[0],
			ch:   clientChan,
			conn: conn,
		}
		instances[istsInfo.pos] = clientAddr
		unusedPos = unusedPos[1:]
	}
	set[clientAddr] = istsInfo
}

func sendChanMsg(ch chan bool) (err error) {
	defer func() {
		if recover() != nil {
			err = myErr.NewError(500, "channel has closed.")
		}
	}()
	ch <- true
	return nil
}

func countDownTime(clientChan chan bool, clientAddr string) {
	defer func() {
		log.Infoln("Closing channel...")
		close(clientChan) // 使用完该通道后，必须关闭该通道。GO的GC不会回收通道。
		log.Infoln("Channel closed.")
	}()
LOOP:
	for true {
		select {
		case <-clientChan:
		case <-time.After(time.Duration(config.HealthCheckTime) * time.Second):
			log.Error("health check timed out, removing " + clientAddr + " server.")
			mutex.Lock()
			idx := set[clientAddr].pos
			instances[idx] = "E" // E - error
			unusedPos = append(unusedPos, idx)
			log.Infoln("Closing connection...")
			err := set[clientAddr].conn.Close()
			if err != nil {
				log.Errorf("Closing grpc connection to %s failed. This may cause some instabilities of your system. Please check.", clientAddr)
			} else {
				log.Infoln("Connection closed.")
			}
			delete(set, clientAddr)
			mutex.Unlock()
			break LOOP
		}
	}
}
