package datastore

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/dsbell81/go-pgsql-api/utils"
	_ "github.com/lib/pq"
)

type (
	PgDatabase struct {
		Db *sql.DB
	}
)

//prepare a database connection
//disable sslmode for dev
func GetPgConnection() (pgdb PgDatabase, err error) {

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		utils.AppConfig.DbHost,
		utils.AppConfig.DbPort,
		utils.AppConfig.DbUser,
		utils.AppConfig.DbPwd,
		utils.AppConfig.DbName)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("[GetPgConnection]: %s\n", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("[TestDbConnection: %s\n", err)
	}

	log.Println("Successfully connected to Database!")
	pgdb.Db = db
	return
}

//relase connection resources
func (p *PgDatabase) Close() (err error) {
	if p.Db == nil {
		return
	}

	err = p.Db.Close()
	if err != nil {
		log.Printf("Error closing database connection: %s\n", err)
	}
	return
}
