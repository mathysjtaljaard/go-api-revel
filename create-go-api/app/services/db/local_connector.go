package db

import (
	"create-go-api/app/models"

	"github.com/revel/revel"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbConnect *gorm.DB

func GetDbConnector() *gorm.DB {
	if dbConnect == nil {
		initDb()
	}
	revel.AppLog.Debugf("GetDbConnector returning dbConnect %v ", dbConnect)
	return dbConnect
}

func initDb() {
	var err error
	dbConnect, err = gorm.Open(sqlite.Open("./local/local.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	revel.AppLog.Debugf("No Error in initDB, dbConnect %v", dbConnect)
	revel.AppLog.Debug("calling CustomAutomigrator")
}

func CreateUsers() {
	// revel.DevMod and revel.RunMode work here
	// Use this script to check for dev mode and set dev/prod startup scripts here!

	revel.AppLog.Debugf("Connection for DB %v for creating users", GetDbConnector())
	//Clear Table
	GetDbConnector().Where("id > ?", 0).Delete(&models.User{})

	users := []models.User{
		{FirstName: "Randy", LastName: "Wilson"},
		{FirstName: "Zhany", LastName: "Drag"},
		{FirstName: "Bob", LastName: "Builder"},
		{FirstName: "Kat", LastName: "House"},
		{FirstName: "Brent", LastName: "Rivers"},
	}
	revel.AppLog.Debugf("Users to insert %v", users)
	GetDbConnector().CreateInBatches(&users, len(users))

}
