package database

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
	"log"
	"nimbus/core"
	"nimbus/core/config"
)

// import _ "github.com/jinzhu/gorm/dialects/postgres"
// import _ "github.com/jinzhu/gorm/dialects/sqlite"
// import _ "github.com/jinzhu/gorm/dialects/mssql"

type DatabaseFacade struct {
	DB      *gorm.DB
	context *core.Context
}

func NewDatabase() *DatabaseFacade {
	appConfig := config.GetAppConfig().Config
	user := appConfig["DB_USER"]
	database := appConfig["DB_NAME"]
	host := appConfig["DB_HOST"]
	port := appConfig["DB_PORT"]
	password := appConfig["DB_PASSWORD"]
	db, err := mysql(user, password, database, host, port)
	if err != nil {
		log.Fatal(err)
	}
	return &DatabaseFacade{
		DB: db,
	}

}

func mysql(user string, password string, databaseName string, host string, port string) (*gorm.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, databaseName)
	log.Println(connectionString)
	db, err := gorm.Open("mysql", connectionString)
	// remember to close the db
	return db, err
}
