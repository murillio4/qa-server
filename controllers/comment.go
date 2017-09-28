package controllers

import (
		m "github.com/murillio4/stack-server/models"
)

type CommentController struct {
	BaseController
}

func (c *CommentController) CreateAnswerComment() {
	aid, _ := c.GetInt(":aid")
	content := c.GetString("c_content")

	co := m.Comment{Content:content}
	a := m.Answer{ID:uint(aid)}

	e := m.CreateComment(&a, &co)

	if e != nil {
		c.Data["json"] = e
	} else {
		c.Data["json"] = co
	}
	c.ServeJSON()
}

func (c *CommentController) CreateQuestionComment() {
	aid, _ := c.GetInt(":qid")
	content := c.GetString("c_content")

	co := m.Comment{Content:content}
	q := m.Question{ID:uint(aid)}

	e := m.CreateComment(&q, &co)

	if e != nil {
		c.Data["json"] = e
	} else {
		c.Data["json"] = co
	}
	c.ServeJSON()
}
