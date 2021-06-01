package repositories

import (
	"create-go-api/app"
	"create-go-api/app/models"

	"github.com/revel/revel"
)

func GetAllUsers() []models.User {
	users := []models.User{}
	var dbConn = app.DB_Connector
	revel.AppLog.Debugf("Database connnector is %v", dbConn)
	dbConn.Find(&users)
	return users
}

func CreateUser(user *models.User) *models.User {

	app.DB_Connector.Create(&user)
	return user
}
