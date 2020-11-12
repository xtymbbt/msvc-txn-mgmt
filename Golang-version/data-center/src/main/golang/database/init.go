package database

import (
	"../../../resources/config"
	myErr "../error"
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // init mysql driver
	log "github.com/sirupsen/logrus"
)

var mainDB = make(map[string]*sql.DB, 0)
var bakDBs = make([]map[string]*sql.DB, len(config.DBBakUrls))

func initDB() (err error) {
	log.Info("connecting to database in " + config.DBUrl + ".")
	// dsn := "user:password@tcp(url:3306)/db_name"
	dsn := config.DBUser + ":" + config.DBPassword + "@tcp(" + config.DBUrl + ")/" + "data_center"
	mainDB["data_center"], err = sql.Open(config.DBDriver, dsn)
	if err != nil {
		return
	}
	mainDB["data_center"].SetMaxIdleConns(config.DBMaxIdleConn) // 最大闲置连接数
	mainDB["data_center"].SetMaxOpenConns(config.DBMaxOpenConn) // 最大连接数
	if err = mainDB["data_center"].Ping(); err != nil {
		log.Errorf("open  in " + config.DBUrl + " failed,\nerror is: %v\n", err)
		return
	}
	for _, dbName := range config.DBNAME {
		// dsn := "user:password@tcp(url:3306)/db_name"
		dsn := config.DBUser + ":" + config.DBPassword + "@tcp(" + config.DBUrl + ")/" + dbName
		mainDB[dbName], err = sql.Open(config.DBDriver, dsn)
		if err != nil {
			return
		}
		mainDB[dbName].SetMaxIdleConns(config.DBMaxIdleConn) // 最大闲置连接数
		mainDB[dbName].SetMaxOpenConns(config.DBMaxOpenConn) // 最大连接数
		if err = mainDB[dbName].Ping(); err != nil {
			log.Errorf("open  in " + config.DBUrl + " failed,\nerror is: %v\n", err)
			return
		}
	}
	log.Info("successfully connected to database in " + config.DBUrl + ".")
	return nil
}

func initBackupDBs() (err error) {
	if len(config.DBBakUsers) != len(config.DBBakPasswords) || len(config.DBBakPasswords) != len(config.DBBakUrls) {
		err = myErr.NewError(500, "Backup Databases' config is not matched. Please check your configurations")
		return err
	}
	for i := range config.DBBakUrls {
		log.Info("connecting to database in " + config.DBBakUrls[i] + ".")
		bakDBs[i] = make(map[string]*sql.DB, 0)
		for _, dbName := range config.DBNAME {
			// dsn := "user:password@tcp(url:3306)/db_name"
			dsn := config.DBBakUsers[i] + ":" + config.DBBakPasswords[i] + "@tcp(" + config.DBBakUrls[i] + ")/" + dbName
			bakDBs[i][dbName], err = sql.Open(config.DBDriver, dsn)
			if err != nil {
				return
			}
			bakDBs[i][dbName].SetMaxIdleConns(config.DBMaxIdleConn) // 最大闲置连接数
			bakDBs[i][dbName].SetMaxOpenConns(config.DBMaxOpenConn) // 最大连接数
			if err = bakDBs[i][dbName].Ping(); err != nil {
				log.Errorf("open database in " + config.DBBakUrls[i] + " failed,\nerror is: %v\n", err)
				return
			}
		}
		log.Info("successfully connected to database in " + config.DBBakUrls[i] + ".")
	}
	return nil
}

func init() {
	err := initDB()
	if err != nil {
		log.Fatalf("init main database failed,\nerror is: %v\n", err)
	}
	err = initBackupDBs()
	if err != nil {
		log.Fatalf("init backup_databases failed,\nerror is: %v\n", err)
	}
	err = initSync()
	if err != nil {
		log.Fatalf("init sync databases failed,\nerror is: %v\n", err)
	}
	// TODO: 再加上个每隔一段时间使用多数投票的方式进行同步的功能。如果发现数据库之间不同步，则进行报警，并记录下数据不一致的数据库
}
