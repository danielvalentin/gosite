package controllers

import (
	"github.com/revel/revel"
	"goapp/app/models"
)

type User struct {
	*revel.Controller
}

func (u User) Index () revel.Result {
	if u.connected() != nil {
		return u.Redirect(User.Login)
	} else {
		active := "user"
		return u.Render(active)
	}
}

func (u User) Login () revel.Result {
	active := "user"
	return u.Render(active)
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
		u.Session["user"] = username
		u.Flash.Success("Logged in!")
		return u.Redirect(User.Index)
	} else {
		u.Flash.Error("Wrong username or password")
		u.FlashParams()
		return u.Redirect(User.Login)
	}
}

func (u User) AddUser() revel.Result {
	if user := u.connected(); user != nil {
		u.RenderArgs["user"] = user
	}
	return nil
}

func (u User) connected() *models.User {
	if u.RenderArgs["user"] != nil {
		return u.RenderArgs["user"].(*models.User)
	}
	if username, ok := u.Session["user"]; ok {
		return models.GetUserByUsername(username)
	}
	return nil
}

func (u User) Logout() revel.Result {
	for k := range u.Session {
		delete(u.Session, k)
	}
	return u.Redirect(User.Login)
}

func (u User) Register () revel.Result {
	active := "user"
	return u.Render(active)
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
		u.Session["user"] = user.Username
		u.Flash.Success("You have been registerred and can now log in. Welcome!");
		return u.Redirect(User.Index)
	}
}

