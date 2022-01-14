package core

import (
	"algorithm-test/config"
	myErr "algorithm-test/error"
	log "github.com/sirupsen/logrus"
	"strconv"
	"sync"
	"time"
)

var (
	set          = make(map[uint32]*IstsInfo, 0)
	mutex        sync.RWMutex
	AllIsts      = make([]*IstsInfo, 0, 0)
	instanceList = make([]uint32, 0, 0)
)

type IstsInfo struct {
	pos  uint32
	ch   chan bool
	Conn *ClientConn
}

type ClientConn struct {
	ClientAddr string
	TxnNum     int
	VirtualNum uint32
	Memory     uint32
}

type ClientStatus struct {
	Online bool
	Port   int32
	Memory uint32
}

func RegMsgHandle(msg *ClientStatus, clientAddr string) error {
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
}

func addInstance(clientAddr string, msg *ClientStatus) error {
	clientChan := make(chan bool, 0)
	log.Infof("msg.GetMemory() is: %d\n", msg.Memory)
	virtualNum := msg.Memory >> 7 // ➗128
	log.Infof("VirtualNum is: %d\n", virtualNum)
	conn := &ClientConn{
		ClientAddr: clientAddr,
		TxnNum:     0,
		VirtualNum: virtualNum,
		Memory:     msg.Memory,
	}
	log.Infof("Connect to instance at %s success.", clientAddr)
	go countDownTime(clientChan, clientAddr, virtualNum)
	var istsInfo *IstsInfo
	for i := 0; i <= int(virtualNum); i++ {
		hash := hashCode(clientAddr + "x" + strconv.Itoa(i))
		istsInfo = &IstsInfo{
			pos:  hash,
			ch:   clientChan,
			Conn: conn,
		}
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
	AllIsts = append(AllIsts, istsInfo)
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

			break LOOP
		}
	}
}
