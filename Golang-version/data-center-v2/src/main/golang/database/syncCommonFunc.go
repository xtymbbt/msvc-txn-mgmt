package database

import (
	"../../../resources/config"
	"../common"
	"database/sql"
	log "github.com/sirupsen/logrus"
	"reflect"
)

func syncDBs(srcDB map[string]*sql.DB, dnsDB map[string]*sql.DB) {
	// 进行数据库同步
	for _, dbname := range config.DBNAME {
		srcDBTables := make(map[string]bool, 0)
		dnsDBTables := make(map[string]bool, 0)
		queryTableName(&srcDBTables, srcDB[dbname])
		queryTableName(&dnsDBTables, dnsDB[dbname])
		dropRduTabOrCreNewTab(srcDBTables, dnsDBTables, srcDB[dbname], dnsDB[dbname])
		// Query SrcDB's data and Record them into a map.
		srcDBData := make(map[string][][]interface{}, 0)
		queryDBData(srcDB[dbname], &srcDBData, srcDBTables)
		// Query DnsDB's data and Record them into a map.
		dnsDBData := make(map[string][][]interface{}, 0)
		queryDBData(dnsDB[dbname], &dnsDBData, srcDBTables)
		// Compare src data with dnsDB's data. If not the same, update it.
		compareAndUpdate(srcDBData, dnsDBData, srcDB[dbname], dnsDB[dbname], srcDBTables)
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

func dropRduTabOrCreNewTab(srcDBTables, dnsDBTables map[string]bool, srcDB, dnsDB *sql.DB) {
	// DROP dnsDB's redundancy table or create new table.
	for k := range srcDBTables {
		if _, ok := dnsDBTables[k]; !ok {
			// Create this table in dns database
			// Step 1: Query from srcDB's table structure.
			rowsInQuerySrcDB, err := srcDB.Query("show create table `" + k + "`")
			if err != nil {
				log.Errorf("Fetching from source database's table creating SQL code failed,error is: %v", err)
				return
			}

			var noNeedVar string
			var querySqlStr string
			for rowsInQuerySrcDB.Next() {
				err = rowsInQuerySrcDB.Scan(&noNeedVar, &querySqlStr)
				if err != nil {
					log.Fatalf("Scanning Creating SQL code failed, error is: %v", err)
				}
			}
			err = rowsInQuerySrcDB.Close()
			if err != nil {
				log.Fatalf("Closing query source DB's rows failed, error is: %v", err)
			}

			// Step 2: Create this table in dnsDB
			rowsInCreate, err := dnsDB.Query(querySqlStr)
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
	//		rowsInDrop, err := dnsDB.Query("drop table if exists " + k)
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

func queryDBData(db *sql.DB, resultMap *map[string][][]interface{}, tables map[string]bool) {
	// Query DB's data and Record them into a map.
	for k := range tables {
		(*resultMap)[k] = make([][]interface{}, 0)
		rows, err := db.Query("select * from `" + k + "`;")
		if err != nil {
			log.Fatalf("execute sql string failed, error is: %v", err)
			return
		}

		columnType, err := rows.ColumnTypes()
		result := make([]interface{}, len(columnType))
		for i := range result {
			switch columnType[i].DatabaseTypeName() {
			// "VARCHAR", "TEXT", "NVARCHAR", "DECIMAL", "BOOL", "INT", and "BIGINT".
			case "VARCHAR":
				result[i] = new(string)
			case "TEXT":
				result[i] = new(string)
			case "NVARCHAR":
				result[i] = new(string)
			case "DECIMAL":
				result[i] = new(float64)
			case "BOOL":
				result[i] = new(bool)
			case "INT":
				result[i] = new(int)
			case "BIGINT":
				result[i] = new(int64)
			default:
				result[i] = new(string)
			}
		}

		for rows.Next() {
			oneResult := make([]interface{}, 0)
			err = rows.Scan(result...)
			if err != nil {
				log.Fatalf("rows scan failed,error is: %v", err)
				return
			}
			for _, s := range result {
				oneResult = append(oneResult, reflect.ValueOf(s).Elem().Interface())
			}
			(*resultMap)[k] = append((*resultMap)[k], oneResult)
		}
		err = rows.Close()
		if err != nil {
			log.Fatalf("rows scan failed,error is: %v", err)
			return
		}
	}
}

func compareAndUpdate(srcDBData, dnsDBData map[string][][]interface{}, srcDB, dnsDB *sql.DB, tables map[string]bool) {
	// Compare src data with dnsDB's data. If not the same, update it.
	for table := range tables {
		breakOrNot := false
		if len(srcDBData[table]) != len(dnsDBData[table]) {
			// Drop dnsDB table & insert srcDBData into it.
			dropAndInsert(srcDBData, srcDB, dnsDB, table)
			continue
		}
		for i, values := range srcDBData[table] {
			for j, value := range values {
				if dnsDBData[table][i][j] == value {
					continue
				} else {
					// Drop dnsDB table & insert srcDBData into it.
					dropAndInsert(srcDBData, srcDB, dnsDB, table)
					breakOrNot = true
					break
				}
			}
			if breakOrNot {
				break
			}
		}
	}
}

func dropAndInsert(srcDBData map[string][][]interface{}, srcDB *sql.DB, dnsDB *sql.DB, table string) {
	// DROP dnsDB's  table
	rowsInDrop, err := dnsDB.Query("drop table if exists `" + table + "`")
	if err != nil {
		log.Fatalf("Dropping table failed, error is: %v", err)
	}
	err = rowsInDrop.Close()
	if err != nil {
		log.Fatalf("Dropping table's query rows close failed, error is: %v", err)
	}
	// Create this table in dns database
	// Step 1: Query from srcDB's table structure.
	rowsInQuerySrcDB, err := srcDB.Query("show create table `" + table + "`")
	if err != nil {
		log.Errorf("Fetching from source database's table creating SQL code failed,error is: %v", err)
		return
	}

	var noNeedVar string
	var querySqlStr string
	for rowsInQuerySrcDB.Next() {
		err = rowsInQuerySrcDB.Scan(&noNeedVar, &querySqlStr)
		if err != nil {
			log.Fatalf("Scanning Creating SQL code failed, error is: %v", err)
		}
	}
	err = rowsInQuerySrcDB.Close()
	if err != nil {
		log.Fatalf("Closing query source DB's rows failed, error is: %v", err)
	}
	// Step 2: Create this table in dnsDB
	rowsInCreate, err := dnsDB.Query(querySqlStr)
	if err != nil {
		log.Fatalf("Creating table failed, error is: %v", err)
	}
	err = rowsInCreate.Close()
	if err != nil {
		log.Fatalf("Creating table's query rows close failed, error is: %v", err)
	}

	// Insert Data.
	for _, oneData := range srcDBData[table] {
		oneQuery := "insert into `" + table + "` " + "values ("
		for _, datum := range oneData {
			switch datum.(type) {
			case string:
				oneQuery += "'" + datum.(string) + "'" + ","
			default:
				oneQuery += common.StrVal(datum) + ","
			}
		}
		oneQuery = oneQuery[:len(oneQuery)-1] + ")"
		result, err := dnsDB.Exec(oneQuery)
		log.Infof("sql string is: %s", oneQuery)
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
	}
}
