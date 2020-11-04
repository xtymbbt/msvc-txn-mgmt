package config

/**
	used to config servers.
 */

const (
	PORT       = ":1996"
	TIMELAPSES = 10
	DBDriver   = "mysql"
	DBUrl = "my.aliyun.com"
	DBUser = "root"
	DBPassword = "Bupt_Bridge_Wang"
	DBMaxIdleConn = 16
	DBMaxOpenConn = 100
)

var (
	DBNAME = []string{"test_order", "test_storage", "test_payment"}
)


