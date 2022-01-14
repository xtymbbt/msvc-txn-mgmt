package core

import (
	"RegCenter/config"
	myErr "RegCenter/error"
	"RegCenter/proto/cluster"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"strconv"
	"sync"
	"time"
)

var (
	set          = make(map[uint32]*IstsInfo, 0)
	mutex        sync.RWMutex
	instanceList = make([]uint32, 0, 0)
)

type IstsInfo struct {
	pos  uint32
	ch   chan bool
	conn *grpc.ClientConn
}

func RegMsgHandle(msg *cluster.ClientStatus, clientAddr string) error {
	if msg.Online {
		mutex.Lock()
		hash := hashCode(clientAddr + "x" + "0")
		_, ok := set[hash]
		if ok {
			clientChan := set[hash].ch
			err := sendChanMsg(clientChan)
			if err != nil { // 代表着channel已经关闭了，即，上一次健康检查已失败
				err := addInstance(clientAddr, msg)
				if err != nil {
					return err
				}
			}
		} else {
			err := addInstance(clientAddr, msg)
			if err != nil {
				return err
			}
		}
		mutex.Unlock()
		return nil
	} else {
		return myErr.NewError(400, "client is not online?")
	}
}

func addInstance(clientAddr string, msg *cluster.ClientStatus) error {
	clientChan := make(chan bool, 0)
	log.Debugf("msg.GetMemory() is: %d\n", msg.GetMemory())
	virtualNum := msg.GetMemory() >> 7 // ➗128
	log.Debugf("VirtualNum is: %d\n", virtualNum)
	conn, err := grpc.Dial(clientAddr, grpc.WithInsecure())
	if err != nil {
		log.Errorf("Cannot connect to instance: %v\n"+
			"Error is: %v\n", clientAddr, err)
	}
	log.Infof("Connect to instance at %s success.", clientAddr)
	go countDownTime(clientChan, clientAddr, virtualNum)
	var istsInfo *IstsInfo
	for i := 0; i <= int(virtualNum); i++ {
		hash := hashCode(clientAddr + "x" + strconv.Itoa(i))
		istsInfo = &IstsInfo{
			pos:  hash,
			ch:   clientChan,
			conn: conn,
		}
		log.Debugf("This instance's hashcode is: %d\n", hash)
		set[hash] = istsInfo

		if len(instanceList) == 0 {
			instanceList = append(instanceList, hash)
		} else {
			if hash <= instanceList[0] {
				instanceList = append([]uint32{hash}, instanceList...)
			} else {
				writen := false
				for j := 1; j < len(instanceList); j++ {
					if hash <= instanceList[j] {
						x := make([]uint32, 0, 0)
						x = append(x, instanceList[:j]...)
						x = append(x, hash)
						instanceList = append(x, instanceList[j:]...)
						writen = true
						break
					}
				}
				if !writen {
					instanceList = append(instanceList, hash)
				}
			}
		}
	}
	return nil
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

func countDownTime(clientChan chan bool, clientAddr string, virtualNum uint32) {
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
			var istsInfo *IstsInfo
			hash := hashCode(clientAddr + "x" + "0")
			istsInfo, ok := set[hash]
			if ok {
				log.Infoln("Closing connection...")
				err := istsInfo.conn.Close()
				if err != nil {
					log.Errorf("Closing grpc connection to %s failed. This may cause some instabilities of your system. Please check.", clientAddr)
				} else {
					log.Infoln("Connection closed.")
				}
				for i := 0; i < int(virtualNum); i++ {
					hash = hashCode(clientAddr + "x" + strconv.Itoa(i))
					delete(set, hash)
					var idx uint32
					for j := 0; j < len(instanceList); j++ {
						if instanceList[j] == hash {
							idx = hash
							break
						}
					}
					instanceList = append(instanceList[:idx], instanceList[idx+1:]...)
				}
			} else {
				log.Errorf("Could not found a corresponding instance which hashCode is: %d\n", hash)
			}
			mutex.Unlock()
			break LOOP
		}
	}
}
