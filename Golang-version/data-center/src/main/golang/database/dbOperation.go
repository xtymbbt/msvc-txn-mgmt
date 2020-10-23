package database

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
	"strings"
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
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Errorf("Get last insert id failed, err: %v\n", err)
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		log.Errorf("Get affected rows failed, err: %v\n", err)
		return
	}
	log.Infof("execute sql string succeeded, inserted id is: %v, affected rows: %v", id, affected)
	return
}

func dbUpdate(dbx *sql.DB, tableName string, data map[string]string, query string) (err error) {
	queryStr := strings.Split(query, ",")
	queryMap := make(map[string]int, 0)
	for i, s := range queryStr {
		queryMap[s] = i
	}
	sqlStr1 := "update `"+tableName+"` set "
	sqlStr2 := " where "
	for key, value := range data {
		if _, ok := queryMap[key]; ok {
			sqlStr2 += key + "=" + value + " and "
			continue
		}
		if len(value) == 0 {
			sqlStr1 += key+"="+key+value+","
			continue
		}
		switch value[0] {
		case '+':
			sqlStr1 += key+"="+key+value+","
		case '-':
			sqlStr1 += key+"="+key+value+","
		case '×':
			sqlStr1 += key+"="+key+"*"+value[1:]+","
		case '*':
			sqlStr1 += key+"="+key+value+","
		case '÷':
			sqlStr1 += key+"="+key+"/"+value[1:]+","
		case '/':
			sqlStr1 += key+"="+key+value+","
		case '=':
			sqlStr1 += key+value+","
		default:
			sqlStr1 += key+"="+value+","
		}
	}
	sqlStr1 = sqlStr1[:len(sqlStr1)-1]
	sqlStr2 = sqlStr2[:len(sqlStr2)-5]+";"
	sqlStr := sqlStr1 + sqlStr2
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
	return
}

func dbDelete(dbx *sql.DB, tableName string, data map[string]string, query string) (err error) {
	queryStr := strings.Split(query, ",")
	sqlStr1 := "delete from `"+tableName+"` "
	sqlStr2 := " where "
	for _, s := range queryStr {
		sqlStr2 += s + "=" + data[s] + " and "
	}
	sqlStr1 = sqlStr1[:len(sqlStr1)-1]
	sqlStr2 = sqlStr2[:len(sqlStr2)-5]+";"
	sqlStr := sqlStr1 + sqlStr2
	result, err := dbx.Exec(sqlStr)
	log.Infof("sql string is: %s", sqlStr)
	if err != nil {
		log.Errorf("execute sql string failed, err is: %v", err)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Errorf("Get last deleted id failed, err: %v\n", err)
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		log.Errorf("Get affected rows failed, err: %v\n", err)
		return
	}
	log.Infof("execute sql string succeeded, deleted id is: %v, affected rows: %v", id, affected)
	return
}

func dbQuery(dbx *sql.DB, tableName string, data map[string]string) (err error) {
	return
}