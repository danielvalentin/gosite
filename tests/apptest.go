package tests

import (
	"github.com/revel/revel"
	"net/url"
)

type AppTest struct {
	revel.TestSuite
}

func (t *AppTest) Before() {
	println("Set up")
}

func (t AppTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t AppTest) TestThatUserLoginWorks() {
	t.Get("/user/login")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t AppTest) TestThatUserLoginPostWorks() {
	t.PostForm("/user/login", url.Values{
		"username":{"daniel"},
		"password":{"test"},
	})
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t AppTest) TestThatUserRegisterWorks() {
	t.PostForm("/user/login", url.Values{
		"username":{"daniel"},
		"email":{"somer@arte.com"},
		"password":{"test"},
	})
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t AppTest) TestThatUserRegisterPostWorks() {
	t.Get("/user/register")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *AppTest) After() {
	println("Tear down")
}
