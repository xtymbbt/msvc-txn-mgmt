package database

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
)

func dbInsert(dbx *sql.DB, tableName string, data map[string]string) (err error) {
	sqlStr1 := "insert into `"+tableName+"` ("
	sqlStr2 := ") values ("
	for key, value := range data {
		sqlStr1 += key+", "
		sqlStr2 += value+", "
	}
	sqlStr1 = sqlStr1[:len(sqlStr1)-2]
	sqlStr2 = sqlStr2[:len(sqlStr2)-2]+");"
	sqlStr := sqlStr1 + sqlStr2
	result, err := dbx.Exec(sqlStr)
	log.Infof("sql string is: %s", sqlStr)
	if err != nil {
		log.Errorf("execute sql string failed, err is: %v", err)
	} else {
		log.Infof("execute sql string succeeded, sql string is: %s, result is: %v", sqlStr, result)
	}
	return
}

func dbUpdate(dbx *sql.DB, tableName string, data map[string]string) (err error) {
	return
}

func dbDelete(dbx *sql.DB, tableName string, data map[string]string) (err error) {
	return
}

func dbQuery(dbx *sql.DB, tableName string, data map[string]string) (err error) {
	return
}