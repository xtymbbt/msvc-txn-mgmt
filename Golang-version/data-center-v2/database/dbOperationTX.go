package database

import (
	"data-center-v2/common"
	"database/sql"
	log "github.com/sirupsen/logrus"
	"strings"
)

func startDBTX(db *sql.DB, root *common.TreeNode) (err error) {
	tx, err := db.Begin()
	defer func() {
		if err != nil {
			log.Errorf("Error occured. err is: %#v\nExecuting rollback function.", err)
			err = tx.Rollback()
			if err != nil {
				log.Errorf("transaction rollback failed. err is: %#v\n", err)
			} else {
				log.Warn("transaction has rolled back.")
			}
		}
	}()
	if err != nil {
		log.Errorf("Transaction Begin failed. err is: %#v\n", err)
		return err
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
		rows, err := tx.Query("use " + tmp.DbName)
		if err != nil {
			log.Errorf("Query: use "+tmp.DbName+" failed. err is: %#v\n", err)
			return err
		}
		err = rows.Close()
		if err != nil {
			log.Errorf("Rows close failed. err is: %#v\n", err)
			return err
		}
		rows, err = tx.Query(tmp.SqlStr)
		if err != nil {
			log.Errorf("Query: "+tmp.SqlStr+" failed. err is: %#v\n", err)
			return err
		}
		err = rows.Close()
		if err != nil {
			log.Errorf("Rows close failed. err is: %#v\n", err)
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Errorf("Transaction Commit failed. err is: %#v\n", err)
	}
	return err
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
