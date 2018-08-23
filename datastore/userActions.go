package datastore

import (
	"github.com/dsbell81/go-pgsql-api/models"
	"time"
)

func CreateUser(user *models.User) (err error) {
	sqlStatement := `
  INSERT INTO users (email, password, created, modified, last_login)
  VALUES (lower($1), crypt($2, gen_salt('bf',8)), $3, $4, $5)
  RETURNING id`

	currentTime := time.Now()
	user.Id = ""
	user.Created = currentTime
	user.Modified = currentTime
	user.LastLogin = currentTime

	err = Db.QueryRow(sqlStatement, user.Email, user.Password, user.Created, user.Modified, user.LastLogin).Scan(&user.Id)
	return
}

func Login(user *models.User) (u models.User, err error) {

	//check db for user/pwd combination
	sqlStatement := `
	SELECT id, email, created, modified, last_login FROM users 
	WHERE email = lower($1) AND password = crypt($2, password);`

	err = Db.QueryRow(sqlStatement, user.Email, user.Password).Scan(&u.Id, &u.Email, &u.Created, &u.Modified, &u.LastLogin)

	//if login successful, update last_login
	if err == nil {
		sqlStatement := `
  	UPDATE users
  	SET last_login = $1
		WHERE email = lower($2)
		RETURNING id;`

		currentTime := time.Now()
		err = Db.QueryRow(sqlStatement, currentTime, user.Email).Scan(&u.Id)
	}

	return
}
