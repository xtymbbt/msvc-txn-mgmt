package test

import (
	"algorithm-test/core"
	"fmt"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"strconv"
	"time"
)

func LoadBalanceTest(实例数目 int, 事务数目 int) {
	// 先进行实例注册，注册若干实例
	// 生成实例
	for i := 0; i < 实例数目; i++ {
		rand.Seed(time.Now().UnixNano())
		clientstatus := &core.ClientStatus{
			Online: true,
			Port:   rand.Int31n(65535),
			//Memory: uint32((rand.Intn(31) + 1) * 1024),
			Memory: uint32(12800),
		}
		log.Infof("clientstatus is: %v\n", clientstatus)
		clientAddr := strconv.Itoa(rand.Intn(255)) + "." +
			strconv.Itoa(rand.Intn(255)) + "." +
			strconv.Itoa(rand.Intn(255)) + "." +
			strconv.Itoa(rand.Intn(255))
		log.Infof("clientAddr is: %s:%d\n", clientAddr, clientstatus.Port)
		core.RegMsgHandle(clientstatus, clientAddr)
	}
	// 然后进行事务信息传递
	// 需要虚拟出若干事务
	for i := 0; i < 事务数目; i++ {
		txnMessage := &core.TxnMessage{
			Online: true,
		}
		txnMessage.TreeUuid = uuid.NewV1().String()
		core.RouteMessage(txnMessage)
	}
	fmt.Printf("Memory\t|VirtualNum\t|TxnNum\n")
	for i := 0; i < len(core.AllIsts); i++ {
		conn := core.AllIsts[i].Conn
		fmt.Printf("%d\t|%d\t\t|%d\t\n",
			conn.Memory, conn.VirtualNum+1, conn.TxnNum)
	}
}
