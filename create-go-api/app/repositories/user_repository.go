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

func CreateUser(firstName, lastName string) models.User {
	user := models.User{
		FirstName: firstName,
		LastName:  lastName,
	}

	app.DB_Connector.Create(user)
	return user
}
