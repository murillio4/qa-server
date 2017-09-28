package controllers

import (
  m "github.com/murillio4/stack-server/models"
  "github.com/sadlil/go-trigger"
)

type QuestionController struct {
	BaseController
}

func (c *QuestionController) GetQuestion() {
  qid, _ := c.GetInt(":qid")
  q := &m.Question{ID:uint(qid)}
  e := m.GetQuestion(q)

  if e != nil {
    c.Data["json"] = e
  } else {
    trigger.Fire("register-view", &q)
    c.Data["json"] = q
  }
  c.ServeJSON()
}

func (c *QuestionController) CreateQuestion() {
  if c.loggedin {
    title := c.GetString("q_title")
    content := c.GetString("q_content")
    tags := c.GetString("q_tags")

    _ = tags

    q := &m.Question{Title:title, Content:content}
    e := m.CreateQuestion(q)

    if e != nil {
      c.Data["json"] = e
    } else {
      c.Data["json"] = q
    }
  } else {
    c.Data["json"] = notAuth()
  }
  c.ServeJSON()
}
