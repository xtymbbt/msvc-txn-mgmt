package database

import (
	"../proto"
)

func Write(dataS []*commonInfo.HttpRequest) (err error) {
	for _, data := range dataS {
		dbx := db[data.DbName]
		/**
		 * 根据data的两个method判断是增删改查的哪个操作
		 * true true = 增
		 * true false = 删
		 * false true = 改
		 * false false = 查
		 */
		if data.Method1 {
			if data.Method2 {
				err = dbInsert(dbx, data.TableName, data.Data)
				if err != nil {
					return
				}
			} else {
				err = dbDelete(dbx, data.TableName, data.Data)
				if err != nil {
					return
				}
			}
		} else {
			if data.Method2 {
				err = dbUpdate(dbx, data.TableName, data.Data)
				if err != nil {
					return
				}
			} else {
				err = dbQuery(dbx, data.TableName, data.Data)
				if err != nil {
					return
				}
			}
		}
	}
	return
}
