package database

import (
	"../proto"
	"sync"
)

var wg sync.WaitGroup
var mutex sync.RWMutex

func Write(dataS []*commonInfo.HttpRequest) (err error) {
	for _, data := range dataS {
		wg.Add(1)
		go goWrite(data, &err)
	}
	wg.Wait()
	return
}

func goWrite(data *commonInfo.HttpRequest, err *error){
	mutex.Lock()
	defer mutex.Unlock()
	dbx := db[data.DbName]
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
	wg.Done()
}