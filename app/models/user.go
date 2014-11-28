package models

import (
	"github.com/revel/revel"
	"time"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type User struct {
	Id						int64
	Email					string
	Username				string
	Password				string
	Logins					int64
	Last_login				int64
	Created					int64
}

func (u User) IsLogged() bool {
	return false
}

func (u User) Validate(v *revel.Validation) {
	v.Required(u.Email).Message("E-mail is required")
	v.Required(u.Username).Message("Username is required")
	v.Required(u.Password).Message("Password is required")
}

func (u User) Save() {
	conn := getDB()
	defer conn.Db.Close()
	
	u.Created = time.Now().Unix()
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		panic(err.Error())
	}
	u.Password = string(hashedPass)
	
	InsertErr := conn.Insert(&u)
	if InsertErr != nil {
		panic(InsertErr.Error())
	}
}

func (u User) CheckCredentials(username, password string, v *revel.Validation) bool {
	
	v.Required(username).Message("Username required")
	v.Required(password).Message("Password required")
	
	conn := getDB();
	defer conn.Db.Close()
	
	var query string
	if strings.Contains(username, "@") {
		query = "SELECT * FROM `users` WHERE `email` = ?"
	} else {
		query = "SELECT * FROM `users` WHERE `username` = ?"
	}
	
	var user User
	err := conn.SelectOne(&user, query, username)
	if err != nil {
		return false
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return false
	}
	return true
	
}
