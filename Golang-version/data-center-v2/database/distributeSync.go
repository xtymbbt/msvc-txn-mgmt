package database

import (
	log "github.com/sirupsen/logrus"
)

func distributeSync() {
	// 使用主数据库对从数据库进行同步。
	for _, db := range bakDBs {
		// 首先判断数据库同步与否。
		// 若不一致，则进行同步。
		syncDBs(mainDB, db)
	}
	// 在case 2的情况下，所有steps执行完毕后，将datacenter数据库中的backup信息修改为0，代表无需进行同步。
	err := updateDataCenterDB(0)
	if err != nil {
		log.Fatalf("Sync backup databases using main database failed. Error is: %#v", err)
	}
}
