package controllers

import (
	"github.com/revel/revel"
	"goapp/app/models"
)

type User struct {
	*revel.Controller
}

func (u User) Index () revel.Result {
	if !new(models.User).IsLogged() {
		return u.Redirect(User.Login)
	} else {
		return u.Render()
	}
}

func (u User) Login () revel.Result {
	return u.Render()
}

func (u User) DoLogin(username, password string) revel.Result {
	user := new(models.User)
	result := user.CheckCredentials(username, password, u.Validation)
	if u.Validation.HasErrors() {
		u.Validation.Keep()
		u.FlashParams()
		return u.Redirect(User.Login)
	}
	if result {
		u.Flash.Success("Logged in!")
	} else {
		u.Flash.Error("Wrong username or password")
		u.FlashParams()
	}
	return u.Redirect(User.Login)
}

func (u User) Register () revel.Result {
	return u.Render()
}

func (u User) DoUserRegistration (email, username, password string) revel.Result {
	user := new(models.User)
	user.Email = email
	user.Username = username
	user.Password = password
	user.Validate(u.Validation)
	if u.Validation.HasErrors() {
		u.Validation.Keep()
		u.FlashParams()
		return u.Redirect(User.Register)
	} else {
		user.Save()
		u.Flash.Success("You have been registerred and can now log in. Welcome!");
		return u.Redirect(User.Index)
	}
}

