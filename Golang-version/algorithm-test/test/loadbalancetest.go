package test

import (
	"algorithm-test/core"
	"fmt"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func LoadBalanceTest(实例数目 int, 事务数目 int) {
	// 先进行实例注册，注册若干实例
	// 生成实例
	for i := 0; i < 实例数目; i++ {
		rand.Seed(time.Now().UnixNano())
		//memory := uint32((rand.Intn(31) + 1) * 1024)
		memory := uint32(12800)
		if 实例数目 == 2 {
			if i == 0 {
				memory = uint32(8192)
			} else if i == 1 {
				memory = uint32(23552)
			}
		} else if 实例数目 == 3 {
			if i == 0 {
				memory = uint32(5120)
			} else if i == 1 {
				memory = uint32(25600)
			} else if i == 2 {
				memory = uint32(23552)
			}
		} else if 实例数目 == 4 {
			if i == 0 {
				memory = uint32(12288)
			} else if i == 1 {
				memory = uint32(18432)
			} else if i == 2 {
				memory = uint32(16384)
			} else if i == 3 {
				memory = uint32(14336)
			}
		} else if 实例数目 == 5 {
			if i == 0 {
				memory = uint32(24576)
			} else if i == 1 {
				memory = uint32(6144)
			} else if i == 2 {
				memory = uint32(17408)
			} else if i == 3 {
				memory = uint32(10240)
			} else if i == 4 {
				memory = uint32(4096)
			}
		}
		clientstatus := &core.ClientStatus{
			Online: true,
			Port:   rand.Int31n(65535),
			Memory: memory,
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
	var (
		avg int64 = 0
		min int64 = math.MaxInt64
		max int64 = math.MinInt64
	)
	for i := 0; i < 事务数目; i++ {
		txnMessage := &core.TxnMessage{
			Online: true,
		}
		txnMessage.TreeUuid = uuid.NewV1().String()
		start := time.Now().UnixNano()
		core.RouteMessage(txnMessage)
		end := time.Now().UnixNano()
		cost := end - start
		if cost < min && cost != 0 {
			min = cost
		}
		if cost > max {
			max = cost
		}
		avg += cost
	}
	avg /= int64(事务数目)
	fmt.Printf("min\t|avg\t|max\n")
	fmt.Printf("%dns\t|%dns\t|%dns\n", min, avg, max)
	fmt.Printf("Memory\t|VirtualNum\t|TxnNum\n")
	for i := 0; i < len(core.AllIsts); i++ {
		conn := core.AllIsts[i].Conn
		fmt.Printf("%d\t|%d\t\t|%d\t\n",
			conn.Memory, conn.VirtualNum+1, conn.TxnNum)
	}
}
