package main

import (
	"fmt"

	"github.com/dsbell81/go-pgsql-api/datastore"
	"github.com/dsbell81/go-pgsql-api/utils"
)

func main() {
	fmt.Println("hello world")

	fmt.Printf("my server is %s\n", utils.AppConfig.Server)

	myDb, err := datastore.GetPgConnection()
	if err != nil {
		myDb.Close()
	}
}
