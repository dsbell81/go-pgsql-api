package datastore

import (
	"github.com/dsbell81/go-pgsql-api/models"
	"time"
)

func CreateUser(user *models.User) (err error) {
	sqlStatement := `
  INSERT INTO users (email, password, created, modified)
  VALUES ($1, crypt($2, gen_salt('bf',8)), $3, $4)
  RETURNING id`

	currentTime := time.Now()
	user.Id = ""
	user.Created = currentTime
	user.Modified = currentTime

	err = Db.QueryRow(sqlStatement, user.Email, user.Password, user.Created, user.Modified).Scan(&user.Id)
	return
}
