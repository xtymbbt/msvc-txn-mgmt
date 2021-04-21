package config

/**
used to config servers.
本系统首先需要在主数据库中新建状态记录数据库：data_center
然后设计表：
表名：backup
字段：id和backup_state
插入一条记录：id=1 & backup_state=0
*/

const (
	PORT          = ":1996"
	TIMELAPSES    = 50
	DBDriver      = "mysql"
	DBUrl         = "my.aliyun.com:3306"
	DBUser        = "root"
	DBPassword    = "Bupt_Bridge_Wang"
	DBMaxIdleConn = 16
	DBMaxOpenConn = 100
)

var (
	DBNAME = []string{"test_order", "test_storage", "test_payment"}

	DBBakUrls = []string{
		"127.0.0.1:3306",
		"10.112.12.81:3306",
		"10.112.221.144:3306",
		"10.112.196.254:3306",
	}
	DBBakUsers = []string{
		"root",
		"root",
		"root",
		"root",
	}
	DBBakPasswords = []string{
		"123456",
		"root123456",
		"123456",
		"123456",
	}
)
