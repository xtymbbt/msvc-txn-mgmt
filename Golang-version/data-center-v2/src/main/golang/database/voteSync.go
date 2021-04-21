package database

import (
	"../../../resources/config"
	"database/sql"
	log "github.com/sirupsen/logrus"
)

func voteSync() {
	// 投票式同步
	// Steps: Find same databases which offers the most number.
	// if all are the same, return
	// if not, sync other databases.
	// -------------------------------------------------------
	// First, find vote map.
	var dataBaseVoteMap = make([][]map[string]*sql.DB, 0)
	dbs := make([]map[string]*sql.DB, len(config.DBBakUrls))
	copy(dbs, bakDBs)
	dbs = append(dbs, mainDB)
	var sameMap = make(map[int]bool, 0)
	for i, db := range dbs {
		if _, ok := sameMap[i]; ok {
			continue
		}
		oneVoteMap := make([]map[string]*sql.DB, 0)
		oneVoteMap = append(oneVoteMap, db)
		for j := i + 1; j < len(dbs); j++ {
			if _, ok := sameMap[j]; ok {
				continue
			}
			if same := compareDBs(db, dbs[j]); same {
				sameMap[j] = true
				oneVoteMap = append(oneVoteMap, dbs[j])
			}
		}
		dataBaseVoteMap = append(dataBaseVoteMap, oneVoteMap)
	}
	// Second, find the most number in vote map & find src database and dns database.
	maxIndex, maxNum := 0, 0
	for i := range dataBaseVoteMap {
		if len(dataBaseVoteMap[i]) > maxNum {
			maxNum = len(dataBaseVoteMap[i])
			maxIndex = i
		}
	}
	syncSrcDB := dataBaseVoteMap[maxIndex][0]
	syncDnsDB := make([]map[string]*sql.DB, 0)
	for i, dbs := range dataBaseVoteMap {
		if i == maxIndex {
			continue
		}
		for _, db := range dbs {
			syncDnsDB = append(syncDnsDB, db)
		}
	}
	// Third, use most number's database to sync other databases.
	for _, dnsDB := range syncDnsDB {
		syncDBs(syncSrcDB, dnsDB)
	}
	// --------------------------------------------------------------------------------------
	// 在case 1的情况下，所有steps执行完毕后，将datacenter数据库中backup的信息修改为0，代表无需进行同步。
	err := updateDataCenterDB(0)
	if err != nil {
		log.Fatalf("Sync backup databases using main database failed. Error is: %#v", err)
	}
}

func compareDBs(db1, db2 map[string]*sql.DB) bool {
	// compare two database.
	for _, dbname := range config.DBNAME {
		db1Tables := make(map[string]bool, 0)
		db2Tables := make(map[string]bool, 0)
		queryTableName(&db1Tables, db1[dbname])
		queryTableName(&db2Tables, db2[dbname])
		for k := range db1Tables {
			if _, ok := db2Tables[k]; !ok {
				return false
			}
		}
		for k := range db2Tables {
			if _, ok := db1Tables[k]; !ok {
				return false
			}
		}
		// Query SrcDB's data and Record them into a map.
		db1Data := make(map[string][][]interface{}, 0)
		queryDBData(db1[dbname], &db1Data, db1Tables)
		// Query DnsDB's data and Record them into a map.
		db2Data := make(map[string][][]interface{}, 0)
		queryDBData(db2[dbname], &db2Data, db2Tables)
		for table := range db1Tables {
			if len(db1Data[table]) != len(db2Data[table]) {
				return false
			}
			for i, values := range db1Data[table] {
				for j, value := range values {
					if db2Data[table][i][j] == value {
						continue
					} else {
						return false
					}
				}
			}
		}
	}
	return true
}
