package database

import (
	"../proto"
	"database/sql"
	log "github.com/sirupsen/logrus"
	"strings"
)

func startDBTX(db *sql.DB, dataS []*commonInfo.HttpRequest, sqlStrS []string, err *error) {
	if len(dataS) != len(sqlStrS) {
		log.Errorf("Sql str length is not as equal to dataS length. Please check!")
		return
	}
	tx, errx := db.Begin()
	if errx != nil {
		*err = errx
		log.Errorf("Transaction Begin failed. err is: %#v\n", errx)
		return
	}
	for i, data := range dataS {
		rows, errx := tx.Query("use " + data.DbName)
		if errx != nil {
			*err = errx
			log.Errorf("Query: use "+data.DbName+" failed. err is: %#v\n", errx)
			return
		}
		errx = rows.Close()
		if errx != nil {
			*err = errx
			log.Errorf("Rows close failed. err is: %#v\n", errx)
			return
		}
		rows, errx = tx.Query(sqlStrS[i])
		if errx != nil {
			*err = errx
			log.Errorf("Query: "+sqlStrS[i]+" failed. err is: %#v\n", errx)
			return
		}
		errx = rows.Close()
		if errx != nil {
			*err = errx
			log.Errorf("Rows close failed. err is: %#v\n", errx)
			return
		}
	}
	*err = tx.Commit()
	if *err != nil {
		log.Errorf("Transaction Commit failed. err is: %#v\n", *err)
	}
}

func dbInsertTX(tableName string, data map[string]string) (sqlStr string, err error) {
	sqlStr1 := "insert into `" + tableName + "` ("
	sqlStr2 := ") values ("
	for key, value := range data {
		sqlStr1 += key + ", "
		sqlStr2 += value + ", "
	}
	sqlStr1 = sqlStr1[:len(sqlStr1)-2]
	sqlStr2 = sqlStr2[:len(sqlStr2)-2] + ");"
	sqlStr = sqlStr1 + sqlStr2
	log.Debugf("sql string is: %s", sqlStr)
	return
}

func dbUpdateTX(tableName string, data map[string]string, query string) (sqlStr string, err error) {
	queryStr := strings.Split(query, ",")
	queryMap := make(map[string]int, 0)
	for i, s := range queryStr {
		queryMap[s] = i
	}
	sqlStr1 := "update `" + tableName + "` set "
	sqlStr2 := " where "
	for key, value := range data {
		if _, ok := queryMap[key]; ok {
			sqlStr2 += key + "=" + value + " and "
			continue
		}
		if len(value) == 0 {
			sqlStr1 += key + "=" + key + value + ","
			continue
		}
		switch value[0] {
		case '+':
			sqlStr1 += key + "=" + key + value + ","
		case '-':
			sqlStr1 += key + "=" + key + value + ","
		case '×':
			sqlStr1 += key + "=" + key + "*" + value[1:] + ","
		case '*':
			sqlStr1 += key + "=" + key + value + ","
		case '÷':
			sqlStr1 += key + "=" + key + "/" + value[1:] + ","
		case '/':
			sqlStr1 += key + "=" + key + value + ","
		case '=':
			sqlStr1 += key + value + ","
		default:
			sqlStr1 += key + "=" + value + ","
		}
	}
	sqlStr1 = sqlStr1[:len(sqlStr1)-1]
	sqlStr2 = sqlStr2[:len(sqlStr2)-5] + ";"
	sqlStr = sqlStr1 + sqlStr2
	log.Debugf("sql string is: %s", sqlStr)
	return
}

func dbDeleteTX(tableName string, data map[string]string, query string) (sqlStr string, err error) {
	queryStr := strings.Split(query, ",")
	sqlStr1 := "delete from `" + tableName + "` "
	sqlStr2 := " where "
	for _, s := range queryStr {
		sqlStr2 += s + "=" + data[s] + " and "
	}
	sqlStr1 = sqlStr1[:len(sqlStr1)-1]
	sqlStr2 = sqlStr2[:len(sqlStr2)-5] + ";"
	sqlStr = sqlStr1 + sqlStr2
	log.Debugf("sql string is: %s", sqlStr)
	return
}

func dbQueryTX(tableName string, data map[string]string) (sqlStr string, err error) {
	// Don't need query operation in transaction management.
	// We recommend you to operate query operation in your own codes.
	return
}
