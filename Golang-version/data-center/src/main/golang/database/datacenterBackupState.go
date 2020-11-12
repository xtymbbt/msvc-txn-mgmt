package database

import (
	log "github.com/sirupsen/logrus"
	"strconv"
)

func updateDataCenterDB(state int) (err error) {
	dbx := mainDB["data_center"]
	sqlStr1 := "update `"+"backup"+"` set backup_state="+strconv.Itoa(state)
	sqlStr2 := " where id=1"

	sqlStr := sqlStr1 + sqlStr2+";"
	result, err := dbx.Exec(sqlStr)
	log.Infof("sql string is: %s", sqlStr)
	if err != nil {
		log.Errorf("execute sql string failed, err is: %v", err)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Errorf("Get last updated id failed, err: %v\n", err)
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		log.Errorf("Get affected rows failed, err: %v\n", err)
		return
	}
	log.Infof("execute sql string succeeded, updated id is: %v, affected rows: %v", id, affected)
	return nil
}
