package controllers

import (
	"github.com/revel/revel"
	
)

type Cms struct {
	*revel.Controller
}

func (cms Cms) Index() revel.Result {
	return nil
}
