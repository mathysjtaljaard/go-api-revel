package db

import (
	"create-go-api/app/models"
	"fmt"

	"github.com/revel/revel"
	"gorm.io/gorm"
)

func CustomAutomigrator(dbConn *gorm.DB) {
	if dbConn == nil {
		panic("No dbConnection given. Cannot Auto Magrate")
	}

	for _, schema := range getDefinedSchemas() {
		revel.AppLog.Debugf("Shema is %v %T", schema, schema)
		err := dbConn.AutoMigrate(schema)

		if err != nil {
			panic(fmt.Errorf("auto migration of schema %v failed with err %v", schema, err))
		}
	}
}

func getDefinedSchemas() []interface{} {
	schemas := []interface{}{
		models.User{},
	}
	return schemas
}
