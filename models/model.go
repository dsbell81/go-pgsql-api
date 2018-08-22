package models

import (
	"time"
)

type (
	User struct {
		Id        string    `json:"id"`
		Email     string    `json:"email"`
		Password  string    `json:"password,omitempty"`
		Created   time.Time `json:"created"`
		Modified  time.Time `json:"modified"`
		LastLogin time.Time `json:"lastlogin,omitempty"`
	}
)
