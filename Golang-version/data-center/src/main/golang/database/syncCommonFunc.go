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
			// TODO: DROP dnsDB & Sync dnsDB using srcDB
			continue
		}

	}
}

func queryTableName(dbTables *map[string]bool, db *sql.DB) {
	rows, err := db.Query("show tables")
	if err != nil {
		log.Fatalf("execute sql string failed,\nerror is: %v\n", err)
	}
	var s string
	for rows.Next() {
		err = rows.Scan(&s)
		if err != nil {
			log.Fatalf("rows scan failed,\nerror is: %v\n", err)
		}
		(*dbTables)[s] = true
	}
	err = rows.Close()
	if err != nil {
		log.Fatalf("rows close failed,\nerror is: %v\n", err)
	}
}
