package controllers

import (
	m "github.com/murillio4/stack-server/models"
  "github.com/sadlil/go-trigger"
  "time"
)

type UserController struct {
	BaseController
}


func (c *UserController) Register() {
  username := c.GetString("username")
  password := c.GetString("password")

  u := &m.User{Username:username, Password:password}
  err := m.CreateUser(u)

  if err != nil {
    c.Data["json"] = err
  } else {
    trigger.FireBackground("send-confirm-email", u)

    rt := RwtToken{
      ID: u.ID,
      Expires:  time.Now().Unix() + 3600,
    }

    tokenString, _ := rt.getToken()
    m.AddRwtToUser(u, tokenString)
    c.Data["json"] = tokenString
  }
  c.ServeJSON()
}

func (c *UserController) SignIn() {
  username := c.GetString("username")
  password := c.GetString("password")

  u := &m.User{Username:username, Password:password}
  err := m.SignInUser(u)

  if err != nil {
    c.Data["json"] = err
  } else {
    trigger.FireBackground("send-confirm-email", u)

    rt := RwtToken{
      ID: u.ID,
      Expires:  time.Now().Unix() + 3600,
    }

    tokenString, _ := rt.getToken()
    m.AddRwtToUser(u, tokenString)
    c.Data["json"] = tokenString
  }
  c.ServeJSON()
}
