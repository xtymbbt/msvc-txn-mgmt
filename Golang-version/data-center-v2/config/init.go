package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	initDefaultValue()
	initLog()
	printv()
}

func printv() {
	fmt.Println("PORT", PORT)
	fmt.Println("TIMELAPSES", TIMELAPSES)
	fmt.Println("DBDriver", DBDriver)
	fmt.Println("DBUrl", DBUrl)
	fmt.Println("DBUser", DBUser)
	fmt.Println("DBPassword", DBPassword)
	fmt.Println("DBMaxIdleConn", DBMaxIdleConn)
	fmt.Println("DBMaxOpenConn", DBMaxOpenConn)
	fmt.Println("EnableBKDB", EnableBKDB)
	fmt.Println("DBNAME", DBNAME)
	fmt.Println("DBBakUrls", DBBakUrls)
	fmt.Println("DBBakUsers", DBBakUsers)
	fmt.Println("DBBakPasswords", DBBakPasswords)
}

// init log config
func initLog() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{
		ForceColors:               true,
		DisableColors:             false,
		ForceQuote:                true,
		DisableQuote:              false,
		EnvironmentOverrideColors: true,
		DisableTimestamp:          false,
		FullTimestamp:             true,
		TimestampFormat:           "MST 2006 Jan 2 Mon PM 15:04:05.000000000",
		DisableSorting:            false,
		SortingFunc:               nil,
		DisableLevelTruncation:    false,
		PadLevelText:              false,
		QuoteEmptyFields:          false,
		FieldMap:                  nil,
		CallerPrettyfier:          nil,
	})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}
