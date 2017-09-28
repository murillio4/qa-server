package controllers

import (
  _ "github.com/murillio4/stack-server/events"
  m "github.com/murillio4/stack-server/models"
	"github.com/astaxie/beego"

)

type BaseController struct {
	beego.Controller

  user      *m.User
  loggedin  bool
}

func (c *BaseController) Prepare(){

  //Check if token is valid and that token is registered to a user
  if c.Ctx.Input.Header("token") != "" {
    token := RwtToken{}
    v, err := token.validToken(c.Ctx.Input.Header("token"))

    if v == true && err == nil {
      c.user = m.GetUserWithRwt(c.Ctx.Input.Header("token"))
      c.loggedin = true
    }
  }
}

func notAuth() (*m.Error){
  return &m.Error{401, "Not authorized"}
}
