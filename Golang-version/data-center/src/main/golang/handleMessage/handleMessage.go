package handleMessage

import (
	"../../../resources/config"
	"../database"
	myErr "../error"
	"../proto"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

var hashMap = make(map[string][]*commonInfo.HttpRequest, 0)
var canWriteMap = make(map[string]map[uint32]map[string][]uint32, 0)
var numMap = make(map[string][]int64, 0)
var msgMap = make(map[string]int64, 0)
var timeMap = make(map[string]chan bool, 0)
var mutex sync.Mutex

func HandleMessage(message *commonInfo.HttpRequest) (err error) {
	mutex.Lock()
	if _, ok := hashMap[message.TreeUuid]; ok {
		hashMap[message.TreeUuid] = append(hashMap[message.TreeUuid], message)
		msgMap[message.TreeUuid]++
		if _, ok := canWriteMap[message.TreeUuid][message.Pos]; ok {
			if _, ok := canWriteMap[message.TreeUuid][message.Pos][message.ServiceUuid]; !ok {
				canWriteMap[message.TreeUuid][message.Pos][message.ServiceUuid] = []uint32{message.MapperNum, message.ServiceNum}
			}
		} else {
			canWriteMap[message.TreeUuid][message.Pos] = map[string][]uint32{message.ServiceUuid: {message.MapperNum, message.ServiceNum}}
		}
	} else {
		hashMap[message.TreeUuid] = []*commonInfo.HttpRequest{message}
		msgMap[message.TreeUuid] = 1
		canWriteMap[message.TreeUuid] = map[uint32]map[string][]uint32{message.Pos: {message.ServiceUuid: []uint32{message.MapperNum, message.ServiceNum}}}
		timeMap[message.TreeUuid] = make(chan bool, 1)
		go timeOut(message.TreeUuid, &err)
	}


	if canWrite(canWriteMap[message.TreeUuid], message.TreeUuid, msgMap[message.TreeUuid]) {
		log.Info("message all received, writing into database...")
		if _, ok := timeMap[message.TreeUuid]; ok {
			timeMap[message.TreeUuid] <- true
			err = database.Write(hashMap[message.TreeUuid])
			if err != nil {
				log.Error("write into database failed, deleting caches...")
			} else {
				log.Info("write into database success, deleting caches...")
			}
		} else {
			log.Error("current service chain already timed out. deleting caches...")
		}
		if _, ok := hashMap[message.TreeUuid]; ok {
			delete(hashMap, message.TreeUuid)
			delete(canWriteMap, message.TreeUuid)
			delete(numMap, message.TreeUuid)
			delete(msgMap, message.TreeUuid)
			delete(timeMap, message.TreeUuid)
			log.Info("caches deleted.")
		} else {
			log.Errorf("no caches found. cannot delete caches at TreeUUID: %s", message.TreeUuid)
		}
	}
	mutex.Unlock()
	return err
}

func timeOut(treeUuid string, err *error){
	mutex.Lock()
	timeChan := timeMap[treeUuid]
	mutex.Unlock()
	select {
	case <- timeChan:
		log.Info("receiving message succeeded, timeOut function stopped, writing into database.")
	case <- time.After(time.Second * config.TIMELAPSES):
		log.Error("receiving message timed out, deleting caches...")
		if _, ok := hashMap[treeUuid]; ok {
			delete(hashMap, treeUuid)
			delete(canWriteMap, treeUuid)
			delete(numMap, treeUuid)
			delete(msgMap, treeUuid)
			delete(timeMap, treeUuid)
			log.Info("caches deleted.")
			*err = myErr.NewError(300, "receive message timed out "+string(rune(config.TIMELAPSES))+" seconds.")
		}
	}
}

func canWrite(canWriteMap map[uint32]map[string][]uint32, treeUuid string, msgNum int64) bool {
	if _, ok := canWriteMap[0]; !ok {
		return false
	}
	var (
		mapperONum  int64 = 0
		serviceONum int64 = 0
	)
	for _, num := range canWriteMap[0] {
		mapperONum += int64(num[0])
		serviceONum += int64(num[1])
	}
	//                          mapperNum, serviceNum
	numMap[treeUuid] = []int64{mapperONum, serviceONum}
	var i uint32 = 1
	for {
		if _, ok := canWriteMap[i]; ok {
			var (
				mapperNum  int64 = 0
				serviceNum int64 = 0
			)
			for _, num := range canWriteMap[i] {
				mapperNum += int64(num[0])
				serviceNum += int64(num[1]) - 1
			}
			numMap[treeUuid][0] += mapperNum
			numMap[treeUuid][1] += serviceNum
		} else {
			break
		}
		i++
	}
	//fmt.Printf("numMap is: %#v\nmsgNum is %v\n", numMap, msgNum)
	if numMap[treeUuid][0] == msgNum && numMap[treeUuid][1] == 0 {
		return true
	}
	return false
}
