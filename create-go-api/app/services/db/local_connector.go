package db

import (
	"github.com/revel/revel"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbConnect *gorm.DB

func GetDbConnector() *gorm.DB {
	if dbConnect == nil {
		initDb()
	}
	return dbConnect
}

func initDb() {
	dbConnect, err := gorm.Open(sqlite.Open("./local/local.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	revel.AppLog.Debugf("No Error in initDB, dbConnect %v", &dbConnect)
	revel.AppLog.Debug("calling CustomAutomigrator")
	CustomAutomigrator(dbConnect)
}
