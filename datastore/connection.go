package datastore

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/dsbell81/go-pgsql-api/utils"
	_ "github.com/lib/pq"
)

var Db *sql.DB

//prepare a database connection
func InitDb() {
	getPgConnection()
}

//get Postgres connection
//disable ssl mode for dev
func getPgConnection() {

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		utils.AppConfig.DbHost,
		utils.AppConfig.DbPort,
		utils.AppConfig.DbUser,
		utils.AppConfig.DbPwd,
		utils.AppConfig.DbName)

	var err error
	Db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("[GetPgConnection]: %s\n", err)
	}

	err = Db.Ping()
	if err != nil {
		log.Fatalf("[TestDbConnection: %s\n", err)
	}

	log.Println("Successfully connected to Database!")
	return
}
