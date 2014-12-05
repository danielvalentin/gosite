package controllers

import (
	"github.com/revel/revel"
	"goapp/app/models"
)

func checkUser(c *revel.Controller) revel.Result {
	if username, ok := c.Session["user"]; ok {
		c.RenderArgs["user"] = models.GetUserByUsername(username)
	}
	return nil
}

func getPages(c *revel.Controller) revel.Result {
	c.RenderArgs["pages"] = new(models.Content).GetAll()
	return nil
}

func templateFuncActive (args ...interface{}) bool {
	if args[0] == "" {
		return false
	}
	return args[0] == args[1]
}

func init() {
	revel.InterceptFunc(checkUser, revel.BEFORE, revel.ALL_CONTROLLERS)
	revel.InterceptFunc(getPages, revel.BEFORE, revel.ALL_CONTROLLERS)
	revel.TemplateFuncs["active"] = templateFuncActive;
}