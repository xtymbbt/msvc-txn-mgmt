package database

import (
	"data-center-v2/common"
	"data-center-v2/proto/execTxnRpc"
	"database/sql"
	log "github.com/sirupsen/logrus"
)

func dbBackup(root *common.TreeNode) (err error) {
	// 设置一个datacenter数据库，用于记录Backup的信息。0代表没有出现错误。1代表在主数据库写入时出现的错误，2代表在备份数据库写入时出现的错误
	// 每次在backup之前，先更改数据库该backup值为2，在backup结束之后，更改数据库该backup值为0
	log.Warn("WRITING INTO BACKUP DATABASES...")
	err = updateDataCenterDB(2)
	if err != nil {
		log.Fatalf("record datacenter state failed.\n"+
			"error is: %#v\n"+
			"Stopping datacenter...", err)
	}
	log.Info("Record Datacenter State succeeded.")
	//=====Deprecated======
	//for _, database := range bakDBs {
	//	for _, data := range dataS {
	//		wg.Add(1)
	//		go goBackup(data, database, &err)
	//	}
	//}
	//wg.Wait()
	//=========end=========
	//=======new way=======
	for _, database := range bakDBs {
		if root != nil {
			err = startDBTX(database[root.DbName], root)
		}
	}
	if err != nil {
		log.Fatalf("Executing Database Transaction failed.\n"+
			"error is: %#v\n", err)
	}
	//=========end=========
	log.Warn("WRITING INTO BACKUP DATABASES SUCCEEDED.")
	err = updateDataCenterDB(0)
	if err != nil {
		log.Fatalf("record datacenter state failed.\n"+
			"error is: %#v\n"+
			"Stopping datacenter...", err)
	}
	log.Info("Record Datacenter State succeeded.")
	return nil
}

func goBackup(data *execTxnRpc.TxnMessage, database map[string]*sql.DB, err *error) {
	mutex.Lock()
	dbx := database[data.DbName]
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
