package database

import (
	log "github.com/sirupsen/logrus"
)


func initSync() (err error) {
	// 初始化时候的区块链同步
	// 先读取datacenter数据库中backup的信息
	dbx := mainDB["data_center"]
	sqlStr := "select backup_state from backup where id = 1"
	queryRow := dbx.QueryRow(sqlStr)
	var backupState int
	err = queryRow.Scan(&backupState)
	if err != nil {
		return err
	}

	switch backupState {
	case 0:
	// 若为0，则无需进行同步。
	// WARN:只是在init的时候需要如此考虑，因为init设置的初衷是失败后进行回滚，而其他的sync设计的初衷是提防数据库被黑客篡改
		log.Info("Backup State is 0, no need to sync.")
	case 1:
	// 若为1，则使用多数投票的方式进行数据库同步。（1代表在主数据库写入的时候发生了错误）
	// 若程序在此步骤中崩溃，由于是采用的多数投票方式进行同步，而多数的数据库不会被更改，只有少数的数据库需要被更改
	// 因此，我们不需要修改该backup的值。因为若崩溃了，我们仍需要采用多数投票的方式进行同步。
		log.Warn("Backup State is 1, executing vote sync...")
		voteSync()
		log.Warn("Sync succeeded.")
	case 2:
	// 若为2，则直接使用主数据库对从数据库进行同步。（2代表在从数据库写入的时候发生了错误）
	//       此同步需要的是读取注数据库中所有的信息，与从数据库进行同步。
	// 若程序在此步骤中崩溃，仍然无需更改backup的值，因为再重新启动的时候，我们仍需使用主数据库对从数据库进行同步。
		log.Warn("Backup State is 1, executing vote sync...")
		distributeSync()
		log.Warn("Sync succeeded.")
	default:
		log.Fatalf("Backup State is neither 0 or 1 or 2, init failed. Please check your database and your program logs.")
	}



	return nil
}
