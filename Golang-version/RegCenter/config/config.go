package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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
	PORT            int
	ClusterEnabled  bool
	HealthCheckTime = 30
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
	HealthCheckTime = app.GetInt("health_check_time")

	cluster := viper.Sub("cluster")
	ClusterEnabled = cluster.GetBool("enabled")
}
