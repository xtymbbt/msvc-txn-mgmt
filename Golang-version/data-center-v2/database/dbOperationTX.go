package database

import (
	"../common"
	"database/sql"
	log "github.com/sirupsen/logrus"
	"strings"
)

func startDBTX(db *sql.DB, root *common.TreeNode, err *error) {
	tx, errx := db.Begin()
	if errx != nil {
		*err = errx
		log.Errorf("Transaction Begin failed. err is: %#v\n", errx)
		return
	}
	// level order query to keep correct sql execute order.
	queue := make([]*common.TreeNode, 0)
	queue = append(queue, root)
	var tmp *common.TreeNode
	for len(queue) != 0 {
		tmp = queue[0]
		queue = queue[1:]
		for _, child := range tmp.Children {
			queue = append(queue, child)
		}
		rows, errx := tx.Query("use " + tmp.Info.DbName)
		if errx != nil {
			*err = errx
			log.Errorf("Query: use "+tmp.Info.DbName+" failed. err is: %#v\n", errx)
			return
		}
		errx = rows.Close()
		if errx != nil {
			*err = errx
			log.Errorf("Rows close failed. err is: %#v\n", errx)
			return
		}
		rows, errx = tx.Query(tmp.SqlStr)
		if errx != nil {
			*err = errx
			log.Errorf("Query: "+tmp.SqlStr+" failed. err is: %#v\n", errx)
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