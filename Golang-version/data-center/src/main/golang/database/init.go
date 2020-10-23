package database

import (
	"../../../resources/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // init mysql driver
	log "github.com/sirupsen/logrus"
)

var db = make(map[string]*sql.DB, 0)

func initDB() (err error) {
	log.Info("connecting to database in "+config.DBUrl+" at port 3306.")
	for _, dbName := range config.DBNAME {
		// dsn := "user:password@tcp(url:3306)/db_name"
		dsn := config.DBUser + ":" + config.DBPassword + "@tcp(" + config.DBUrl + ":3306)/"+dbName
		db[dbName], err = sql.Open(config.DBDriver, dsn)
		if err != nil {
			return
		}
		db[dbName].SetMaxIdleConns(config.DBMaxIdleConn) // 最大闲置连接数
		db[dbName].SetMaxOpenConns(config.DBMaxOpenConn) // 最大连接数
		if err = db[dbName].Ping(); err != nil {
			log.Errorf("open database failed, error is: %v\n", err)
			return
		}
	}
	log.Info("successfully connected to database in "+config.DBUrl+" at port 3306.")
	return nil
}

func init() {
	err := initDB()
	if err != nil {
		log.Fatalf("init database failed, error is: %v\n", err)
	}
}
