package database

import (
	"../proto"
	log "github.com/sirupsen/logrus"
	"sync"
)

var wg sync.WaitGroup
var mutex sync.RWMutex

func Write(dataS []*commonInfo.HttpRequest) (err error) {
	// 设置一个datacenter数据库，用于记录Backup的信息。0代表没有出现错误。1代表在主数据库写入时出现的错误，2代表在备份数据库写入时出现的错误
	// 每次在写入之前，先更改数据库该backup值为1，在写入结束之后，更改数据库该backup值为0
	log.Warn("WRITING INTO MAIN DATABASE...")
	err = updateDataCenterDB(1)
	if err != nil {
		log.Fatalf("record datacenter state failed.\n"+
			"error is: %#v\n"+
			"Stopping datacenter...", err)
	}
	log.Info("Record Datacenter State succeeded.")
	//=====Deprecated======
	//for _, data := range dataS {
	//	wg.Add(1)
	//	go goWrite(data, &err)
	//}
	//=========end=========
	//=======new way=======
	sqlStrS := make([]string, len(dataS))
	signalChan := make([]bool, 0, len(dataS))
	//fmt.Println(len(signalChan))
	for i, data := range dataS {
		go goWriteTX(data, &sqlStrS[i], &err, &signalChan)
	}
	//fmt.Println(len(signalChan))
	for len(signalChan) != len(dataS) {
		//fmt.Println(len(signalChan))
	}
	if err != nil {
		log.Fatalf("Generate SQL str failed.\n"+
			"error is: %#v\n", err)
	}
	if dataS[0] != nil {
		mutex.Lock()
		startDBTX(mainDB[dataS[0].DbName], dataS, sqlStrS, &err)
		mutex.Unlock()
		if err != nil {
			log.Fatalf("Executing Database Transaction failed.\n"+
				"error is: %#v\n", err)
		}
	}
	//=========end=========
	log.Warn("WRITING INTO MAIN DATABASE SUCCEEDED.")
	err = updateDataCenterDB(0)
	if err != nil {
		log.Fatalf("record datacenter state failed.\n"+
			"error is: %#v\n"+
			"Stopping datacenter...", err)
	}
	log.Info("Record Datacenter State succeeded.")
	err = dbBackup(dataS, sqlStrS)
	return
}

func goWriteTX(data *commonInfo.HttpRequest, sqlStr *string, err *error, signalChan *[]bool) {
	/**
	 * 根据data的两个method判断是增删改查的哪个操作
	 * true true = 增
	 * true false = 删
	 * false true = 改
	 * false false = 查
	 */
	if data.Method1 {
		if data.Method2 {
			*sqlStr, *err = dbInsertTX(data.TableName, data.Data)
		} else {
			*sqlStr, *err = dbDeleteTX(data.TableName, data.Data, data.Query)
		}
	} else {
		if data.Method2 {
			*sqlStr, *err = dbUpdateTX(data.TableName, data.Data, data.Query)
		} else {
			*sqlStr, *err = dbQueryTX(data.TableName, data.Data)
		}
	}
	if *err != nil {
		log.Fatalf("database write in failed.\n"+
			"error is: %#v\n"+
			"Stopping datacenter...", *err)
	}
	*signalChan = append(*signalChan, true)
}

func goWrite(data *commonInfo.HttpRequest, err *error) {
	mutex.Lock()
	dbx := mainDB[data.DbName]
	/**
	 * 根据data的两个method判断是增删改查的哪个操作
	 * true true = 增
	 * true false = 删
	 * false true = 改
	 * false false = 查
	 */
	if data.Method1 {
		if data.Method2 {
			*err = dbInsert(dbx, data.TableName, data.Data)
		} else {
			*err = dbDelete(dbx, data.TableName, data.Data, data.Query)
		}
	} else {
		if data.Method2 {
			*err = dbUpdate(dbx, data.TableName, data.Data, data.Query)
		} else {
			*err = dbQuery(dbx, data.TableName, data.Data)
		}
	}
	if *err != nil {
		log.Fatalf("database write in failed.\n"+
			"error is: %#v\n"+
			"Stopping datacenter...", *err)
	}
	mutex.Unlock()
	wg.Done()
}
