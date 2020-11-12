package database

import (
	"../../../resources/config"
	"database/sql"
	log "github.com/sirupsen/logrus"
)

func syncDBs(srcDB map[string]*sql.DB, dnsDB map[string]*sql.DB) {
	// TODO: 进行数据库同步
	for _, dbname := range config.DBNAME {
		srcDBTables := make(map[string]bool, 0)
		dnsDBTables := make(map[string]bool, 0)
		queryTableName(&srcDBTables, srcDB[dbname])
		queryTableName(&dnsDBTables, dnsDB[dbname])
		if len(srcDBTables) != len(dnsDBTables) {
			// DROP dnsDB's redundancy table or create new table.
			for k := range srcDBTables {
				if _, ok := dnsDBTables[k]; !ok {
					// Create this table in dns database
					// Step 1: Query from srcDB's table structure.
					rowsInQuerySrcDB, err := srcDB[dbname].Query("show create table `"+k+"`")
					if err != nil {
						log.Errorf("Fetching from source database's table creating SQL code failed,error is: %v", err)
						return
					}

					var noNeedVar string
					var querySqlStr string
					for rowsInQuerySrcDB.Next() {
						err = rowsInQuerySrcDB.Scan(&noNeedVar,&querySqlStr)
						if err != nil {
							log.Fatalf("Scanning Creating SQL code failed, error is: %v", err)
						}
					}
					err = rowsInQuerySrcDB.Close()
					if err != nil {
						log.Fatalf("Closing query source DB's rows failed, error is: %v", err)
					}

					// Step 2: Create this table in dnsDB
					rowsInCreate, err := dnsDB[dbname].Query(querySqlStr)
					if err != nil {
						log.Fatalf("Creating table failed, error is: %v", err)
					}
					err = rowsInCreate.Close()
					if err != nil {
						log.Fatalf("Creating table's query rows close failed, error is: %v", err)
					}
				}
			}
			// This Code Here We'd better not use it. Because We better not delete all of the records in one table.
			//for k := range dnsDBTables{
			//	if _, ok := srcDBTables[k]; ! ok {
			//		// DROP dnsDB's redundancy table
			//		rowsInDrop, err := dnsDB[dbname].Query("drop table if exists " + k)
			//		if err != nil {
			//			log.Fatalf("Dropping table failed, error is: %v", err)
			//		}
			//		err = rowsInDrop.Close()
			//		if err != nil {
			//			log.Fatalf("Dropping table's query rows close failed, error is: %v", err)
			//		}
			//	}
			//}
		}
		//
	}
}

func queryTableName(dbTables *map[string]bool, db *sql.DB) {
	rows, err := db.Query("show tables")
	if err != nil {
		log.Fatalf("execute sql string failed, error is: %v", err)
	}
	var s string
	for rows.Next() {
		err = rows.Scan(&s)
		if err != nil {
			log.Fatalf("rows scan failed, error is: %v", err)
		}
		(*dbTables)[s] = true
	}
	err = rows.Close()
	if err != nil {
		log.Fatalf("rows close failed, error is: %v", err)
	}
}
