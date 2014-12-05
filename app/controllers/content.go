package controllers

import (
	"strings"
	"github.com/revel/revel"
	"goapp/app/models"
)

type Content struct{
	*revel.Controller
}

func (c Content) Index () revel.Result {
	active := "content"
	return c.Render(active)
}

func (c Content) getSlug () string {
	var url string = string(c.Request.URL.Path)
	url = strings.TrimLeft(url, "/")
	url = strings.TrimRight(url, "/")
	return url
}

func (c Content) Find () revel.Result {
	slug := c.getSlug()
	content := new(models.Content).GetByGuid(slug)
	return c.Render(slug, content)
}
