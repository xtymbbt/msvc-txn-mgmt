package database

import (
	log "github.com/sirupsen/logrus"
)

func voteSync() {
	// TODO: 投票式同步
	// Steps: Find same databases which offers the most number.

	// if all are the same, return

	// if not, sync other databases.

	// 在case 1的情况下，所有steps执行完毕后，将datacenter数据库中backup的信息修改为0，代表无需进行同步。
	err := updateDataCenterDB(0)
	if err != nil {
		log.Fatalf("Sync backup databases using main database failed. Error is: %#v", err)
	}
}

