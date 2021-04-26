package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

/**
used to config servers.
本系统首先需要在主数据库中新建状态记录数据库：data_center
然后设计表：
表名：backup
字段：id和backup_state
插入一条记录：id=1 & backup_state=0
*/

var (
	PORT           int
	TIMELAPSES     time.Duration
	DBDriver       string
	DBUrl          string
	DBUser         string
	DBPassword     string
	DBMaxIdleConn  int
	DBMaxOpenConn  int
	EnableBKDB     bool
	DBNAME         []string
	DBBakUrls      []string
	DBBakUsers     []string
	DBBakPasswords []string
)

func initDefaultValue() {
	viper.SetConfigFile("application.yml")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Fatalf("Config file not found. Err is %#v\n", err)
		} else {
			// Config file was found but another error was produced
			log.Fatalf("Config file was found but another error was produced.\nerr is %#v\n", err)
		}
	}
	app := viper.Sub("app")
	PORT = app.GetInt("port")
	TIMELAPSES = time.Second * time.Duration(app.GetInt("timelapses"))

	dbConf := viper.Sub("database")
	DBDriver = dbConf.GetString("driver")
	EnableBKDB = dbConf.GetBool("enable_slave_database")
	DBNAME = []string{}
	for _, db := range dbConf.GetStringSlice("database_names") {
		DBNAME = append(DBNAME, db)
	}
	DBMaxIdleConn = dbConf.GetInt("max_idle_connection")
	DBMaxOpenConn = dbConf.GetInt("max_open_connection")
	mainDB := viper.Sub("database.main")
	DBUrl = mainDB.GetString("url")
	DBUser = mainDB.GetString("user")
	DBPassword = mainDB.GetString("password")

	slaves := viper.Sub("database.slaves")

	DBBakUrls = []string{}
	for _, url := range slaves.GetStringSlice("urls") {
		DBBakUrls = append(DBBakUrls, url)
	}
	DBBakUsers = []string{}
	for _, user := range slaves.GetStringSlice("users") {
		DBBakUsers = append(DBBakUsers, user)
	}
	DBBakPasswords = []string{}
	for _, password := range slaves.GetStringSlice("passwords") {
		DBBakPasswords = append(DBBakPasswords, password)
	}
}
